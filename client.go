package main

// "encoding/json"
// "log"
// "net/http"
// "context"
// "fmt"
// "os"

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	var response string

	fmt.Print("Do you want to clean data yes (y) no (n) >>> ")
	fmt.Scanln(&response)

	for ; response != "y" && response != "n"; fmt.Scanln(&response) {
		fmt.Print("Do you want to clean data yes (y) no (n) >>> ")
	}

	if response == "y" {
		cleanData()
	}

	conn, err := pgxpool.Connect(context.Background(), "postgres://csce315904_4user:Helpme012@csce-315-db.engr.tamu.edu/csce315904_4db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//Close the connection
	defer conn.Close()

	/*
		//Hello World
		var greeting string
		err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(greeting)
	*/

	//Verify that tables exists

	// Send the query to the server.

	/*
		//Insert something into the table
		_, err = conn.Exec(context.Background(), "INSERT INTO TeamMembers VALUES ('Testy Testingson', 000, 'Test Test', '10/02/21')")
		if err != nil {
			fmt.Print("Fail here")
			log.Fatal(err)
		}
	*/

	//fmt.Print(response)

	//Create tables if they don't exist

	//Upload data to the tables
}
