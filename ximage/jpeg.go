// Copyright 2020 Singularity, Inc. All rights reserved.

package ximage

import (
	"image"
	"image/jpeg"
	"os"
)

func SaveAsFile(path string, img image.Image, opt *jpeg.Options) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return jpeg.Encode(file, img, opt)
}
