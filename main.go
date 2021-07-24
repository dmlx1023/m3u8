package main

import (
	"flag"
	"fmt"
	"m3u8/dl"
	"os"
)

var (
	url      string
	output   string
	videoName string
	chanSize int

)

func init() {
	flag.StringVar(&url, "u", "", "M3U8 URL, required.")
	flag.IntVar(&chanSize, "c", 25, "Maximum number of occurrences,default value is enough,don't change it if you are not understand this!")
	flag.StringVar(&output, "o", "f://videos", "Output folder, required")
	flag.StringVar(&videoName, "n", "m3u8_video", "File Name")
}

func main() {
	if len(flag.Args())==0 {
		fmt.Println("please use [-help] for more details!")
		return
	}
	flag.Parse()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	if url == "" {
		panicParameter("u")
	}
	if output == "" {
		panicParameter("o")
	}
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}
	downloader, err := dl.NewTask(output, url,videoName)
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize); err != nil {
		panic(err)
	}
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}
