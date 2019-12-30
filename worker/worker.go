// Revision history:
//     Init: 2019/12/6    Jon Snow

package worker

import (
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"lissajous/engine"
)

func Worker(src, dst string, opt *jpeg.Options) (engine.Result, error) {
	var result engine.Result

	info, err := os.Stat(src)
	if err != nil {
		log.Printf("[E] os.Stat(%s): %v", src, err)
		return result, err
	}

	if info.IsDir() {
		files, err := DirOpener(src)
		if err != nil {
			log.Printf("[E] DirOpener(%s): %v", src, err)
			return result, err
		}

		result.Tasks = make([]engine.Task, 0, len(files))
		for _, fileName := range files {
			result.Tasks = append(result.Tasks, engine.Task{
				Src:        fileName,
				Dst:        dst,
				JpegOpt:    opt,
				WorkerFunc: Worker,
			})
		}

		dumpDst := filepath.Join(dst, src)
		err = os.MkdirAll(dumpDst, os.ModePerm)
		if err != nil {
			log.Printf("[E] os.MkdirAll(%s, %d): %v", dumpDst, os.ModePerm, err)
			return result, err
		}
	} else {
		img, err := ImageReader(src)
		if err != nil {
			log.Printf("[E] ImageReader(%s): %v", src, err)
			return result, err
		}

		dumpDst := filepath.Join(dst, src)
		names := strings.Split(dumpDst, ".")
		names[len(names)-1] = JPEG
		dumpDst = strings.Join(names, ".")

		err = JPEGConverter(dumpDst, img, opt)
		if err != nil {
			log.Printf("[E] JPEGConverter(%s): %v", src, err)
			return result, err
		}

		result.Items = append(result.Items, "converted image: "+src)
	}

	return result, err
}
