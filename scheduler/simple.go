// Revision history:
//     Init: 2019/12/29    Jon Snow

package scheduler

import (
	"lissajous/engine"
)

type Simple struct {
	workChan chan engine.Task
}

func (s *Simple) WorkerChan() chan engine.Task {
	return s.workChan
}

func (s *Simple) WorkerReady(chan engine.Task) {}

func (s *Simple) Submit(t engine.Task) {
	go func() { s.workChan <- t }()
}

func (s *Simple) Run() {
	s.workChan = make(chan engine.Task)
}
