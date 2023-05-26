package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

const alfabetString = "abcdefghijklmnopqrstuvwxyz"
const willYouMerry = "Will you marry me?"

func encodeMessage(offset int, inputString string) string {
	var encodedString string
	var position int
	for _, element := range inputString {
		position = strings.Index(alfabetString, strings.ToLower(string(element)))
		if position == -1 {
			encodedString += string(element)
		} else {
			if unicode.IsUpper(element) {
				newPosition := (position + offset) % len(alfabetString)
				encodedString += strings.ToUpper(string(alfabetString[newPosition]))
			} else {
				newPosition := (position + offset) % len(alfabetString)
				encodedString += string(alfabetString[newPosition])
			}
		}
	}
	return encodedString
}

func decodeMessage(offset int, inputString string) string {
	var decodedString string
	var position int
	for _, element := range inputString {
		position = strings.Index(alfabetString, strings.ToLower(string(element)))
		if position == -1 {
			decodedString += string(element)
		} else {
			if unicode.IsUpper(element) {
				newPosition := int(math.Abs(float64((position - offset) % len(alfabetString))))
				decodedString += strings.ToUpper(string(alfabetString[newPosition]))
			} else {
				newPosition := int(math.Abs(float64((position - offset) % len(alfabetString))))
				decodedString += string(alfabetString[newPosition])
			}
		}
	}
	return decodedString
}

// replace all calculation by modulus

func main() {
	// write your code here
	var g, p int
	_, err := fmt.Scanf("g is %d and p is %d", &g, &p)
	if err != nil {
		fmt.Print("Cannot parse input")
	}
	// fmt.Printf("g=%d and p=%d\n", g, p)
	fmt.Print("OK\n")

	b := 7
	var A, B, S int
	A, B, S = 1, 1, 1

	for i := 0; i < b; i++ {
		B = B * g % p
	}

	fmt.Scanf("A is %d", &A)

	for i := 0; i < b; i++ {
		S = S * A % p
	}

	fmt.Printf("B is %d\n", B)
	// fmt.Printf("A is %d\n", A)
	// fmt.Printf("S is %d\n", S)

	// index := strings.Index(alfabetString, "e")

	/*var inputForEncode string
	fmt.Print("Please enter string for encode:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputForEncode = scanner.Text()*/

	offset := S % len(alfabetString)

	fmt.Println(encodeMessage(offset, willYouMerry))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputForDecode := scanner.Text()

	fmt.Println(decodeMessage(offset, inputForDecode))

	/*for _, element := range strings.ToLower(inputForEncode) {
		position := strings.Index(alfabetString, string(element)) // if not exist we receive -1 paste as is
		fmt.Printf("Original position: %d\n", position)
	}
	*/
	/* if (position + offset) > len(alfabetString) {

	}
	*/
	// apply s to modulus 26
	// for encode - take latter index apply offset
	// for decode - take latter index apply offset

	// position + offset > 25 / overload {absolute function(25 - (offset + position))}
	// position + offset <= 25 / normal {position + offset}

}
