package main

import "fmt"

func bubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func main() {
	slice := []int{5, 2, 4, 3, 1}
	fmt.Println(slice)
	bubbleSort(slice)
	fmt.Println(slice)
}
