package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string

	fmt.Scan(&source)

	var hasSuffix string
	if strings.HasSuffix(source, "ing") {
		hasSuffix = "yes"
	} else {
		hasSuffix = "no"
	}
	fmt.Println(hasSuffix)
}
