package main

import (
	"os"

	"github.com/initialcapacity/golang-starter/pkg/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	_ "github.com/lib/pq"
)

func newDataAnalyzer() workflowsupport.WorkScheduler[dataanalyzer.Task] {
	url := os.Getenv("POSTGRESQL_URL")
	if url == "" {
		panic("oops, unable to find postgres url.")
	}
	db, _ := databasesupport.Open(url)
	gateway := dataanalyzer.DataGateway{DB: db}

	worker := dataanalyzer.Worker[dataanalyzer.Task]{Gateway: gateway}
	finder := dataanalyzer.WorkFinder[dataanalyzer.Task]{Gateway: gateway, Results: make(chan bool)}
	list := []workflowsupport.Worker[dataanalyzer.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[dataanalyzer.Task](&finder, list, 12_000)
	return scheduler
}

func main() {
	scheduler := newDataAnalyzer()
	scheduler.Start()
	select {}
}
