package main

import "fmt"

var (
	nemos = []string{"nemo", "not nemo"}
)

func main() {
	findNemo(nemos)
}

func findNemo(array []string) {
	for _, nemo := range array {
		if nemo == "nemo" {
			fmt.Println("found nemo")
		}
	}
}
