package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// HINT: The error is within the 'medals' map declaration:
	medals := map[string]interface{}{
		"first":  "Gold",
		"second": "Silver",
		"third":  "Bronze",
	}

	medalsJSON, err := json.Marshal(medals)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(medalsJSON))
}
