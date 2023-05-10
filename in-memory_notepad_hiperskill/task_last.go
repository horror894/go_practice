package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 	write your code here

	var line string
	var noteCapasity uint
	wordScanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the maximum number of notes:")
	fmt.Scan(&noteCapasity) // add validation

	var notes = make([]string, note_capasity)

	for {
		fmt.Print("Enter a command and data:")

		wordScanner.Scan()
		//err := wordScanner.Err()
		//if err != nil {
		//	fmt.Println("'Scanln' error:", err)
		//	return
		//}
		line = wordScanner.Text()
		splitText := strings.SplitN(line, " ", 2)

		switch splitText[0] {
		case "exit":
			fmt.Println("[Info] Bye!\\n")
			os.Exit(0)
		case "create":
			if len(splitText) <= 1 {
				fmt.Println("[Error] Missing note argument")
			} else {
				for index, element := range notes { // complexity can be constant but my is n
					if element == "" {
						notes[index] = splitText[1]
						fmt.Println("[OK] The note was successfully created")
						break
					} else if index+1 == len(notes) {
						fmt.Println("[Error] Notepad is full\\n")
					}
				}
			}

		case "list":
			i := 0
			for index, element := range notes {
				if element != "" {
					fmt.Printf("[Info] %d: %s\n", index+1, element)
					i++
				} else {
				}
			}
			if i == 0 {
				fmt.Println("[Info] Notepad is empty")
			}
		case "clear":
			notes = []string{}
			fmt.Println("[OK] All notes were successfully deleted")

		default:
			fmt.Println("[Error] Unknown command")

		}
	}
}
