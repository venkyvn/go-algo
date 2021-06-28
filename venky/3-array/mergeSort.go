package main

import "fmt"

var (
	firstArr  = []int{}
	secondArr = []int{4, 6, 30, 33, 35}
)

func main() {
	fmt.Println("result", mergeSort1(firstArr, secondArr))

}

func mergeSort1(firstArr []int, secondArr []int) []int {

	totalLen := len(firstArr) + len(secondArr)
	resultArr := make([]int, totalLen)
	arr1 := 0
	arr2 := 0
	i := 0
	for i < totalLen {
		if arr2 >= len(secondArr) {
			resultArr[i] = firstArr[arr1]
			arr1++
		} else if arr1 >= len(firstArr) {
			resultArr[i] = secondArr[arr2]
			arr2++
		} else if firstArr[arr1] <= secondArr[arr2] {
			resultArr[i] = firstArr[arr1]
			arr1++
		} else {
			resultArr[i] = secondArr[arr2]
			arr2++
		}
		i++
	}

	return resultArr
}
