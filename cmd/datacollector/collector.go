package main

import (
	"github.com/initialcapacity/golang-starter/internal/datacollector"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
)

func newDataCollector() workflowsupport.WorkScheduler[datacollector.Task] {
	gateway := datacollector.DataGateway{Urls: []string{"https://feed.infoq.com/"}}
	worker := datacollector.Worker[datacollector.Task]{}
	finder := datacollector.WorkFinder[datacollector.Task]{Gateway: gateway, Results: make(chan bool)}
	list := []workflowsupport.Worker[datacollector.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[datacollector.Task](&finder, list, 12_000)
	return scheduler
}

func main() {
	scheduler := newDataCollector()
	scheduler.Start()
	select {}
}
