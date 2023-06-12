package main

import (
	"fmt"
	"math/rand"
	"os"
)

/* finaly I want to chage logic of work
it's the same as I saw in  bufio.NewScanner().
I want to do the same, we initialize object giving it some params, after that call method and recive return
in my case with matrix.
Because we create struct not for external use, for use it with specific methods.


Rewrite part with new generation
To many "if", reduce count of if statement

Add real time show in console

*/

type univirceObject struct {
	alive  int
	size   int
	seed   int
	matrix [][]string
}

func (s *univirceObject) newMatrix(size int, seed int) {
	s.matrix = make([][]string, size)
	for i := range s.matrix {
		s.matrix[i] = make([]string, size)
	}

	s.size = size
	s.seed = seed

	rand.Seed(int64(seed))

	for i := range s.matrix {
		for a := range s.matrix[i] {
			randomNumber := rand.Intn(2)
			if randomNumber == 0 {
				s.matrix[i][a] = " "
			} else {
				s.matrix[i][a] = "O"
			}

		}
	}

}

func (original *univirceObject) newGeneration() {

	tempSlice := make([][]string, original.size)
	for i := range tempSlice {
		tempSlice[i] = make([]string, original.size)
	}

	for i := range original.matrix {
		for a := range original.matrix[i] {
			// find naibors
			// fist step set coordinats for all neighbors
			// i - SN, a - EW;
			var iN, iS, aE, aW int // coordinats of neighbors
			// find N
			if i-1 < 0 {
				iN = (i-1)%original.size + original.size
			} else {
				iN = i - 1
			}
			// find S
			if i+1 >= original.size {
				iS = (i + 1) % original.size
			} else {
				iS = i + 1
			} // find E
			if a-1 < 0 {
				aE = (a-1)%original.size + original.size
			} else {
				aE = a - 1
			}
			// find W
			if a+1 >= original.size {
				aW = (a + 1) % original.size
			} else {
				aW = a + 1
			}
			// check neigbours
			countOfNeighbors := 0
			if original.matrix[iN][a] == "O" {
				countOfNeighbors++
			}
			if original.matrix[iS][a] == "O" {
				countOfNeighbors++
			}
			if original.matrix[i][aE] == "O" {
				countOfNeighbors++
			}
			if original.matrix[i][aW] == "O" {
				countOfNeighbors++
			}
			if original.matrix[iN][aE] == "O" {
				countOfNeighbors++
			}
			if original.matrix[iN][aW] == "O" {
				countOfNeighbors++
			}
			if original.matrix[iS][aE] == "O" {
				countOfNeighbors++
			}
			if original.matrix[iS][aW] == "O" {
				countOfNeighbors++
			}

			// fmt.Printf("for element %d and %d, Neibors: coordinats:%d %d %d %d\n", i, a, iN, iS, aE, aW)

			// make disigion
			if original.matrix[i][a] == "O" {
				if countOfNeighbors < 2 || countOfNeighbors > 3 {
					tempSlice[i][a] = " "
				} else {
					tempSlice[i][a] = "O"
				}

			} else {
				if countOfNeighbors == 3 {
					tempSlice[i][a] = "O"
				} else {
					tempSlice[i][a] = " "
				}
			}

			// above line

		}
	}

	original.alive = 0
	for i := range original.matrix {
		for a := range original.matrix[i] {
			if tempSlice[i][a] == "O" {
				original.matrix[i][a] = "O"
				original.alive++
			} else {
				original.matrix[i][a] = " "
			}
		}
	}
}

func main() {

	testMap := map[string][][]int {
		"N":
	}



	var size, seed, generationCount int
	generationCount = 1
	for tryCount := 10; tryCount > 0; tryCount-- {
		if n, err := fmt.Scanf("%d", &size); n < 1 {
			fmt.Printf("Wrong input %s\n", err)
			if tryCount > 1 {
				continue
			} else {
				os.Exit(-1)
			}
		}
		break
	}

	myUniverce := univirceObject{}
	myUniverce.newMatrix(size, seed)

	for generationCount <= 10 {
		myUniverce.newGeneration()
		fmt.Printf("Generation #%d\n", generationCount)
		fmt.Printf("Alive: %d\n", myUniverce.alive)
		for i := range myUniverce.matrix {
			for _, element := range myUniverce.matrix[i] {
				fmt.Print(element)
			}
			fmt.Print("\n")
		}
		generationCount++
	}

}

/* I like this example !!!
// DO NOT delete or modify the contents of the main() function!
func main() {
	date1, date2, date3 := readDate(), readDate(), readDate()

	// The first travel — checks conditions for the latest date:
	checkFirstTravel(date1, date2, date3)

	// The second travel — checks conditions for the earliest date:
	checkSecondTravel(date1, date2, date3)

	// The third travel — checks conditions for the middle date:
	checkThirdTravel(date1, date2, date3)
}

// DO NOT modify the readDate function!
func readDate() time.Time {
	var year, month, day int
	fmt.Scan(&year, &month, &day)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
*/
