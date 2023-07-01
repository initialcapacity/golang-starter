package datacollector

import (
	"sync/atomic"
)

type Task struct {
	Url string
}

type Worker struct {
	Gateway DataGateway
}

func (a *Worker) Run(w Task) error {
	body, _ := MakeRequest(w)

	return a.Gateway.Save(body)
}

func NewWorkFinder(gateway DataGateway) WorkFinder {
	return WorkFinder{Gateway: gateway, Results: atomic.Int64{}}
}

type WorkFinder struct {
	Gateway DataGateway
	Results atomic.Int64
	Panics  uint64
}

func (f *WorkFinder) MarkErroneous(task Task) {
	f.Results.Add(1)
}

func (f *WorkFinder) MarkCompleted(task Task) {
	f.Results.Add(1)
}

func (f *WorkFinder) Stop() {
}

func (f *WorkFinder) FindRequested() []Task {
	gateway := f.Gateway
	records, _ := gateway.Find()
	s := make([]Task, len(records))
	for i, r := range records {
		s[i] = Task{r.Url}
	}
	return s
}
