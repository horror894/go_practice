package main

import (
	"fmt"
	"math"
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
			if (offset + position) >= len(alfabetString) {
				fmt.Println(len(alfabetString))
				fmt.Println(offset + position)
				newPosition := math.Abs(float64(len(alfabetString) - (offset + position)))
				fmt.Println(newPosition)
				if unicode.IsUpper(element) {
					encodedString += strings.ToUpper(string(alfabetString[int(newPosition)]))
				} else {
					encodedString += string(alfabetString[int(newPosition)])
				}
			} else if (offset + position) < len(alfabetString) {
				newPosition := offset + position
				if unicode.IsUpper(element) {
					encodedString += strings.ToUpper(string(alfabetString[newPosition]))
				}
				encodedString += string(alfabetString[newPosition])
			}
		}
	}
	return encodedString
}

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

	for i := 1; i <= b; i++ {
		B = B * g % p
	}

	fmt.Scanf("A is %d", &A)

	for i := 1; i <= A; i++ {
		S = S * B % p
	}

	fmt.Printf("B is %d\n", B)
	// fmt.Printf("A is %d\n", A)
	// fmt.Printf("S is %d\n", S)

	// index := strings.Index(alfabetString, "e")

	offset := S % len(alfabetString) // create real offset

	/*var inputForEncode string
	fmt.Print("Please enter string for encode:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputForEncode = scanner.Text()*/

	fmt.Println(S)
	fmt.Println(offset)
	fmt.Print(encodeMessage(offset, willYouMerry))

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
