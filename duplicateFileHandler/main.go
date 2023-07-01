package main

// for the future research, how I can use closure if I make separate func for filepath.Walk
// chat Gpt show good example how function return two values string and _, review it and implement
// make decomposition for main()

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
)

func twoTypeSliceSort(sliceForSort []int, sortType bool) {
	sort.Slice(sliceForSort, func(l, j int) bool {
		if sortType {
			return sliceForSort[l] > sliceForSort[j]
		} else {
			return sliceForSort[l] < sliceForSort[j]
		}
	})

}

func main() {
	var fileDict map[int][]string
	fileDict = make(map[int][]string)

	if len(os.Args) < 2 {
		fmt.Print("Directory is not specified")
	} else {
		var inputFileFormant string
		fmt.Println("Enter file format:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputFileFormant = scanner.Text()

		inputFileFormant = "." + inputFileFormant

		err := filepath.Walk(os.Args[1], func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			if !info.IsDir() {
				if filepath.Ext(path) == inputFileFormant || inputFileFormant == "." {
					key := int(info.Size())
					fileDict[key] = append(fileDict[key], path)
				}
			}
			return nil
		})
		if err != nil {
			log.Fatal(err)
		}
		keyList := []int{}

		for k := range fileDict {
			keyList = append(keyList, k)
		}
		var sortType string
		fmt.Print("Enter a sorting options:\n" +
			"1. Descending\n" +
			"2. Ascending\n" +
			"\n" +
			"Enter a sorting option:\n")

	myloop1:
		for {
			myscanner := bufio.NewScanner(os.Stdin)
			myscanner.Scan()
			sortType = myscanner.Text()
			switch sortType {
			case "1":
				twoTypeSliceSort(keyList, true)
				break myloop1
			case "2":
				twoTypeSliceSort(keyList, false)
				break myloop1
			default:
				fmt.Println("Wrong option")
			}
		}

		for _, value := range keyList {
			fmt.Printf("%d bytes\n", value)
			for _, val := range fileDict[value] {
				fmt.Println(val)
			}
			fmt.Print("\n")
		}

	myloop2:
		for {
			fmt.Print("Check for duplicates?\n")
			var ifWeCheckDuplicates string
			myscanner := bufio.NewScanner(os.Stdin)
			myscanner.Scan()
			ifWeCheckDuplicates = myscanner.Text()
			switch ifWeCheckDuplicates {
			case "yes":
				var tempHashDict = make(map[int]map[string][]string)
				for _, value := range keyList {
					for _, val := range fileDict[value] {
						file, err := os.Open(val)
						if err != nil {
							fmt.Println(err)
						}

						md5Hash := md5.New()
						if _, err := io.Copy(md5Hash, file); err != nil {
							fmt.Println(err)
						}
						stringHash := hex.EncodeToString(md5Hash.Sum(nil))
						if _, ok := tempHashDict[value]; !ok {
							tempHashDict[value] = make(map[string][]string)
						}
						tempHashDict[value][stringHash] = append(tempHashDict[value][stringHash], val)
						file.Close()
					}
				}

				lineNumber := 1
				for element := range tempHashDict {
					for element1 := range tempHashDict[element] {
						if len(tempHashDict[element][element1]) > 1 {
							fmt.Printf("%d bytes\n", element)
							fmt.Printf("Hash: %v\n", element1)
							for _, element2 := range tempHashDict[element][element1] {
								fmt.Printf("%d. %v\n", lineNumber, element2)
								lineNumber++
							}
							fmt.Println()
						}
					}
				}
				break myloop2
			case "no":
				break myloop2
			default:
				fmt.Println("Wrong option")
			}
		}
	}
	// write your code here
}
