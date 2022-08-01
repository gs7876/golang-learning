package main

import "fmt"

func TwoNumberSumSol2(array []int, target int) []int {
	nums := map[int]bool{}

	for _, num := range array {

		fmt.Println(num)

		potentialMatch := target - num
		if found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}

	return []int{}

}

func main() {
	var array = []int{3, 5, -4, 8, 11, 9, -1, 2}
	// var array = []int{2, 18}
	var target = 20

	fmt.Println(TwoNumberSumSol2(array, target))

}
