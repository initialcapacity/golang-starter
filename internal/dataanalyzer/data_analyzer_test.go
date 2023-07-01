package dataanalyzer_test

import (
	"testing"
	"time"

	"github.com/initialcapacity/golang-starter/internal/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	"github.com/stretchr/testify/assert"
)

func TestWorkflow(t *testing.T) {
	db, _ := databasesupport.Open("postgres://starter:starter@localhost:5432/starter_test?sslmode=disable")
	gateway := dataanalyzer.DataGateway{DB: db}
	worker := dataanalyzer.Worker{Gateway: gateway}
	finder := dataanalyzer.NewWorkFinder(gateway)
	list := []workflowsupport.Worker[dataanalyzer.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[dataanalyzer.Task](&finder, list, 10)
	scheduler.Start()
	for finder.Results.Load() <= 2 {
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
	assert.True(t, finder.Results.Load() > 2)
	scheduler.Stop()
}

func TestWorkflow_Stop(t *testing.T) {
	gateway := dataanalyzer.DataGateway{}
	finder := dataanalyzer.NewWorkFinder(gateway)
	finder.Stop()
}
