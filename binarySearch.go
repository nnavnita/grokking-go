package main

import "fmt"

func binarySearch(value int, items []int) bool {
	low := 0
	high := len(items) - 1

	for low <= high {
		mid := (low + high) / 2
		if items[mid] == value {
			return true
		}
		if items[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	items := []int{2, 3, 4, 5, 10, 20, 35, 46, 77, 91}
	fmt.Println(binarySearch(4, items))
	fmt.Println(binarySearch(50, items))
}
