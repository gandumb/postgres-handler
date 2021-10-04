package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
	//https://github.com/gocarina/gocsv
)

//Global variables
var fileNames = []string{"crew.csv", "customer_ratings.csv", "names.csv", "principals.csv", "titles.csv"}
var rawExten = "rawData\\"
var cleanExten = "cleanedData\\"

//Helper function for cleanData() string building
func buildString(properties []string) string {
	var output string

	for i, data := range properties {
		data := strings.Map(func(r rune) rune {
			if r > unicode.MaxASCII {
				return -1
			}
			return r
		}, data)

		data = strings.TrimSpace(data)

		if i < (len(properties) - 1) {
			output += data + "\t"
		} else {
			output += data
		}
	}

	output += "\n"

	return output
}

func cleanData() {
	for _, fileName := range fileNames {

		//_ = fileName // Used for debugging

		fmt.Println("Cleaning " + rawExten + fileName + " now...")

		//Establish link to output file
		readFile, err := os.Open(rawExten + fileName)

		if err != nil {
			fmt.Printf("failed reading data from file: %s", err)
		}

		defer readFile.Close()

		//Create / Establish link to output file
		outputFile, err := os.Create(cleanExten + fileName)

		if err != nil {
			fmt.Printf("failed writing to file: %s", err)
		}

		defer outputFile.Close()

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		hashMap := make(map[string]bool)

		for fileScanner.Scan() {
			line := fileScanner.Text()

			properties := strings.Split(line, "\t")

			//Clean Data Here

			properties = properties[1:]

			//Check for duplicates
			if _, ok := hashMap[properties[0]]; ok {
				continue
			} else {
				hashMap[properties[0]] = true
			}

			//Convert WIN1252 characters to UTF8

			// //Check for nulls
			// for i, data := range properties {
			// 	if data == "" {
			// 		properties[i] = "NULL"
			// 	}
			// }

			//fmt.Println(buildString(properties)) //Used for debugging

			_, err = io.WriteString(outputFile, buildString(properties))

			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
