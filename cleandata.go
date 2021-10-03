package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/gocarina/gocsv"
	//https://github.com/gocarina/gocsv
)

//Global variables
var fileNames = []string{"crew.csv", "customer_ratings.csv", "names.csv", "principals.csv", "titles.csv"}
var rawExten = "rawData\\"
var cleanExten = "cleanedData\\"

//Structs used for unmarshalling lines into easy format to clean and write
type Crew struct {
	titleID   string `csv:"titleID"`
	directors string `csv:"directors"`
	writers   string `csv:"writers"`
}

func cleanCrew() {
	//Used for debugging
	//_ = fileName
	fmt.Println("Cleaning " + rawExten + "crew.csv" + " now...")

	//Establish link to output file
	readFile, err := os.Open(rawExten + "crew.csv")

	if err != nil {
		fmt.Printf("failed reading data from file: %s", err)
	}

	defer readFile.Close()

	//Create / Establish link to output file
	outputFile, err := os.Create(cleanExten + "crew.csv")

	if err != nil {
		fmt.Printf("failed writing to file: %s", err)
	}

	defer outputFile.Close()

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(readFile)
		r.Comma = '\t'
		return r
	})

	crew := []*Crew{}
	gocsv.UnmarshalFile(readFile, &crew)

	fmt.Println(crew[0])

	//for _, crew := range crew {
	//	fmt.Println(crew)

	//_, err = io.WriteString(outputFile, crew.titleID+"\t"+crew.directors+"\t"+crew.writers+"\n")

	// if err != nil {
	// 	fmt.Print(err)
	// }
	//}

	/*

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			line := fileScanner.Text()

			properties := strings.Split(line, "\t")

			_, err = io.WriteString(outputFile, properties)

			if err != nil {
				fmt.Print(err)
			}
		}
	*/
}

func cleanData() {
	cleanCrew()
}
