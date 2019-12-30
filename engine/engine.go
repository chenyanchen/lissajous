// Revision history:
//     Init: 2019/12/6    Jon Snow

package engine

import (
	"log"
)

type ReadyNotifier interface {
	WorkerReady(chan Task)
}

type Scheduler interface {
	// get worker from scheduler
	WorkerChan() chan Task
	// tell scheduler worker is ready
	ReadyNotifier
	// submit tasks to scheduler
	Submit(Task)
	// init scheduler and run
	Run()
}

type ConcurrentEngine struct {
	WorkerNum int
	Scheduler Scheduler
}

func (e *ConcurrentEngine) Run(seeds ...Task) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[E] %v", r)
		}
		log.Println("Finish")
	}()

	out := make(chan Result)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerNum; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, task := range seeds {
		e.Scheduler.Submit(task)
	}

	itemCounter := 0
	for result := range out {
		for _, task := range result.Tasks {
			e.Scheduler.Submit(task)
		}

		for _, item := range result.Items {
			itemCounter++
			log.Printf("[I] #%d %v\n", itemCounter, item)
		}
	}
}

func createWorker(in chan Task, out chan<- Result, notifier ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)

			task := <-in

			result, err := task.WorkerFunc(task.Src, task.Dst, task.JpegOpt)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
