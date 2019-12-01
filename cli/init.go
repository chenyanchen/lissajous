// Revision history:
//     Init: 2019/12/1    Jon Snow

package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"os"
	"runtime"
)

const (
	appName = "Lissajous"
	version = "0.0.1"
)

var (
	source  string
	output  string
	quality int
)

func usage() {
	fmt.Fprintf(os.Stderr, "%s 帮助您将 jpg, png, gif 格式的图片转为更小的 jpeg 图片\n", appName)
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", appName)
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage

	v := flag.Bool("v", false, "当前版本")
	e := flag.Bool("e", false, "运行环境")
	h := flag.Bool("h", false, "使用说明")

	flag.StringVar(&source, "d", ".", "指定需要转换的文件名称或文件夹")
	flag.Parse()

	flag.StringVar(&output, "o", source+"_converted", "输出文件夹")
	flag.IntVar(&quality, "q", jpeg.DefaultQuality, "输出图片的质量(0~100)")

	flag.Parse()

	if *v || *e || *h {
		if *v {
			fmt.Println(appName, "version:", version)
		}
		if *e {
			fmt.Printf("go version: %s\n GOOS: %s\n GRASH: %s\n",
				runtime.Version(), runtime.GOOS, runtime.GOARCH)
		}
		if *h {
			flag.Usage()
		}

		os.Exit(0)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())
}
