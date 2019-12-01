// Revision history:
//     Init: 2019/12/1    Jon Snow

package engine

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// read image from a file path
func ImageReader(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	img, _, err = image.Decode(file)
	return img, err
}
