// Copyright 2020 Singularity, Inc. All rights reserved.

package ximage

import (
	"image"
	"os"
)

func OpenFromFile(path string) (image.Image, string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	return image.Decode(file)
}
