// Revision history:
//     Init: 2019/12/1    Jon Snow

package main

import (
	"image/jpeg"

	"github.com/JonSnow47/lissajous/engine"
)

func main() {
	e := engine.New(&engine.Options{
		Source:  source,
		Output:  output,
		JpegOpt: &jpeg.Options{Quality: quality},
	})

	e.Run()
}
