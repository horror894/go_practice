package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

// create const for promt
// make main more clear
// create one function for encript/decript
// create struct for secretes

const alfabetString = "abcdefghijklmnopqrstuvwxyz"
const willYouMerry = "Will you marry me?"

func encodeMessage(offset int, inputString string, alphabet string) string {
	var encodedString string
	var position int
	size := len(alphabet)

	for _, element := range inputString {
		position = strings.Index(alphabet, strings.ToLower(string(element)))
		if position == -1 { // if not exist in alphabet paste as is
			encodedString += string(element)
		} else {
			newPosition := (position + offset) % size // calculate position
			if unicode.IsUpper(element) {
				encodedString += strings.ToUpper(string(alphabet[newPosition]))
			} else {
				encodedString += string(alphabet[newPosition])
			}
		}
	}
	return encodedString
}

func decodeMessage(offset int, inputString string, alphabet string) string {
	var decodedString string
	var position int
	size := len(alphabet)

	for _, element := range inputString {
		position = strings.Index(alfabetString, strings.ToLower(string(element)))
		if position == -1 {
			decodedString += string(element)
		} else {
			newPosition := (position - offset) % size
			if newPosition < 0 {
				newPosition += len(alfabetString)
			}
			if unicode.IsUpper(element) {
				decodedString += strings.ToUpper(string(alfabetString[newPosition]))
			} else {
				decodedString += string(alfabetString[newPosition])
			}
		}
	}
	return decodedString
}

// replace all calculation by modulus

func receiveSharedNum() (g, p int) {
	_, err := fmt.Scanf("g is %d and p is %d", &g, &p)
	if err != nil {
		fmt.Print("Cannot parse input")
	}
	fmt.Print("OK\n")
	return g, p
}

func generatePublicKey(g, p, b int) int {
	key := 1
	for i := 0; i < b; i++ {
		key = key * g % p
	}
	return key
}

func main() {
	var g, p int       // large prime number
	var b int          // my secret
	A, B, S := 1, 1, 1 // Alice public key, Bob public key, Secret for encript

	g, p = receiveSharedNum() // receive variable from stdin
	b = rand.Intn(300)        // generate my secret

	B = generatePublicKey(g, p, b) // generate my public key

	fmt.Scanf("A is %d", &A)

	S = generatePublicKey(A, p, b)

	fmt.Printf("B is %d\n", B)

	fmt.Println(encodeMessage(S, willYouMerry, alfabetString))

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputForDecode := scanner.Text()

	answer := decodeMessage(S, inputForDecode, alfabetString)

	switch answer {
	case "Yeah, okay!":
		fmt.Println(encodeMessage(S, "Great!", alfabetString))
	case "Let's be friends.":
		fmt.Println(encodeMessage(S, "What a pity!", alfabetString))

	}

}
