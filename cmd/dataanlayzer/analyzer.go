package main

import (
	"os"

	dataanalyzer2 "github.com/initialcapacity/golang-starter/internal/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	_ "github.com/lib/pq"
)

func newDataAnalyzer() workflowsupport.WorkScheduler[dataanalyzer2.Task] {
	url := os.Getenv("POSTGRESQL_URL")
	if url == "" {
		panic("oops, unable to find postgres url.")
	}
	db, _ := databasesupport.Open(url)
	gateway := dataanalyzer2.DataGateway{DB: db}

	worker := dataanalyzer2.Worker[dataanalyzer2.Task]{Gateway: gateway}
	finder := dataanalyzer2.WorkFinder[dataanalyzer2.Task]{Gateway: gateway, Results: make(chan bool)}
	list := []workflowsupport.Worker[dataanalyzer2.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[dataanalyzer2.Task](&finder, list, 12_000)
	return scheduler
}

func main() {
	scheduler := newDataAnalyzer()
	scheduler.Start()
	select {}
}
