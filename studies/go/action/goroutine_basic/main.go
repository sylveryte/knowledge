package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Sleeping for 10ms")
	buf()
}

func buf() {
	processConcurrently([]int{1, 2, 3, 4})
	time.Sleep(100 * time.Millisecond)
}

func processConcurrently(vals []int) {
	in := make(chan int)
	out := make(chan int)

	//load in
	go func() {
		// fmt.Println("Loading in")
		for _, v := range vals {
			fmt.Println("Loaded in", v)
			in <- v
		}
		fmt.Println("Closing in")
		close(in)
	}()
	go func() {
		for v := range in {
			out <- process(v)
		}
	}()

	for v := range out {
		fmt.Println("Processed ", v)
	}
	fmt.Println("Closing out")
	close(out)
}

func process(val int) int {
	// fmt.Println("Processing", val)
	return val * val
}
