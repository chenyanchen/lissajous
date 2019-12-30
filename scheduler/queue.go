// Revision history:
//     Init: 2019/12/29    Jon Snow

package scheduler

import (
	"lissajous/engine"
)

type Queue struct {
	taskChan   chan engine.Task
	workerChan chan chan engine.Task
}

func (q *Queue) WorkerChan() chan engine.Task {
	return make(chan engine.Task)
}

func (q *Queue) WorkerReady(c chan engine.Task) {
	q.workerChan <- c
}

func (q *Queue) Submit(t engine.Task) {
	q.taskChan <- t
}

func (q *Queue) Run() {
	q.taskChan = make(chan engine.Task)
	q.workerChan = make(chan chan engine.Task)

	go func() {
		var (
			taskQueue   []engine.Task
			workerQueue []chan engine.Task
		)

		for {
			var (
				curTask   engine.Task
				curWorker chan engine.Task
			)

			if len(taskQueue) > 0 && len(curWorker) > 0 {
				curTask = taskQueue[0]
				curWorker = workerQueue[0]
			}

			select {
			case t := <-q.taskChan:
				taskQueue = append(taskQueue, t)
			case w := <-q.workerChan:
				workerQueue = append(workerQueue, w)
			case curWorker <- curTask:
				taskQueue = taskQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
