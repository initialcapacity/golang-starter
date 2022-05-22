package datacollector

import (
	"io"
	"log"
	"net/http"
	"sync/atomic"
)

type Task struct {
	Url string
}

type Record struct {
	Url string
}

type DataGateway struct {
	Urls []string
}

func (d *DataGateway) Find() []Record {
	var tasks []Record
	for _, u := range d.Urls {
		tasks = append(tasks, Record{u})
	}
	return tasks
}

type Worker[T Task] struct {
}

func (a *Worker[T]) Run(w Task) error {
	get, err := http.Get(w.Url)
	if err != nil {
		return err
	}
	_, err = io.ReadAll(get.Body)
	return err
}

func NewWorkFinder[T Task](gateway DataGateway, c chan bool) WorkFinder[T] {
	return WorkFinder[T]{Gateway: gateway, Results: c}
}

type WorkFinder[T Task] struct {
	Gateway interface{}
	Results chan bool
	Panics  uint64
}

func (f *WorkFinder[T]) MarkErroneous(task T) {
	defer func() {
		if err := recover(); err != nil {
			atomic.AddUint64(&f.Panics, 1)
			log.Println("recovering collector work.")
		}
	}()
	f.Results <- false
}

func (f *WorkFinder[T]) MarkCompleted(task T) {
	defer func() {
		if err := recover(); err != nil {
			atomic.AddUint64(&f.Panics, 1)
			log.Println("recovering collector work.")
		}
	}()
	f.Results <- true
}

func (f *WorkFinder[T]) Stop() {
	close(f.Results)
}

func (f *WorkFinder[T]) FindRequested() []Task {
	gateway := f.Gateway.(DataGateway)
	results := gateway.Find()
	s := make([]Task, len(results))
	for i, r := range results {
		s[i] = Task{r.Url}
	}
	return s
}
