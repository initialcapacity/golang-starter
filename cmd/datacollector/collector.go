package main

import (
	"github.com/initialcapacity/golang-starter/pkg/datacollector"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
)

func main() {
	gateway := datacollector.DataGateway{Urls: []string{"https://feed.infoq.com/"}}

	worker := datacollector.Worker[datacollector.Task]{}
	finder := datacollector.WorkFinder[datacollector.Task]{Gateway: gateway, Results: make(chan bool)}
	list := []workflowsupport.Worker[datacollector.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[datacollector.Task](&finder, list, 12_000)
	scheduler.Start()
	select {}
}
