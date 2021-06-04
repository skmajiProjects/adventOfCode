package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sumTotal int = 2020

func main() {
	var loopCount int // Count of numbers adding up to 2020
	sum := 0
	product := 1
	loopInit := 0

	// Read the list from a file //
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// store the
	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please enter the count of numbers :")
	fmt.Scanf("%d", &loopCount)

	//Function to recursively find the entries which sum up to 2020
	checkCombinationRecursive(lines, loopCount, loopInit, sum, product)

}

func readLines(file *os.File) ([]string, error) {
	//reads the lines from the file and returns a slice
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func checkCombinationRecursive(lines []string, loopCount int, loopInit int, sum int, product int) {

	if loopCount > 0 {
		for i := loopInit; i < len(lines); i++ {
			number1, err := strconv.Atoi(lines[i])
			if err != nil {
				log.Fatal(err)
			}

			sumCurr := sum + number1
			productCurr := product * number1
			checkCombinationRecursive(lines, loopCount-1, i+1, sumCurr, productCurr)

			if sumCurr == sumTotal && loopCount == 1 {
				fmt.Println("Product of Numbers =", productCurr)
			}
		}
	}
}
