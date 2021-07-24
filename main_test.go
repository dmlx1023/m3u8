package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"m3u8/dl"
	"os"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	start:=time.Now()
	defer func() {
		fmt.Printf("total cost:%vs ",time.Since(start).Seconds())
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
			os.Exit(-1)
		}
	}()
	output="f:\\优质资源"
	url="https://video.dious.cc/20200612/tVmwUYdV/index.m3u8"
	videoName="是小甜心吗"
	chanSize=25
	downloader, err := dl.NewTask(output, url,videoName)
	if err != nil {
		t.Fatal("aaa")
		panic(err)
	}
	if err := downloader.Start(chanSize); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func Test_bar(t *testing.T)  {
	bar := progressbar.Default(100)
	bar.Describe("merging...")
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(100 * time.Millisecond)
	}
}
