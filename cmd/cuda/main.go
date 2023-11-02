// What it does:
//
// 	This program outputs the current OpenCV library version and CUDA version the console.
//
// How to run:
//
// 		go run --tags cuda ./cmd/cuda/main.go
//
//go:build cuda
// +build cuda

package main

import (
	"fmt"

	"github.com/xuxiaoshuo/gocv"
	"github.com/xuxiaoshuo/gocv/cuda"
)

func main() {
	fmt.Printf("gocv version: %s\n", gocv.Version())
	fmt.Println("cuda information:")
	devices := cuda.GetCudaEnabledDeviceCount()
	for i := 0; i < devices; i++ {
		fmt.Print("  ")
		cuda.PrintShortCudaDeviceInfo(i)
	}
}
