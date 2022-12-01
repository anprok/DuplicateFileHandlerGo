package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string

	fmt.Scan(&source)

	// Your code here
	if strings.Count(source, "0") > strings.Count(source, "1") {
		fmt.Println("The count of zeros is bigger")
	} else if strings.Count(source, "0") < strings.Count(source, "1") {
		fmt.Println("The count of units is bigger")
	} else {
		fmt.Println("Counts are equal")

	}
}
