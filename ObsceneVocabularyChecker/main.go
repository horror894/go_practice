package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func reciveTabooWords(file *os.File) []string {
	var tabooSlice []string
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		tabooSlice = append(tabooSlice, fileScanner.Text())
	}
	return tabooSlice
}

func checkWord(word string, tabooWords *[]string) string {
	var outPutWord string
	var isItbadWord bool
	for _, element := range *tabooWords {
		if strings.EqualFold(word, element) {
			isItbadWord = true
			break
		}
	}
	if isItbadWord {
		for range word {
			outPutWord += "*"
		}
	} else {
		outPutWord = word
	}

	return outPutWord
}

func main() {
	var userInput string
	var userInputSlice []string
	// user provide file name
	if _, err := fmt.Scan(&userInput); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(userInput)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tabooWords := reciveTabooWords(file)
	// user provide string for analize

	for {

		userInputScanner := bufio.NewScanner(os.Stdin)
		//userInputScanner.Split(bufio.ScanWords)
		userInputScanner.Scan()
		myInput := userInputScanner.Text()

		if strings.HasPrefix(myInput, "exit") {
			fmt.Println("Bye!")
		}

		userInputSlice = strings.Split(myInput, "\n")

		var finalString string
		for _, element := range userInputSlice {
			checkedWord := checkWord(element, &tabooWords)
			finalString += checkedWord
		}

		fmt.Println(finalString)

	}
}
