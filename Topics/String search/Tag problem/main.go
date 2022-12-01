package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string

	fmt.Scan(&source)

	openingTagStart := strings.Index(source, "<")
	openingTagEnd := strings.Index(source, ">")

	closingTagStart := strings.LastIndex(source, "<")
	closingTagEnd := strings.LastIndex(source, ">")

	fmt.Printf("%d %d\n", openingTagStart, openingTagEnd)
	fmt.Printf("%d %d\n", closingTagStart, closingTagEnd)
}
