// Copyright 2020 Singularity, Inc. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"log"
	_ "net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/JonSnow47/lissajous/cmd/lissajous/internal"
	"github.com/JonSnow47/lissajous/pkg/directory"
)

const (
	appName = "Lissajous"
	version = "1.0.0"
)

var (
	source      string
	destination string
	quality     int
)

func printVersion() {
	fmt.Println(appName, "version:", version)
}

func printEnvironment() {
	fmt.Printf("go version: %s\n GOOS: %s\n ARCH: %s\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `%s is an image convertor.
    Help you convert your images into jpeg(jpg), current supported formats is jpeg(jpg), png, gif.
`, appName)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", appName)
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage

	v := flag.Bool("v", false, "print lissajous version")
	e := flag.Bool("env", false, "print runtime environment")
	h := flag.Bool("h", false, "usage of lissajous")

	dir := filepath.Dir(os.Args[0])
	flag.StringVar(&source, "src", dir, "set the source directory or image file path")
	flag.StringVar(&destination, "dst", filepath.Join(dir, "tmp"), "set the destination directory path")
	flag.IntVar(&quality, "q", jpeg.DefaultQuality, "set the quality of the converted image, ranges from 1 to 100 inclusive, higher is better")

	flag.Parse()

	if *v {
		printVersion()
		os.Exit(0)
	}
	if *e {
		printEnvironment()
		os.Exit(0)
	}
	if *h {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	convertor, err := internal.NewJpegConvertor(source, destination, quality)
	if err != nil {
		log.Fatalln("failed to NewJpegConvertor:", err)
	}
	if err := convertor.Convert(); err != nil {
		log.Fatalln("failed to convert:", err)
	}
	if err := directory.RemoveEmpty(destination); err != nil {
		log.Fatalln("failed to clean empty directories:", err)
	}
}
