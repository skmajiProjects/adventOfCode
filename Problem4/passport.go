package main

import (
	"regexp"
	"strconv"
	"strings"
)

type year struct {
	length  int
	minSize int
	maxSize int
}

type strPattern string

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

var byrYear = year{4, 1920, 2002}
var iyrYear = year{4, 2010, 2020}
var eyrYear = year{4, 2020, 2030}

var hclStrPattern = strPattern("#[0-9a-z]+")
var pidStrPattern = strPattern("^\\d{9}$")

var eyeColorList = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

var requiredField = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
	"cid": false,
}

//returns a struc of type Passport for every record
func getPassport(record []string) (Passport, error) {
	var p Passport
	for _, r := range record {
		fieldName := strings.Split(r, ":")[0]
		if fieldName == "byr" {
			p.byr = strings.Split(r, ":")[1]
		}
		if fieldName == "iyr" {
			p.iyr = strings.Split(r, ":")[1]
		}
		if fieldName == "eyr" {
			p.eyr = strings.Split(r, ":")[1]
		}
		if fieldName == "hgt" {
			p.hgt = strings.Split(r, ":")[1]
		}
		if fieldName == "hcl" {
			p.hcl = strings.Split(r, ":")[1]
		}
		if fieldName == "ecl" {
			p.ecl = strings.Split(r, ":")[1]
		}
		if fieldName == "pid" {
			p.pid = strings.Split(r, ":")[1]
		}
		if fieldName == "cid" {
			p.cid = strings.Split(r, ":")[1]
		}

	}
	return p, nil
}

//validate by method1 (manditory)
func isPassportValidMethod1(p Passport) bool {
	return areRequiredFieldsPresent(p)
}

//validate by method1(manditory and more)
func isPassportValidMethod2(p Passport) bool {
	var isValid bool
	if !(areRequiredFieldsPresent(p)) {
		return false
	}

	isValid = isYearValid(p.byr, byrYear)
	if !isValid {
		return false
	}
	isValid = isYearValid(p.iyr, iyrYear)
	if !isValid {
		return false
	}
	isValid = isYearValid(p.eyr, eyrYear)
	if !isValid {
		return false
	}
	isValid = isHgtValid(p.hgt)
	if !isValid {
		return false
	}
	isValid = isPatternValid(p.hcl, hclStrPattern)
	if !isValid {
		return false
	}
	isValid = isEclValid(p.ecl, eyeColorList)
	if !isValid {
		return false
	}
	isValid = isPatternValid(p.pid, pidStrPattern)
	if !isValid {
		return false
	}
	isValid = isCidValid(p.cid)
	if !isValid {
		return false
	}
	return true
}

//checks if Passport fields of type year are valid
func isYearValid(fieldVal string, by year) bool {
	fieldValNum, err := strconv.Atoi(fieldVal)
	if err != nil {
		return false
	}
	if len(fieldVal) < by.length {
		return false
	}
	if !((by.minSize <= fieldValNum) && (fieldValNum <= by.maxSize)) {
		return false
	}
	return true
}

//checks if Passport field Hgt is valid
func isHgtValid(fieldVal string) bool {
	var minValCm int = 150
	var maxValCm int = 193
	var minValIn int = 59
	var maxValIn int = 76

	if len(fieldVal) < 3 {
		return false
	}

	var dim = fieldVal[len(fieldVal)-2:]
	var fieldValNum, err = strconv.Atoi(fieldVal[:len(fieldVal)-2])
	if err != nil {
		return false
	}
	if dim == "cm" {
		if !((minValCm <= fieldValNum) && (fieldValNum <= maxValCm)) {
			return false
		}
	}
	if dim == "in" {
		if !((minValIn <= fieldValNum) && (fieldValNum <= maxValIn)) {
			return false
		}
	}
	return true
}

//checks if Passport field Hcl, Pid is valid (pattern matching)
func isPatternValid(fieldVal string, sp strPattern) bool {
	match, _ := regexp.MatchString(string(sp), fieldVal)
	return match
}

//checks if Passport field Ecl is one of the colors in the Color list
func isEclValid(fieldVal string, eyeColorList []string) bool {
	for _, val := range eyeColorList {
		if fieldVal == val {
			return true
		}
	}
	return false
}

//checks if Passport field Cid is valid (no validations as of now)
func isCidValid(fieldVal string) bool {
	return true
}

// Checks if required fields have values
func areRequiredFieldsPresent(p Passport) bool {
	switch {
	case requiredField["byr"] && p.byr == "":
		return false
	case requiredField["iyr"] && p.iyr == "":
		return false
	case requiredField["eyr"] && p.eyr == "":
		return false
	case requiredField["hgt"] && p.hgt == "":
		return false
	case requiredField["hcl"] && p.hcl == "":
		return false
	case requiredField["ecl"] && p.ecl == "":
		return false
	case requiredField["pid"] && p.pid == "":
		return false
	case requiredField["cid"] && p.cid == "":
		return false
	}
	return true
}
