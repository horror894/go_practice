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

const errorDirectoryNotSpecified = "directory is not specified"
const errorToManyTry = "too many try"
const errorWrongOptionForSort = "wrong option for sort, only 1 or 2"

const errorWrongOption = "\nWrong option"
const fileFormatPrompt = "Enter file format:\n"
const sortOptionPrompt = "\nSize sorting options:\n" +
	"1. Descending\n" +
	"2. Ascending\n" +
	"\n" +
	"Enter a sorting option:\n"
const checkDuplicatesPrompt = "\nCheck for duplicates?\n"
const deleteFilesPrompt = "\nDelete files?\n"
const numbersForDeletingPrompt = "\nEnter file numbers to delete:\n"

type PathStorage struct {
	rootPath      string
	extFilter     string
	sortedKeys    []int
	pathMap       map[int]map[string][]string
	numberedPaths []string
	error         error
}

func NewStorage(rootPath string, ExtFilter string) *PathStorage {
	return &PathStorage{
		rootPath:  rootPath,
		extFilter: ExtFilter,
	}
}

func (f *PathStorage) Error() error {
	return f.error
}

func (f *PathStorage) loadStorage() {
	f.pathMap = make(map[int]map[string][]string)

	err := filepath.Walk(f.rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if !info.IsDir() {
			if filepath.Ext(path) == "."+f.extFilter || f.extFilter == "" {
				sizeKey := int(info.Size())
				hashKey, err2 := getHashStringForFile(path)
				if err2 != nil {
					log.Fatal(err2)
				}
				if _, ok := f.pathMap[sizeKey]; !ok {
					f.pathMap[sizeKey] = make(map[string][]string)
				}
				f.pathMap[sizeKey][hashKey] = append(f.pathMap[sizeKey][hashKey], path)
			}
		}
		return nil
	})
	if err != nil {
		f.error = err
	}

	var keyList []int
	for k := range f.pathMap {
		keyList = append(keyList, k)
	}
	f.sortedKeys = keyList
}

func (f *PathStorage) printStorageContent(withHash bool) {
	lineNumber := 1
	var pathSlice []string

	for _, key := range f.sortedKeys {
		hashMap, _ := f.pathMap[key]

		if !withHash {
			fmt.Printf("%d bytes\n", key)
			for _, paths := range hashMap {
				for _, path := range paths {
					fmt.Println(path)
				}
			}
			fmt.Println()
		} else {
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
	}
	f.numberedPaths = pathSlice
}

func (f *PathStorage) sortContent(sortType string) {
	if sortType == "1" || sortType == "2" {
		sort.Slice(f.sortedKeys, func(l, j int) bool {
			switch sortType {
			case "1":
				return f.sortedKeys[l] > f.sortedKeys[j]
			case "2":
				return f.sortedKeys[l] < f.sortedKeys[j]
			default:
				return true
			}
		})
	} else {
		f.error = errors.New(errorWrongOptionForSort)
	}
}

func (f *PathStorage) deleteFiles(listDeletingNumbers []string) (bool, int64, error) {
	var freedUpSpace int64
	borderNum := len(f.numberedPaths)
	for _, number := range listDeletingNumbers {
		num, err := strconv.Atoi(number)
		if err != nil {
			return false, 0, err
		}
		if num > borderNum || num < 0 {
			return false, 0, errors.New("wrong format")
		}
	}
	for _, number := range listDeletingNumbers {
		num, err := strconv.Atoi(number)
		num--
		path := f.numberedPaths[num]
		fileInfo, err := os.Stat(path)
		if err != nil {
			log.Fatal(err)
		}
		freedUpSpace += fileInfo.Size()
		err = os.Remove(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	return true, freedUpSpace, nil
}

func receiveArguments() (string, error) {
	if len(os.Args) < 2 {
		return "", errors.New(errorDirectoryNotSpecified)
	}
	return os.Args[1], nil
}

func userInputReceiver(prompt string, possibleOptions []string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 10; i++ {
		if _, err := fmt.Print(prompt); err != nil {
			return "", err
		}

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return "", err
		}

		userInput := scanner.Text()
		if err := scanner.Err(); err != nil {
			return "", err
		}

		for _, option := range possibleOptions {
			if option == userInput {
				return userInput, nil
			}
			if option == "any" {
				return userInput, nil
			}
		}
		fmt.Println(errorWrongOption)
	}

	return "", errors.New(errorToManyTry)
}

func getUniversalOptions() []string {
	return []string{"any"}
}

func getSortingOptions() []string {
	return []string{"1", "2"}
}

func getYesNoOptions() []string {
	return []string{"yes", "no"}
}

func getHashStringForFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	md5Hash := md5.New()
	if _, err := io.Copy(md5Hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", md5Hash.Sum(nil)), nil
}

func main() {

	rootPath, err := receiveArguments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		//log.Fatal(err)
	}

	formatFiler, err := userInputReceiver(fileFormatPrompt, getUniversalOptions())
	if err != nil {
		log.Fatal(err)
	}

	sortingOption, err := userInputReceiver(sortOptionPrompt, getSortingOptions())
	if err != nil {
		log.Fatal(err)
	}

	storage := NewStorage(rootPath, formatFiler)

	storage.loadStorage()
	if err := storage.Error(); err != nil {
		log.Fatal(err)
	}

	storage.sortContent(sortingOption)
	if err := storage.Error(); err != nil {
		log.Fatal(err)
	}

	storage.printStorageContent(false)

	duplicatesOption, err := userInputReceiver(checkDuplicatesPrompt, getYesNoOptions())
	if err != nil {
		log.Fatal(err)
	}

	switch duplicatesOption {
	case "yes":
		storage.printStorageContent(true)
	case "no":
		os.Exit(0)
	}

	deleteOption, err := userInputReceiver(deleteFilesPrompt, getYesNoOptions())
	if err != nil {
		log.Fatal(err)
	}

	switch deleteOption {
	case "yes":
		deleteNumbers, err := userInputReceiver(numbersForDeletingPrompt, getUniversalOptions())
		if err != nil {
			log.Fatal(err)
		}
		deleteNumbersSlice := strings.Split(deleteNumbers, " ")

		status, freedSpace, err := storage.deleteFiles(deleteNumbersSlice)
		if err != nil {
			log.Fatal(err)
		}
		if status {
			fmt.Printf("Total freed up space: %d bytes", freedSpace)
		}
	case "no":

	}
}

// If duplicated not exist - print prompt and exit
// https://hyperskill.org/learn/step/18567
// https://gosamples.dev/read-user-input/
// read about input - I suppose that use wrong func for user input of user numbers
// check that user write numbers
// convert strings to numners
// check that numbers exist in slice
// n-1 use as index for deleting
// take path receice size add in variable and then delete file
// create slice when print
// return this slice
// read input
// check input
// delete files by numbers and summarize size and print
// for the future research, how I can use closure if I make separate func for filepath.Walk
// chat Gpt show good example how function return two values string and _, review it and implement
// make decomposition for main()
