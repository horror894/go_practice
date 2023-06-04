package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func regexpMatch(regexp, input string) bool {
	if regexp == "" {
		return true
	}
	if string(regexp[0]) == "^" {
		return regexpMatch(regexp[1:], input)
	}
	if input == "" && string(regexp[0]) == "$" {
		return true
	}
	if input == "" {
		return false
	}
	if len(regexp) >= 2 {
		if string(regexp[1]) == "?" {
			if regexp[0] == input[0] || string(regexp[0]) == "." {
				return regexpMatch(regexp[2:], input[1:])
			} else {
				return regexpMatch(regexp[2:], input)
			}
		}
		if string(regexp[1]) == "*" {
			if regexp[0] == input[0] || string(regexp[0]) == "." {
				if !regexpMatch(regexp, input[1:]) {
					if len(regexp[2:]) == len(input[1:]) {
						return regexpMatch(regexp[2:], input[1:])
					}
					return regexpMatch(regexp[2:], input[1:])
				}
				return regexpMatch(regexp, input[1:])
			} else {
				return regexpMatch(regexp[2:], input)
			}
		}
		if string(regexp[1]) == "+" {
			if regexp[0] == input[0] || string(regexp[0]) == "." {
				if !regexpMatch(regexp, input[1:]) {
					if len(regexp[2:]) == len(input[1:]) {
						return regexpMatch(regexp[2:], input[1:])
					}
					return regexpMatch(regexp[2:], input[1:])
				}
				return regexpMatch(regexp, input[1:])
			} else {
				return false
			}
		}
	}
	if regexp[0] == input[0] || string(regexp[0]) == "." {
		return regexpMatch(regexp[1:], input[1:])
	}
	return false
}

func entriPoint(regexp, input string) bool {
	if regexpMatch(regexp, input) {
		return true
	}
	if input != "" && string(regexp[0]) != "^" {
		return entriPoint(regexp, input[1:])
	}
	return false
}

func main() {
	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()

	splitedInput := strings.Split(userInput.Text(), "|")
	regexpLine := splitedInput[0]
	inputLine := splitedInput[1]

	fmt.Println(entriPoint(regexpLine, inputLine))

}
