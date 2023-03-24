package dataanalyzer

import (
	"log"
	"sync/atomic"
)

type Task struct {
	Data []byte
}

type Worker[T Task] struct {
	Gateway interface{}
}

func (a *Worker[T]) Run(w Task) error {
	dataGateway := a.Gateway.(DataGateway)
	return dataGateway.Save(w.Data)
}

func NewWorkFinder[T Task](gateway DataGateway, c chan bool) WorkFinder[T] {
	return WorkFinder[T]{Gateway: gateway, Results: c}
}

type WorkFinder[T Task] struct {
	Gateway interface{}
	Results chan bool
	Panics  uint64
}

func (a *WorkFinder[T]) MarkErroneous(task T) {
	defer func() {
		if err := recover(); err != nil {
			atomic.AddUint64(&a.Panics, 1)
			log.Println("recovering analyzer work.")
		}
	}()
	a.Results <- false
}

func (a *WorkFinder[T]) MarkCompleted(task T) {
	defer func() {
		if err := recover(); err != nil {
			atomic.AddUint64(&a.Panics, 1)
			log.Println("recovering analyzer work.")
		}
	}()
	a.Results <- true
}

func (a *WorkFinder[T]) Stop() {
	close(a.Results)
}

func (a *WorkFinder[T]) FindRequested() []Task {
	dataGateway := a.Gateway.(DataGateway)
	records, _ := dataGateway.Find()
	s := make([]Task, len(records))
	for i, r := range records {
		s[i] = Task{r.Data}
	}
	return s
}
