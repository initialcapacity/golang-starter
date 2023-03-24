package dataanalyzer_test

import (
	"testing"

	"github.com/initialcapacity/golang-starter/pkg/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	"github.com/stretchr/testify/assert"
)

func TestWorkflow(t *testing.T) {
	db, _ := databasesupport.Open("postgres://starter:starter@localhost:5432/starter_development?sslmode=disable")
	gateway := dataanalyzer.DataGateway{DB: db}
	worker := dataanalyzer.Worker[dataanalyzer.Task]{Gateway: gateway}
	finder := dataanalyzer.NewWorkFinder[dataanalyzer.Task](gateway, make(chan bool))
	list := []workflowsupport.Worker[dataanalyzer.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[dataanalyzer.Task](&finder, list, 100)
	scheduler.Start()
	assert.True(t, <-finder.Results)
}

func TestWorkflow_Stop(t *testing.T) {
	gateway := dataanalyzer.DataGateway{}
	finder := dataanalyzer.NewWorkFinder(gateway, make(chan bool))
	finder.Stop()
}

func TestWorker_PanicOnCompleted(t *testing.T) {
	withClosedChannel(func(finder dataanalyzer.WorkFinder[dataanalyzer.Task]) {
		finder.MarkCompleted(dataanalyzer.Task{})
		assert.Equal(t, uint64(1), finder.Panics)
	})
}

func TestWorker_PanicOnErroneous(t *testing.T) {
	withClosedChannel(func(finder dataanalyzer.WorkFinder[dataanalyzer.Task]) {
		finder.MarkErroneous(dataanalyzer.Task{})
		assert.Equal(t, uint64(1), finder.Panics)
	})
}

func withClosedChannel(f func(finder dataanalyzer.WorkFinder[dataanalyzer.Task])) {
	results := make(chan bool)
	finder := dataanalyzer.WorkFinder[dataanalyzer.Task]{Results: results}
	close(results)
	f(finder)
}
