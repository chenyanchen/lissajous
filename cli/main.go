// Revision history:
//     Init: 2019/12/1    Jon Snow

package main

import (
	"image/jpeg"
	_ "net/http"

	"lissajous/engine"
	"lissajous/scheduler"
	"lissajous/worker"
)

func main() {
	e := &engine.ConcurrentEngine{
		WorkerNum: 10,
		Scheduler: &scheduler.Simple{},
	}

	e.Run(engine.Task{
		Src:        source,
		Dst:        output,
		JpegOpt:    &jpeg.Options{Quality: quality},
		WorkerFunc: worker.Worker,
	})
}
