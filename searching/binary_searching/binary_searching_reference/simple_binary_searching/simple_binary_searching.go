package main

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

func simpleBinarySearching(list []int, target int) int {
	min, max := 0, len(list)
	for min != max {
		mid := (min + max) / 2
		if target == list[mid] {
			return mid
		} else if target < list[mid] {
			max = mid
		} else {
			min = mid + 1
		}
	}
	return -1
}

func main() {
	var l int
	for {
		_, err := fmt.Scan(&l)
		if err == io.EOF {
			return
		}

		list := make([]int, l)
		for i := range list {
			fmt.Scan(&list[i])
		}

		for i, val := range list {
			fmt.Printf("Searching for %d, in list %v\n", val, list)
			fmt.Printf("Number was finded in position %d\n\n", i)
		}

		/* simplified speed test */
		testList := make([]int, 10000000)
		for i := range testList {
			testList[i] = i
		}

		start := time.Now()
		for i := 0; i < 100; i++ {
			target := rand.Intn(10000000)
			fmt.Println("target:", target)
			simpleBinarySearching(testList, target)
		}
		end := time.Since(start)
		nanoseconds := end.Nanoseconds() / 100
		fmt.Println("Average time:", nanoseconds)
	}
}
