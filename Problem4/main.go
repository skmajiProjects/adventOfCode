package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal()
	}
	defer f.Close()

	recordList, err := GetRecordList(f, "\r\n\r\n", "\r\n", " ")
	if err != nil {
		log.Fatal(err)
	}
	var validPassportCountWithMethod1 int = 0
	var validPassportCountWithMethod2 int = 0
	for _, record := range recordList {
		p, err := getPassport(record)
		if err != nil {
			log.Fatal(err)
		}

		if isPassportValidMethod1(p) {
			validPassportCountWithMethod1++
		}

		if isPassportValidMethod2(p) {
			validPassportCountWithMethod2++
		}

	}
	fmt.Println("count of valid Passports with method 1=", validPassportCountWithMethod1)
	fmt.Println("count of valid Passports with method 2=", validPassportCountWithMethod2)

}
