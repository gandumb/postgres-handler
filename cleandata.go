package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	//https://github.com/gocarina/gocsv
)

//Global variables
var fileNames = []string{"crew.csv", "customer_ratings.csv", "names.csv", "principals.csv", "titles.csv"}
var rawExten = "rawData\\"
var cleanExten = "cleanedData\\"

//Helper function for cleanData() string building
func buildString(properties []string) string {
	var output string

	for _, data := range properties {
		output += data + "\t"
	}

	output += "\n"

	return output
}

func cleanLine(line []string) {
	for _, section := range line {
		_ = section
	}
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

		//minLength := 0

		for fileScanner.Scan() {
			line := fileScanner.Text()

			properties := strings.Split(line, "\t")

			//Clean Data Here

			properties = properties[1:]

			//fmt.Println(buildString(properties)) //Used for debugging

			_, err = io.WriteString(outputFile, buildString(properties))

			if err != nil {
				fmt.Print(err)
			}
		}
	}
}
