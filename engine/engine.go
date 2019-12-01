// Revision history:
//     Init: 2019/12/1    Jon Snow

package engine

import (
	"image/jpeg"
	"log"
	"os"
	"strings"
)

type Engine struct {
	queue  []string
	source string
	output string

	jpegOpt *jpeg.Options
}

type Options struct {
	Source  string
	Output  string
	JpegOpt *jpeg.Options
}

func New(opt *Options) *Engine {
	e := &Engine{
		source: ".",
		output: ".output",
	}

	if opt != nil {
		e.source = opt.Source
		e.output = opt.Output

		if opt.JpegOpt != nil {
			e.jpegOpt = opt.JpegOpt
		}
	}

	e.queue = []string{e.source}

	return e
}

func (e *Engine) targetPath(path string) string {
	return strings.Replace(path, e.source, e.output, 1)
}

func (e *Engine) Run() {
	for len(e.queue) > 0 {
		cur := e.queue[0]
		e.queue = e.queue[1:]

		e.worker(cur)
	}
}

func (e *Engine) worker(path string) {
	info, err := os.Stat(path)
	if err != nil {
		log.Printf("[E] os.Stat(%s): %v", path, err)
		return
	}

	if info.IsDir() {
		files, err := DirOpener(path)
		if err != nil {
			log.Printf("[E] DirOpener(%s): %v", path, err)
			return
		}
		e.queue = append(e.queue, files...)

		dumpDir := e.targetPath(path)
		err = os.MkdirAll(dumpDir, os.ModePerm)
		if err != nil {
			log.Printf("[E] os.MkdirAll(%s, %d): %v", path, os.ModePerm, err)
			return
		}
	} else {
		img, err := ImageReader(path)
		if err != nil {
			log.Printf("[E] ImageReader(%s): %v", path, err)
			return
		}

		dstPath := e.targetPath(path)
		names := strings.Split(dstPath, ".")
		names[len(names)-1] = JPEG
		dstPath = strings.Join(names, ".")

		err = JPEGConverter(dstPath, img, e.jpegOpt)
		if err != nil {
			log.Printf("[E] JPEGConverter(%s): %v", path, err)
			return
		}
	}
}
