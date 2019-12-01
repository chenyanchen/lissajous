// Revision history:
//     Init: 2019/12/1    Jon Snow

package engine

import (
	"image"
	"image/jpeg"
	"os"
)

const JPEG = "jpeg"

func JPEGConverter(output string, img image.Image, opt *jpeg.Options) error {
	newFile, err := os.Create(output)
	if err != nil {
		return err
	}
	defer newFile.Close()

	return jpeg.Encode(newFile, img, opt)
}
