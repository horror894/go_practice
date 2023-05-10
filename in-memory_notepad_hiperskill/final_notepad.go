package main

// put all logic from switch case block inside functions
// I need to create to variable user position and offset position

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type notepad struct {
	note            []string
	indexFreeRecord int
	size            int
}

func initializeNote() notepad {

	var s string // why I check type if type of variable int !!!
	fmt.Print("Enter the maximum number of notes:")
	fmt.Scan(&s)
	if size, err := strconv.ParseInt(s, 10, 32); err == nil && size > 0 {
		notepadObj := notepad{note: make([]string, size), indexFreeRecord: 0, size: int(size)}
		for index, _ := range notepadObj.note {
			notepadObj.note[index] = ""
		}
		return notepadObj
	} else {
		fmt.Printf("[Error] Wrong size for note '%s'/n", s)
		os.Exit(0)
	}
	return notepad{}
}

func readStdin() []string { // question disscuss with someone

	var userLine string
	fmt.Print("Enter a command and data:")

	wordScanner := bufio.NewScanner(os.Stdin)
	wordScanner.Scan()
	userLine = wordScanner.Text()

	parametersSlice := strings.Fields(userLine)
	if len(parametersSlice) == 0 {
		parametersSlice = []string{""}
	}

	return parametersSlice
}

func gracefullExit() {
	fmt.Print("[Info] Bye!\n")
	os.Exit(0)
}

func default_mesage() {
	fmt.Print("[Error] Unknown command\n")
}

func create(s *notepad, e []string) {
	if len(e) <= 1 {
		fmt.Print("[Error] Missing note argument\n")
	} else {
		noteString := strings.Join(e[1:], " ")
		if s.indexFreeRecord >= s.size { // we decide that notepad full when indexFreeRecord eq to size
			fmt.Print("[Error] Notepad is full\n")
		} else {
			s.note[s.indexFreeRecord] = noteString
			s.indexFreeRecord++
			fmt.Print("[OK] The note was successfully created\n")
		}
	}

}

func printNoteList(s *notepad) {

	if s.indexFreeRecord == 0 {
		fmt.Print("[Info] Notepad is empty\n")
		return
	}
	for index, element := range s.note {
		if element != "" {
			fmt.Printf("[Info] %d: %s\n", index+1, element)
		}
	}
}

func clearNote(s *notepad) {
	s.note = []string{}
	s.indexFreeRecord = 0
	fmt.Print("[OK] All notes were successfully deleted\n")
}

func updateNotePosition(s *notepad, e []string) {
	// expected format "update <position> <text of note>"
	switch {
	case len(e) == 1:
		fmt.Print("[Error] Missing position argument\n")
	case len(e) == 2:
		fmt.Print("[Error] Missing note argument\n")
	case len(e) >= 3:
		if i, err := strconv.ParseInt(e[1], 10, 16); err == nil {
			noteString := strings.Join(e[2:], " ")
			position := int(i)
			offsetPosition := position - 1
			if position > s.size || position < 0 {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, s.size)
			} else if s.note[offsetPosition] == "" {
				fmt.Print("[Error] There is nothing to update\n")
			} else {
				s.note[offsetPosition] = noteString
				fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
			}
		} else if err != nil {
			fmt.Printf("[Error] Invalid position: %s\n", e[1])
		}
	}
}

func deleteNotePosition(s *notepad, e []string) {
	switch {
	case len(e) == 1:
		fmt.Print("[Error] Missing position argument\n")
	case len(e) == 2:
		if i, err := strconv.ParseInt(e[1], 10, 16); err == nil {
			position := int(i)
			offsetPosition := position - 1
			if position > s.size || position < 0 {
				fmt.Printf("[Error] Position %d is out of the boundaries [1, %d]\n", position, s.size)
			} else if s.note[offsetPosition] != "" {
				for r := offsetPosition; r < s.indexFreeRecord; r++ {
					s.note[offsetPosition] = s.note[offsetPosition+1]
					if r == s.indexFreeRecord-1 {
						s.indexFreeRecord = s.indexFreeRecord - 1
						s.note[s.indexFreeRecord] = ""
						break
					}
					// sort by replace
				}
				fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)
			} else {
				fmt.Print("[Error] There is nothing to delete\n")
			}
		} else if err != nil {
			fmt.Printf("[Error] Invalid position: %s\n", e[1])
		}
	}
}

func processingUserInput(s *notepad, e []string) {

	commad := e[0] // I'm not sure about it
	// add check that we receive valid string

	switch commad {
	case "exit":
		gracefullExit()
	case "create":
		create(s, e)
	case "list":
		printNoteList(s)
	case "clear":
		clearNote(s)
	case "update":
		updateNotePosition(s, e)
	case "delete":
		deleteNotePosition(s, e)
	default:
		default_mesage()

	} // all behavior must be in function I suppose, bat some actions about note struck could be methods of class note

}

func main() {

	myNote := initializeNote()

	// add loop for this
	for true {
		input := readStdin()
		processingUserInput(&myNote, input)
	}

}

