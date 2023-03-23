package workflowsupport_test

import (
	"errors"
	"log"
	"testing"

	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	"github.com/stretchr/testify/assert"
)

type NoopTask struct {
	data string
}

type NoopWorker[T any] struct {
}

func (n *NoopWorker[T]) Run(task T) error {
	log.Printf("doing work.\n")
	return nil
}

type NoopWorkFinder[T NoopTask] struct {
	Results chan bool
}

func NewNoopWorkFinder[T NoopTask]() NoopWorkFinder[T] {
	return NoopWorkFinder[T]{Results: make(chan bool)}
}

func (n *NoopWorkFinder[T]) MarkErroneous(task T) {
	n.Results <- false
	log.Println("non completed task")
}

func (n *NoopWorkFinder[T]) MarkCompleted(task T) {
	n.Results <- true
	log.Println("completed task")
}

func (n *NoopWorkFinder[T]) Stop() {
	close(n.Results)
}

func (n NoopWorkFinder[T]) FindRequested() []NoopTask {
	log.Printf("finding work.\n")

	return []NoopTask{
		{"someInfo"},
		{"someMoreInfo"},
		{"andSomeMoreInfo"},
	}
}

func TestWorkflow(t *testing.T) {
	var worker NoopWorker[NoopTask]
	finder := NewNoopWorkFinder[NoopTask]()

	list := []workflowsupport.Worker[NoopTask]{&worker}
	scheduler := workflowsupport.NewScheduler[NoopTask](&finder, list, 50)
	scheduler.Start()

	for i := 0; i < 3; i++ {
		assert.True(t, <-finder.Results)
	}

	scheduler.Stop()
}

type ErroneousWorker[T any] struct {
}

func (n *ErroneousWorker[T]) Run(task T) error {
	log.Printf("doing work.\n")
	return errors.New("oops")
}

func TestErroneousWorkflow(t *testing.T) {
	var worker ErroneousWorker[NoopTask]
	finder := NewNoopWorkFinder()

	list := []workflowsupport.Worker[NoopTask]{&worker}
	scheduler := workflowsupport.NewScheduler[NoopTask](&finder, list, 50)
	scheduler.Start()

	for i := 0; i < 3; i++ {
		assert.False(t, <-finder.Results)
	}

	scheduler.Stop()
}
