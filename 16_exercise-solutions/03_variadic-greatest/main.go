package main

import "fmt"

func greatest(si ...int) int {
	max := 0
	for i, n := range si {
		if i == 0 || n > max {
			max = n
		}
	}

	return max
}

func main() {
	max := greatest(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(max)
}
