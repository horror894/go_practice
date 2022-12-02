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
	var notes = make([]string, 5)

	wordScanner := bufio.NewScanner(os.Stdin)

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
			for index, element := range notes {
				if element == "" {
					notes[index] = splitText[1]
					fmt.Println("[OK] The note was successfully created")
					break
				} else if index+1 == len(notes) {
					fmt.Println("[Error] Notepad is full\\n")
				}
			}

		case "list":
			for index, element := range notes {
				if element != "" {
					fmt.Printf("[Info] %d: %s\n", index+1, element)
				} else {

				}

			}
		case "clear":
			notes = []string{}
			fmt.Println("[OK] All notes were successfully deleted")

		default:
			fmt.Println(splitText[0])

		}
	}
}

