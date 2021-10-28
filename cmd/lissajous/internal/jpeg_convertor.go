// Copyright 2020 Singularity, Inc. All rights reserved.

package internal

import (
	"errors"
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/chenyanchen/lissajous/ximage"
)

type jpegConvertor struct {
	src string        // source path
	dst string        // destination path
	ext string        // file extension
	opt *jpeg.Options // jpeg options
}

func newJpegConvertor(src, dst string, quality int) (*jpegConvertor, error) {
	opt := &jpeg.Options{Quality: quality}
	srcAbs, err := filepath.Abs(src)
	if err != nil {
		return nil, err
	}
	dstAbs, err := filepath.Abs(dst)
	if err != nil {
		return nil, err
	}
	return &jpegConvertor{
		src: srcAbs,
		dst: dstAbs,
		ext: ".jpeg",
		opt: opt,
	}, nil
}

func NewJpegConvertor(src, dst string, quality int) (Convertor, error) {
	return newJpegConvertor(src, dst, quality)
}

func (c *jpegConvertor) Convert() error {
	return filepath.Walk(c.src, c.convert)
}

func (c *jpegConvertor) convert(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	dst := strings.Replace(path, c.src, c.dst, 1)
	if info.IsDir() {
		if path == c.dst {
			return filepath.SkipDir
		}
		return os.MkdirAll(dst, info.Mode())
	}

	img, _, err := ximage.OpenFromFile(path)
	if err != nil {
		if errors.Is(err, image.ErrFormat) {
			return nil
		}
		log.Printf("[E] cant read file %s as image, cause: %v", path, err)
		return nil
	}

	err = ximage.SaveAsFile(dst, img, c.opt)
	if err != nil {
		log.Printf("[E] failed to save image %s as file %s: %v", path, dst, err)
		return nil
	}

	log.Printf("[I] image %s convert as %s", path, dst)
	return nil
}
