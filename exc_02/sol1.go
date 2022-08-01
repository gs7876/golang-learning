package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0

	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			fmt.Println(array[arrIdx], sequence[seqIdx])
			seqIdx += 1
		}
		arrIdx += 1
	}
	// Write your code here.
	return seqIdx == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 5, 1, 8, 9, 32, 45}

	fmt.Println(IsValidSubsequence(array, sequence))

}
