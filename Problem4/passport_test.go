package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_passportValidate(t *testing.T) {

	//Given
	passportList := `ecl:gry pid:475471726 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	//When
	t.Run("Given the Passport List with validation list 1", func(t *testing.T) {

		recordList, err := GetRecordList(strings.NewReader(passportList), "\n\n", "\n", " ")
		if err != nil {
			t.Fatal(err)
		}

		var validPassportCount int = 0
		for _, record := range recordList {
			p, err := getPassport(record)
			if err != nil {
				t.Fatal(err)
			}

			if isPassportValidMethod1(p) {
				validPassportCount++
			}

		}

		if validPassportCount != 2 {
			fmt.Println("Calculated Count=", validPassportCount)
			fmt.Println("Correct Count=", 2)
			t.Fatal("test Failed! the count of invalid passports is incorrect.")
		}
	})

	//When
	t.Run("Given the Passport List with validation list 2", func(t *testing.T) {

		recordList, err := GetRecordList(strings.NewReader(passportList), "\n\n", "\n", " ")
		if err != nil {
			t.Fatal(err)
		}

		var validPassportCount int = 0
		for _, record := range recordList {
			p, err := getPassport(record)
			if err != nil {
				t.Fatal(err)
			}
			//fmt.Println("Passport Birth Year=", p.byr)
			if isPassportValidMethod2(p) {
				validPassportCount++
			}

		}

		if validPassportCount != 2 {
			fmt.Println("Calculated Count=", validPassportCount)
			fmt.Println("Correct Count=", 2)
			t.Fatal("test Failed! the count of invalid passports is incorrect.")
		}
	})

}
