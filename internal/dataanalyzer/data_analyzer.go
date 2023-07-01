package dataanalyzer

import (
	"sync/atomic"
)

type Task struct {
	Data []byte
}

type Worker struct {
	Gateway DataGateway
}

func (a *Worker) Run(w Task) error {
	return a.Gateway.Save(w.Data)
}

func NewWorkFinder(gateway DataGateway) WorkFinder {
	return WorkFinder{Gateway: gateway, Results: atomic.Int64{}}
}

type WorkFinder struct {
	Gateway DataGateway
	Results atomic.Int64
	Panics  uint64
}

func (a *WorkFinder) MarkErroneous(task Task) {
	a.Results.Add(1)
}

func (a *WorkFinder) MarkCompleted(task Task) {
	a.Results.Add(1)
}

func (a *WorkFinder) Stop() {
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
