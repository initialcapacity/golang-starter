/*
 * Copyright 2022 Hexa-org
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code from the Hexa project available at https://github.com/hexa-org/policy-orchestrator
// Inspired by code from the Hexa project available at https://github.com/Netflix/conductor

package workflowsupport

import (
	"log"
	"time"
)

type Worker[T any] interface {
	Run(task T) error
}

type WorkFinder[T any] interface {
	FindRequested() []T
	MarkCompleted(task T)
	MarkErroneous(task T)
	Stop()
}

type WorkScheduler[T any] struct {
	Finder  WorkFinder[T]
	Workers []Worker[T]
	Delay   int64
	done    chan bool
}

func NewScheduler[T any](finder WorkFinder[T], workers []Worker[T], delay int64) WorkScheduler[T] {
	return WorkScheduler[T]{
		Finder:  finder,
		Workers: workers,
		Delay:   delay,
		done:    make(chan bool),
	}
}

func (ws *WorkScheduler[T]) Start() {
	log.Printf("Starting the scheduler.\n")
	ticker := time.NewTicker(time.Duration(ws.Delay) * time.Millisecond)
	for _, w := range ws.Workers {
		go func(worker Worker[T]) {
			for {
				select {
				case <-ws.done:
					return
				case <-ticker.C:
					log.Printf("Scheduling work.\n")
					ws.checkForWork(worker)
				}
			}
		}(w)
	}
}

func (ws *WorkScheduler[T]) checkForWork(worker Worker[T]) {
	finder := ws.Finder
	log.Printf("Checking for work.\n")

	for _, r := range finder.FindRequested() {
		log.Printf("Found work.\n")

		go func(task T) {
			if err := worker.Run(task); err != nil {
				finder.MarkErroneous(task)
				return
			}
			log.Printf("Completed work.\n")
			finder.MarkCompleted(task)
		}(r)
	}
}

func (ws *WorkScheduler[T]) Stop() {
	ws.done <- true
	ws.Finder.Stop()
	log.Printf("Scheduler stopped.\n")
}
