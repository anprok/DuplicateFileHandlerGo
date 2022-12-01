package main

import (
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// DO NOT delete this code block! Mr. Krabs wrote it to open the "secret_formula.txt" file.
	file, err := os.Open("secret_formula.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use the correct function to create a 'new' SHA-512 Hash interface below:
	sha512Hash := sha512.New()

	// Copy the data from 'secret_formula.txt' to the 'sha512Hash'
	// Using the correct function from the `io` package:
	if _, err := io.Copy(sha512Hash, file); err != nil {
		log.Fatal(err)
	}

	// Call the method that returns the computed SHA-512 hash slice of the file:
	// And print the hash in hexadecimal notation below:
	fmt.Printf("%x\n", sha512Hash.Sum(nil))
}
