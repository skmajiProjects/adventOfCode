package main

import (
	"io"
	"io/ioutil"
	"strings"
)

func GetRecordList(r io.Reader, recDelim string, newlineChar string, fieldDelim string) ([][]string, error) {
	//Convert reader to string
	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		return [][]string{{}, {}}, err
	}

	inputStr := string(bytes)

	var recordList [][]string
	records := strings.Split(inputStr, recDelim)
	for _, rec := range records {
		cleanRec := strings.Replace(rec, newlineChar, fieldDelim, -1)
		//fmt.Println("Field Value=", cleanRec)
		recordList = append(recordList, strings.Split(cleanRec, fieldDelim))
	}
	return recordList, nil
}
