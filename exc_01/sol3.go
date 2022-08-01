package main

import (
	"fmt"
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		currentSum := array[left] + array[right]
		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

func main() {
	var array = []int{3, 5, -4, 8, 11, 9, -1, 2}
	var target = 20

	fmt.Println(TwoNumberSum(array, target))

}
