package main

import "fmt"

var ()

func main() {
	arr1 := []int{1, 2, 3, 4, 5}

	fmt.Println(findArraySum(arr1, 15))

}

func findArraySum(arr1 []int, sum int) string {
	pastArr := make(map[int]bool)

	for i := range arr1 {
		if pastArr[i] {
			return "founded"
		}
		pastArr[sum-i] = true
	}
	return "not found"
}
