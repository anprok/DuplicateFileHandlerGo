package main

import (
	"fmt"
	"strings"
)

func main() {
	algorithms := parseAlgorithms()
	fmt.Println(strings.Join(algorithms, ";")) // TODO: format algorithms
}

// DO NOT CHANGE
func parseAlgorithms() []string {
	var algorithms []string
	for {
		var algo string
		if _, err := fmt.Scan(&algo); err != nil {
			break
		}
		algorithms = append(algorithms, algo)
	}
	return algorithms
}
