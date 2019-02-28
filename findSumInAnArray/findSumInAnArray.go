package main

import "log"

// findSumInArray Checks if there are 2 numbers present in an array whose sum is a given value
func findSumInArray(arr []int, sum int) (int, int, bool) {
	sumMap := make(map[int]bool)
	size := len(arr)

	for i := 0; i < size; i++ {
		sumMap[arr[i]] = true
	}

	for i := 0; i < size; i++ {
		if sumMap[sum-arr[i]] {
			return arr[i], sum - arr[i], true
		}
	}
	return 0, 0, false
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sum := 10
	a, b, ok := findSumInArray(array, sum)
	if ok {
		log.Println("Sum Exist:", a, b)
	} else {
		log.Println("Sum doesnot exist")
	}
}
