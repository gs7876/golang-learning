package main

import "fmt"

func FourNumberSum(array []int, target int) [][]int {

	results := make([][]int, 0)
	doubleSum := make(map[int][][]int)

	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if values, ok := doubleSum[target-(array[i]+array[j])]; ok {

				for _, each := range values {
					results = append(results, []int{each[0], each[1], array[i], array[j]})
				}
			}
		}

		for j := i - 1; j >= 0; j-- {
			doubleSum[array[i]+array[j]] = append(doubleSum[array[i]+array[j]], []int{array[i], array[j]})
		}
	}
	return results
}

func main() {

	inputArray := []int{7, 6, 4, -1, 1, 2, 5, 8}
	targetSum := 16

	fmt.Println(FourNumberSum(inputArray, targetSum))

}
