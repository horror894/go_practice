package main

// for the future research, how I can use closure if I make separate func for filepath.Walk

import (
	"bufio"
	"fmt"
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
		keyList := []int{}
		for k := range fileDict {
			keyList = append(keyList, k)
		}
		var sortType int
		fmt.Print("Enter a sorting options:\n" +
			"1. Descending\n" +
			"2. Ascending\n" +
			"\n" +
			"Enter a sorting option:\n")

		fmt.Scan(&sortType)

		switch sortType {
		case 1:
			twoTypeSliceSort(keyList, true)
		case 2:
			twoTypeSliceSort(keyList, false)
		default:
			fmt.Println("Wrong option")
		}

		if err != nil {
			log.Fatal(err)
		}

		for _, value := range keyList {
			fmt.Printf("%d bytes\n", value)
			for _, val := range fileDict[value] {
				fmt.Println(val)
			}
			fmt.Print("\n")
		}
	}
	// write your code here
}
