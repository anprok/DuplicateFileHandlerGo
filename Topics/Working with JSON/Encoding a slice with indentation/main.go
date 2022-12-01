package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// DO NOT delete the code block below, it populates the 'fruits' slice!
	var f1, f2, f3 string
	fmt.Scanln(&f1, &f2, &f3)
	fruits := []string{f1, f2, f3}

	// Remember to specify 4 blank spaces "    " as the indent parameter below:
	fruitsJSON, err := json.MarshalIndent(fruits, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(fruitsJSON))
}
