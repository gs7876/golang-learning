package main

import "fmt"

func TwoNumberSumSol1(array []int, target int) []int {

	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := 1 + i; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
		}
	}
	return []int{}
}

func main() {
	var array = []int{3, 5, -4, 8, 11, 9, -1, 2}
	var target = 20

	fmt.Println(TwoNumberSumSol1(array, target))

}
