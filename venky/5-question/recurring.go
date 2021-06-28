package main

import "fmt"

// given an array =[2, 5, 1, 2, 3, 5, 1, 2, 4 ] => it should return 2
// given an array =[2, 1, 1, 2, 3, 5, 1, 2, 4 ] => it should return 1
// given an array =[2, 5, 1, 4 ] => it should return nil

var (
	arr1 = []int{2, 5, 1, 2, 3, 5, 1, 2, 4}
	arr2 = []int{2, 1, 1, 2, 3, 5, 1, 2, 4}
	arr3 = []int{2, 5, 1, 4}
)

func main() {
	//fmt.Println(recurring2(arr1))
	//fmt.Println(recurring2(arr2))
	//fmt.Println(recurring2(arr3))

	//fmt.Println(recurringWithTwoLoop(arr1))
	//fmt.Println(recurringWithTwoLoop(arr2))
	//fmt.Println(recurringWithTwoLoop(arr3))

	fmt.Println(recurringWithTwoLoop(arr1))
	fmt.Println(recurringWithTwoLoop(arr2))
	fmt.Println(recurringWithTwoLoop(arr3))
}

// given an array =[2, 5, 1, 2, 3, 5, 1, 2, 4 ] => it should return 2
// given an array =[2, 1, 1, 2, 3, 5, 1, 2, 4 ] => it should return 1
// given an array =[2, 5, 1, 4 ] => it should return nil

// best solution with map
func recurringWithMap(arr []int) int {
	table := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		if _, ok := table[arr[i]]; ok {
			return arr[i]
		}
		table[arr[i]] = arr[i]
	}
	return -1
}

// solution with two loop
func recurringWithTwoLoop(arr []int) int {

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[i] == arr[j] {
				return arr[i]
			}
		}
	}
	return -1
}

func recurring3(arr []int) int {
	existedArr := []int{}

	for i := 0; i < len(arr); i++ {
		if existedArr[arr[i]] != 0 {
			return arr[i]
		}
		existedArr[arr[i]] = arr[i]
	}
	return -1
}
