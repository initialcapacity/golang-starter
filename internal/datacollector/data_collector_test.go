package datacollector_test

import (
	"testing"

	"github.com/initialcapacity/golang-starter/internal/datacollector"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	"github.com/stretchr/testify/assert"
)

func TestWorkflow(t *testing.T) {
	withOpenChannel("https://feed.infoq.com/", func(finder datacollector.WorkFinder[datacollector.Task]) {
		assert.True(t, <-finder.Results)
	})
}

func TestWorkflow_Fails(t *testing.T) {
	withOpenChannel("https://feed.infoq./", func(finder datacollector.WorkFinder[datacollector.Task]) {
		assert.False(t, <-finder.Results)
	})
}

func TestWorker_PanicOnCompleted(t *testing.T) {
	withClosedChannel(func(finder datacollector.WorkFinder[datacollector.Task]) {
		finder.MarkCompleted(datacollector.Task{})
		assert.Equal(t, uint64(1), finder.Panics)
	})
}

func TestWorker_PanicOnErroneous(t *testing.T) {
	withClosedChannel(func(finder datacollector.WorkFinder[datacollector.Task]) {
		finder.MarkErroneous(datacollector.Task{})
		assert.Equal(t, uint64(1), finder.Panics)
	})
}

func TestWorkflow_Stop(t *testing.T) {
	finder := datacollector.NewWorkFinder(datacollector.DataGateway{Urls: []string{"https://feed.infoq.com/"}}, make(chan bool))
	finder.Stop()
}

func withOpenChannel[T datacollector.Task](url string, f func(finder datacollector.WorkFinder[datacollector.Task])) {
	var worker datacollector.Worker[datacollector.Task]
	finder := datacollector.NewWorkFinder[datacollector.Task](datacollector.DataGateway{Urls: []string{url}}, make(chan bool))
	list := []workflowsupport.Worker[datacollector.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[datacollector.Task](&finder, list, 100)
	scheduler.Start()
	f(finder)
}

func withClosedChannel[T datacollector.Task](f func(finder datacollector.WorkFinder[datacollector.Task])) {
	results := make(chan bool)
	finder := datacollector.WorkFinder[datacollector.Task]{Results: results}
	close(results)
	f(finder)
}
