package dataanalyzer_test

import (
	"testing"

	dataanalyzer2 "github.com/initialcapacity/golang-starter/internal/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	"github.com/stretchr/testify/assert"
)

func TestWorkflow(t *testing.T) {
	db, _ := databasesupport.Open("postgres://starter:starter@localhost:5432/starter_development?sslmode=disable")
	gateway := dataanalyzer2.DataGateway{DB: db}
	worker := dataanalyzer2.Worker[dataanalyzer2.Task]{Gateway: gateway}
	finder := dataanalyzer2.NewWorkFinder[dataanalyzer2.Task](gateway, make(chan bool))
	list := []workflowsupport.Worker[dataanalyzer2.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[dataanalyzer2.Task](&finder, list, 100)
	scheduler.Start()
	assert.True(t, <-finder.Results)
}

func TestWorkflow_Stop(t *testing.T) {
	gateway := dataanalyzer2.DataGateway{}
	finder := dataanalyzer2.NewWorkFinder(gateway, make(chan bool))
	finder.Stop()
}

func TestWorker_PanicOnCompleted(t *testing.T) {
	withClosedChannel(func(finder dataanalyzer2.WorkFinder[dataanalyzer2.Task]) {
		finder.MarkCompleted(dataanalyzer2.Task{})
		assert.Equal(t, uint64(1), finder.Panics)
	})
}

func TestWorker_PanicOnErroneous(t *testing.T) {
	withClosedChannel(func(finder dataanalyzer2.WorkFinder[dataanalyzer2.Task]) {
		finder.MarkErroneous(dataanalyzer2.Task{})
		assert.Equal(t, uint64(1), finder.Panics)
	})
}

func withClosedChannel(f func(finder dataanalyzer2.WorkFinder[dataanalyzer2.Task])) {
	results := make(chan bool)
	finder := dataanalyzer2.WorkFinder[dataanalyzer2.Task]{Results: results}
	close(results)
	f(finder)
}
