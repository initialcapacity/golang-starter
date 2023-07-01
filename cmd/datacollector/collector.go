package main

import (
	"log"
	"os"
	"sync/atomic"

	"github.com/initialcapacity/golang-starter/internal/datacollector"
	"github.com/initialcapacity/golang-starter/pkg/databasesupport"
	"github.com/initialcapacity/golang-starter/pkg/workflowsupport"
	_ "github.com/lib/pq"
)

func newDataCollector() workflowsupport.WorkScheduler[datacollector.Task] {
	url := os.Getenv("POSTGRESQL_URL")
	if url == "" {
		panic("oops, unable to find postgres url.")
	}
	db, _ := databasesupport.Open(url)
	gateway := datacollector.DataGateway{DB: db}

	worker := datacollector.Worker{Gateway: gateway}
	finder := datacollector.WorkFinder{Gateway: gateway, Results: atomic.Int64{}}
	list := []workflowsupport.Worker[datacollector.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[datacollector.Task](&finder, list, 12_000)
	return scheduler
}

func main() {
	log.Println("Starting the data collector.")
	scheduler := newDataCollector()
	scheduler.Start()
	select {}
}
