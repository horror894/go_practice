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

func checkUserInput(word string, tabooWords *[]string) string {
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
	if _, err := fmt.Scan(&userInput); err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Scan(&userInput)
		if userInput == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		checkWord := checkUserInput(userInput, &tabooWords)

		fmt.Println(checkWord)

	}
}
