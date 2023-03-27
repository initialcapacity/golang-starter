package dataanalyzer

import (
	"log"
	"sync/atomic"
)

type Task struct {
	Data []byte
}

type Worker struct {
	Gateway DataGateway
}

func (a *Worker) Run(w Task) error {
	dataGateway := a.Gateway
	return dataGateway.Save(w.Data)
}

func NewWorkFinder(gateway DataGateway, c chan bool) WorkFinder {
	return WorkFinder{Gateway: gateway, Results: c}
}

type WorkFinder struct {
	Gateway DataGateway
	Results chan bool
	Panics  uint64
}

func (a *WorkFinder) MarkErroneous(task Task) {
	defer func() {
		if err := recover(); err != nil {
			atomic.AddUint64(&a.Panics, 1)
			log.Println("recovering analyzer work.")
		}
	}()
	a.Results <- false
}

func (a *WorkFinder) MarkCompleted(task Task) {
	defer func() {
		if err := recover(); err != nil {
			atomic.AddUint64(&a.Panics, 1)
			log.Println("recovering analyzer work.")
		}
	}()
	a.Results <- true
}

func (a *WorkFinder) Stop() {
	close(a.Results)
}

func (a *WorkFinder) FindRequested() []Task {
	dataGateway := a.Gateway
	records, _ := dataGateway.Find()
	s := make([]Task, len(records))
	for i, r := range records {
		s[i] = Task{r.Data}
	}
	return s
}
