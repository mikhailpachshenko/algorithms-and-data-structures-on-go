package main

import (
	"fmt"
	"io"
)

func bubbleSortingFlagsI(list []int) {
	for i := range list {
		swapped := false
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func bubbleSortingFlagsII(list []int) {
	swapped, i := true, 0
	for swapped {
		swapped = false
		for j := 0; j < len(list)-i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				swapped = true
			}
		}
		i++
	}
}

func main() {
	var l, count int
	for {
		_, err := fmt.Scan(&l)
		if err == io.EOF {
			return
		}

		list := make([]int, l)
		for i := range list {
			fmt.Scan(&list[i])
		}

		if count == 0 {
			bubbleSortingFlagsI(list)
			fmt.Println(list)
			count++
		} else {
			bubbleSortingFlagsII(list)
			fmt.Println(list)
			count--
		}
	}
}
