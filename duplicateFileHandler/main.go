package main

import (
	"bufio"
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type sizeMap map[int][]string
type hashMap map[int]map[string][]string

func twoTypeSliceSort(sliceForSort []int, sortType bool) {
	sort.SliceStable(sliceForSort, func(l, j int) bool {
		if sortType {
			return sliceForSort[l] > sliceForSort[j]
		} else {
			return sliceForSort[l] < sliceForSort[j]
		}
	})
}

func getHashStringForFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}

func receiveFormatFromUser() string {
	var inputFileFormant string
	fmt.Println("Enter file format:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputFileFormant = scanner.Text()

	return "." + inputFileFormant
}

func receiveMapOfFiles(myPath string, filterFileFormat string) (map[int][]string, error) {
	myMap := make(map[int][]string)
	err := filepath.Walk(myPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			if filepath.Ext(path) == filterFileFormat || filterFileFormat == "." {
				key := int(info.Size())
				myMap[key] = append(myMap[key], path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return myMap, err
}

func receiveSortTypeFromUser() bool {
	var sortType string
	fmt.Print("\nSize sorting options:\n" +
		"1. Descending\n" +
		"2. Ascending\n" +
		"\n" +
		"Enter a sorting option:\n")
	myScanner := bufio.NewScanner(os.Stdin)
	myScanner.Scan()
	sortType = myScanner.Text()
	switch sortType {
	case "1":
		return true
	case "2":
		return false
	default:
		fmt.Println("Wrong option")
		return receiveSortTypeFromUser()
		// May be it's not good use recursion when it's not need. Replace it by loop.
	}
}

func printMap(myMap map[int][]string, sortedKeys []int) {
	for _, value := range sortedKeys {
		fmt.Printf("%d bytes\n", value)
		for _, val := range myMap[value] {
			fmt.Println(val)
		}
	}
}

func checkDuplicates(myMap map[int][]string) map[int]map[string][]string {
	tempHashDict := make(map[int]map[string][]string)
	for value, _ := range myMap {
		for _, val := range myMap[value] {
			stringHash, err := getHashStringForFile(val)
			if err != nil {
				fmt.Print(err)
			}
			if _, ok := tempHashDict[value]; !ok {
				tempHashDict[value] = make(map[string][]string)
			}
			tempHashDict[value][stringHash] = append(tempHashDict[value][stringHash], val)
		}
	}
	return tempHashDict
}

func printNewMap(sizeMap map[int]map[string][]string, sortedKeys []int) []string {
	lineNumber := 1
	var pathSlice []string

	for _, key := range sortedKeys {
		hashMap, _ := sizeMap[key]
		duplicateExist := false
		for _, path := range hashMap {
			if len(path) > 1 {
				duplicateExist = true
				break
			}
		}
		if duplicateExist {
			fmt.Printf("%d bytes\n", key)
			for hash, paths := range hashMap {
				if len(paths) > 1 {
					fmt.Printf("Hash: %v\n", hash)
					for _, path := range paths {
						fmt.Printf("%d. %v\n", lineNumber, path)
						pathSlice = append(pathSlice, path)
						lineNumber++
					}
				}
			}
		}
	}
	return pathSlice
}

func askUserAboutDeleting() bool {
	var userInput string
	fmt.Println("Delete files?")
	fmt.Scan(&userInput)

	switch userInput {
	case "yes":
		return true
	case "no":
		return false
	default:
		fmt.Println("Wrong option")
	}
	return false
}

func askWhatFilesWeWillDelete() []string {
	fmt.Println("Enter file numbers to delete:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	userInput := scanner.Text()

	numbersfordeleting := strings.Split(userInput, " ")

	return numbersfordeleting
}

func checkNumbersExists(listOfNumbers []string, pathForDelete []string) (bool, error) {
	borderNum := len(pathForDelete)
	for _, number := range listOfNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return false, err
		}
		if num > borderNum || num < 0 {
			return false, errors.New("wrong format")
		}
	}
	return true, nil
}

func deleteFiles(listOfNumbers []string, pathForDelete []string) (bool, int64) {
	var freedUpSpace int64
	for _, number := range listOfNumbers {
		num, err := strconv.Atoi(number)
		num--
		path := pathForDelete[num]
		fileInfo, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}
		freedUpSpace += fileInfo.Size()
		error := os.Remove(path)
		if error != nil {
			log.Fatal(error)
		}
	}
	return true, freedUpSpace
}

func main() {

	if len(os.Args) < 2 {
		fmt.Print("Directory is not specified")
	} else {
		inputFileFormant := receiveFormatFromUser()

		// Scan directory and save path for filtered files to map
		myMap, err := receiveMapOfFiles(os.Args[1], inputFileFormant)
		if err != nil {
			log.Fatal(err)
		}
		// Ask user about sort type
		sortType := receiveSortTypeFromUser()
		// I think it's not cool take string and convert it to bool, it will be not obvious in the future
		var keyList []int
		for k := range myMap {
			keyList = append(keyList, k)
		}
		// Sort slice in a two-way
		twoTypeSliceSort(keyList, sortType)
		// print map in sorted order
		printMap(myMap, keyList)

	myLoop2:
		for {
			fmt.Print("\nCheck for duplicates?\n")
			var ifWeCheckDuplicates string
			myScanner := bufio.NewScanner(os.Stdin)
			myScanner.Scan()
			ifWeCheckDuplicates = myScanner.Text()
			switch ifWeCheckDuplicates {
			case "yes":
				newMap := checkDuplicates(myMap)
				pathsForDelete := printNewMap(newMap, keyList)
				userChoice := askUserAboutDeleting()

				if userChoice {
					numbersForDeleting := askWhatFilesWeWillDelete()
					status, _ := checkNumbersExists(numbersForDeleting, pathsForDelete)
					if status {
						stat, size := deleteFiles(numbersForDeleting, pathsForDelete)
						if stat {
							fmt.Printf("Total freed up space: %d bytes", size)
						}
					}
					fmt.Println("Wrong format")
				}

				break myLoop2
			case "no":
				break myLoop2
			default:
				fmt.Println("Wrong option")
			}
		}

	}
}

// https://hyperskill.org/learn/step/18567
// https://gosamples.dev/read-user-input/
// read about input - I suppose that use wrong func for user input of user numbers
// check that user write numbers
// convert strings to numners
// check that numbers exist in slice
// n-1 use as index for deleting
// take path receice size add in variable and than delete file
// create slice when print
// return this slice
// read input
// check input
// delete files by numbers and summarize size and print
// for the future research, how I can use closure if I make separate func for filepath.Walk
// chat Gpt show good example how function return two values string and _, review it and implement
// make decomposition for main()
