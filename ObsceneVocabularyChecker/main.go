package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var userInput string
	var wordForOutput string
	var tabooWords []string
	var matchStatus bool
	fmt.Scan(&userInput)

	file, err := os.Open(userInput)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScaner := bufio.NewScanner(file)
	for fileScaner.Scan() {
		tabooWords = append(tabooWords, fileScaner.Text())
	}

	for {
		wordForOutput = ""
		matchStatus = false
		fmt.Scan(&userInput)
		if userInput == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		for _, element := range tabooWords {
			if strings.ToLower(element) == strings.ToLower(userInput) {
				for range userInput {
					wordForOutput += "*"
				}
				matchStatus = true
				break
			}
		}

		if matchStatus {
			fmt.Println(wordForOutput)
		} else {
			fmt.Println(userInput)
		}

	}
}
