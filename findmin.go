package main

import "fmt"

func findMin(n []int) {
	var current_min = n[0]

	for i := 1; i < len(n); i++ {
		if (n[i] < current_min) {
			current_min = n[i]
		}
	}
	fmt.Println(current_min)
}

func main() {
	findMin([]int {1,2,3})
	findMin([]int {-10, 42, 18, 37, 99, -100})
	findMin([]int {-1,-2,-3})
}