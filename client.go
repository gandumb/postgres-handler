package main

// "encoding/json"
// "log"
// "net/http"
// "context"
// "fmt"
// "os"

import (
	"fmt"
)

// func createTables(conn *pgx.Conn) {
// 	_, err := conn.Query(context.Background(), "CREATE TABLE COMPANY (ID INT PRIMARY KEY NOT NULL, NAME TEXT NOT NULL, AGE INT NOT NULL, ADDRESS CHAR(50), SALARY REAL)")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// }

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

	//postgres://csce315904_4user:Helpme012@csce-315-db.engr.tamu.edu/csce315904_4db
}
