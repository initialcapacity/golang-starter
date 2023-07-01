package main

import (
	"log"
	"os"
	"sync/atomic"

	"github.com/initialcapacity/golang-starter/internal/dataanalyzer"
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

	worker := dataanalyzer.Worker{Gateway: gateway}
	finder := dataanalyzer.WorkFinder{Gateway: gateway, Results: atomic.Int64{}}
	list := []workflowsupport.Worker[dataanalyzer.Task]{&worker}
	scheduler := workflowsupport.NewScheduler[dataanalyzer.Task](&finder, list, 12_000)
	return scheduler
}

func main() {
	log.Println("Starting the data analyzer.")
	scheduler := newDataAnalyzer()
	scheduler.Start()
	select {}
}
