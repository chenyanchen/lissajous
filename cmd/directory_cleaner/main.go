// Copyright 2020 Singularity, Inc. All rights reserved.

package main

import (
	"flag"
	"log"

	"github.com/JonSnow47/lissajous/pkg/directory"
)

var path string

func init() {
	flag.StringVar(&path, "p", ".", "root path")

	flag.Parse()
}

func main() {
	if err := directory.RemoveEmpty(path); err != nil {
		log.Fatalln("failed to clean directory:", err)
	}
}
