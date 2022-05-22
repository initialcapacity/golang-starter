package main

import (
	"github.com/initialcapacity/golang-starter/pkg/dataanalyzer"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	_ "github.com/lib/pq"
	"os"
)

func main() {
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
	scheduler.Start()
	select {}
}
