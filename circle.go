package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

const (
	span = 1024 * 1024 * 128
)

type Link struct {
	dummy [span]byte
	next  *Link
}

// create a loop and return it
func circle() *Link {
	one := new(Link)
	root := one

	for i := 0; i < 100; i++ {
		two := new(Link)
		one.next = two
		one = two
	}
	one.next = root
	return root
}

// create a loop and don't return it
func line() {
	one := new(Link)

	for i := 0; i < 100; i++ {
		two := new(Link)
		one.next = two
		one = two
	}
}

// don't create a loop
func doughnut() {
	one := new(Link)
	root := one

	for i := 0; i < 100; i++ {
		two := new(Link)
		one.next = two
		one = two
	}
	one.next = root
}

// create a chain and return it
func arrow() *Link {
	one := new(Link)
	root := one

	for i := 0; i < 100; i++ {
		two := new(Link)
		one.next = two
		one = two
	}
	// one.next = root
	return root
}

// just a sequence of disconected allocations
func dots() {
	one := new(Link)

	for i := 0; i < 100; i++ {
		two := new(Link)
		_ = two
	}
	_ = one
}

func PrintMemUsage(prefix string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	if len(prefix) > 0 {
		fmt.Printf("    %s\t", prefix)
	}

	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func main() {
	appstart := time.Now()
	task := "circle"
	if len(os.Args) > 1 {
		task = os.Args[1]
	}

	PrintMemUsage("at startup: " + task)
	for i := 0; i < 12; i++ {
		start := time.Now()

		switch task {
		case "circle":
			circle()
		case "line":
			line()
		case "doughnut":
			doughnut()
		case "arrow":
			arrow()
		case "dots":
			dots()
		default:
			fmt.Println("", task)
			return
		}

		took := time.Since(start).Microseconds()
		msg := fmt.Sprintf("%v (%v ms)", i, took)
		// PrintMemUsage(msg)
		_ = msg
	}

	PrintMemUsage("at the end")
	runtime.GC()
	PrintMemUsage("after gc")

	fmt.Printf("app run:%v\n", time.Since(appstart))
}
