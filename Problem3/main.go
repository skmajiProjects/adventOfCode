package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//get the filename from cmd aggr
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//read the file and store the content in a string slice
	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	//path to be taken
	gridMatrix := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	var product int = 1
	for _, val := range gridMatrix {
		countTrees := traverseGrid(lines, val[0], val[1])
		product = product * countTrees

		fmt.Println("Count Of Trees with path (", val[0], ",", val[1], ") is = ", countTrees)
	}
	fmt.Println("Product of Count Of Trees:", product)
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func traverseGrid(lines []string, col int, row int) int {
	var x, treeCount int
	var treeStr string = "#"

	var newLine string

	//iterate through each grid line
	for i := 0; i < len(lines); i = i + row {
		currLine := lines[i]
		newLine = currLine
		if i > 0 {
			if len(currLine) <= x {
				multiple := x / len(currLine)
				for count := 0; count <= multiple; count++ {
					newLine = newLine + currLine
				}
			}
			charAtPos := newLine[x : x+1]
			if charAtPos == treeStr {
				treeCount++
			}
		}
		x = x + col
	}

	return treeCount
}
