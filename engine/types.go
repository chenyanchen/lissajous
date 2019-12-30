// Revision history:
//     Init: 2019/12/6    Jon Snow

package engine

import (
	"image/jpeg"
)

type Task struct {
	// source of file
	Src string
	// destination of new file
	Dst string

	JpegOpt *jpeg.Options

	WorkerFunc func(src, dst string, opt *jpeg.Options) (Result, error)
}

type Result struct {
	// new tasks
	Tasks []Task

	// results
	Items []interface{}
}
