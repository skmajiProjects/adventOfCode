package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read Filename as argument
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Store all the lines in the file to String array
	lines, err := readLines(file)
	if err != nil {
		log.Fatal(err)
	}

	//get count for password policy 1
	validPasswordPolicyCount1 := getValidPasswordPolicyCount1(lines)
	fmt.Println("Valid Count with password policy 1:", validPasswordPolicyCount1)

	//get count for password policy 1
	validPasswordPolicyCount2 := getValidPasswordPolicyCount2(lines)
	fmt.Println("Valid Count with password policy 2:", validPasswordPolicyCount2)
}

func readLines(file *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func getValidPasswordPolicyCount1(lines []string) int {
	var validPdCount int
	var character, passwd string

	for _, val := range lines {
		passwd = strings.Split(val, ": ")[1]
		character = strings.Split(strings.Split(val, ":")[0], " ")[1]
		min, err := strconv.Atoi(strings.Split(strings.Split(strings.Split(val, ":")[0], " ")[0], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(strings.Split(strings.Split(strings.Split(val, ":")[0], " ")[0], "-")[1])
		if err != nil {
			log.Fatal(err)
		}

		countOfCharinString := strings.Count(passwd, character)
		if countOfCharinString >= min && countOfCharinString <= max {
			validPdCount++
		}
		//fmt.Println("min:", min, ", max:", max, ", Char:", character, ", password:", passwd)
	}
	return validPdCount
}

func getValidPasswordPolicyCount2(lines []string) int {
	var validPdCount int
	var character, passwd string
	for _, val := range lines {
		passwd = strings.Split(val, ": ")[1]
		character = strings.Split(strings.Split(val, ":")[0], " ")[1]
		min, err := strconv.Atoi(strings.Split(strings.Split(strings.Split(val, ":")[0], " ")[0], "-")[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(strings.Split(strings.Split(strings.Split(val, ":")[0], " ")[0], "-")[1])
		if err != nil {
			log.Fatal(err)
		}

		//countOfCharinString := strings.Count(passwd, character)
		//fmt.Println("min:", min, "max:", max)
		//fmt.Println("check slice on Shring:", passwd[min-1:min], passwd[max-1:max], "Char:", character, "password:", passwd)

		if passwd[min-1:min] == character && passwd[max-1:max] != character {
			validPdCount++
		} else if passwd[min-1:min] != character && passwd[max-1:max] == character {
			validPdCount++
		}
	}
	return validPdCount
}
