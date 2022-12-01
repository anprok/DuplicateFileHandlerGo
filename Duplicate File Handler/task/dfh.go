package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	type File struct {
		Hash string
		Path string
	}
	var filesList map[int][]File
	filesList = make(map[int][]File)
	filesToDelete := make(map[int]string)

	if len(os.Args) != 2 {
		fmt.Println("Directory is not specified!")
		return
	}

	fmt.Println("Enter file format:")
	fileFormat := readString()

	err := filepath.Walk(os.Args[1], func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() && fileExists(path) && (fileFormat == "" || filepath.Ext(path) == "."+fileFormat) {
			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			md5Hash := md5.New()
			if _, err := io.Copy(md5Hash, file); err != nil {
				log.Fatal(err)
			}
			hash := string(md5Hash.Sum(nil))
			filesList[int(info.Size())] = append(filesList[int(info.Size())], File{
				Hash: hash,
				Path: path,
			})

		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	sizes := make([]int, 0, len(filesList))
	for key := range filesList {
		sizes = append(sizes, key)
	}
	fmt.Println("Size sorting option:\n1. Descending\n2. Ascending")
	for {
		fmt.Println("\nEnter a sorting option:")
		sortOption := readString()
		if sortOption == "1" { // Descending
			sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
		} else if sortOption == "2" { // Ascending
			sort.Ints(sizes)
		} else {
			fmt.Println("Wrong option")
			continue
		}
		for _, size := range sizes {
			fmt.Println()
			fmt.Println(size, "bytes")
			for _, s := range filesList[size] {
				fmt.Println(s.Path)
			}
		}
		break
	}
	for {
		fmt.Println("\nCheck for duplicates?")
		answer := readString()
		if answer == "yes" {
			fmt.Println()
			var lineNumber int
			lineNumber = 1
			for _, size := range sizes {

				if len(filesList[size]) > 1 {
					fmt.Println(size, "bytes")
					hashes := make(map[string][]string)
					for _, s := range filesList[size] {
						hashes[s.Hash] = append(hashes[s.Hash], s.Path)
					}
					for hashVal, i := range hashes {
						if len(i) > 1 {
							fmt.Printf("Hash: %x\n", hashVal)
							for _, s := range i {
								fmt.Printf("%d. %v\n", lineNumber, s)
								filesToDelete[lineNumber] = s
								lineNumber++
							}
						}
					}
				}
			}
			break
		} else if answer == "no" {
			break
		} else {
			fmt.Println("Wrong option")
			continue
		}
	}

	for {
		fmt.Println("\nDelete files?")
		answer := readString()
		if answer == "yes" {
			for {
				fmt.Println("Enter file numbers to delete:")
				numbersToDelete := readString()
				newSlice := make([]string, 0, len(filesToDelete))
				contFlag := false
				for _, k := range strings.Split(numbersToDelete, " ") {
					fileNum, _ := strconv.Atoi(k)
					if val, ok := filesToDelete[fileNum]; ok {
						newSlice = append(newSlice, val)
					} else {
						fmt.Println("Wrong format")
						contFlag = true
						break
					}
				}
				if contFlag {
					continue
				}
				var freeTotal int64 = 0

				for _, s := range newSlice {
					fileInfo, err := os.Stat(s)
					if err != nil {
						log.Fatal(err)
					}
					freeTotal += fileInfo.Size()
					err = os.Remove(s)
					if err != nil {
						log.Fatal(err)
					}
				}
				fmt.Printf("Total freed up space: %d bytes", freeTotal)
				break

			}
		} else if answer == "no" {
			break
		} else {
			fmt.Println("Wrong option")
			continue
		}
	}
}

func readString() string {
	var userInput string
	reader := bufio.NewReader(os.Stdin)
	userInput, _ = reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
