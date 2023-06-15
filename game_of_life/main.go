package main

import (
	"fmt"
	"math/rand"
	"os"
)

/*
finaly I want to chage logic of work
it's the same as I saw in  bufio.NewScanner().
I want to do the same, we initialize object giving it some params, after that call method and receive return
in my case with matrix.
Because we create struct not for external use, for use it with specific methods.

Add real time show in console


*/

type universeObject struct {
	alive  int
	size   int
	seed   int
	status bool
	matrix [][]string
}

func newUniverse(size int) *universeObject {
	return &universeObject{
		size: size,
	}
}

func (s *universeObject) newMatrix() {

	s.status = true
	// matrix initialization
	s.matrix = make([][]string, s.size)
	for i := range s.matrix {
		s.matrix[i] = make([]string, s.size)
	}

	rand.Seed(14)
	// fill matrix randomly
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

func (original *universeObject) newGeneration() {

	if original.status == false {
		panic("Generations call for empty universe")
	}

	aliveCount := 0

	// create temporary slice for intermediate calculations
	tempSlice := make([][]string, original.size)
	for i := range tempSlice {
		tempSlice[i] = make([]string, original.size)
	}

	for i := range original.matrix {
		for a := range original.matrix[i] {
			countOfNeigbours := original.countNeigboursForPossition(a, i)
			if shouldLive(original.matrix[i][a], countOfNeigbours) {
				tempSlice[i][a] = "O"
				aliveCount++
			} else {
				tempSlice[i][a] = " "
			}

		}
	}
	original.alive = aliveCount
	original.matrix = tempSlice
}

func (original *universeObject) countNeigboursForPossition(a, i int) int {
	neighboursPaterns := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	neighCount := 0

	// check neighbours
	for _, dir := range neighboursPaterns {
		y := (dir[0] + i + original.size) % original.size
		x := (dir[1] + a + original.size) % original.size
		if original.matrix[y][x] == "O" {
			neighCount++
		}
	}
	return neighCount
}

func shouldLive(currentState string, neighbours int) bool {
	if currentState == "O" {
		if neighbours < 2 || neighbours > 3 {
			return false
		}
		return true
	} else if currentState == " " {
		if neighbours == 3 {
			return true
		}
		return false
	}
	return false
}

func scanUserInput(input *int) *int {
	for tryCount := 10; tryCount > 0; tryCount-- {
		if n, err := fmt.Scanf("%d", input); n < 1 {
			fmt.Printf("Wrong input %s\n", err)
			if tryCount > 1 {
				continue
			} else {
				os.Exit(-1)
			}
		}
		break
	}
	return input
}

func main() {
	var size int
	generationCount := 0
	// input read
	scanUserInput(&size)

	myUniverce := newUniverse(size)
	myUniverce.newMatrix()

	for generationCount <= 10 {
		fmt.Printf("Generation #%d\n", generationCount)
		fmt.Printf("Alive: %d\n", myUniverce.alive)
		for i := range myUniverce.matrix {
			for _, element := range myUniverce.matrix[i] {
				fmt.Print(element)
			}
			fmt.Print("\n")
		}
		myUniverce.newGeneration()
		generationCount++
	}

}
