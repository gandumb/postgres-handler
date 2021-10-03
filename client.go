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
	"os/exec"

	"github.com/jackc/pgx/v4"
)

func createTables(conn *pgx.Conn) {
	_, err := conn.Query(context.Background(), "CREATE TABLE COMPANY (ID INT PRIMARY KEY NOT NULL, NAME TEXT NOT NULL, AGE INT NOT NULL, ADDRESS CHAR(50), SALARY REAL)")
	if err != nil {
		fmt.Println(err)
	}

}

func commandLine() {
	cmd := "psql"
	args := []string{"-U", "postgres", "-d", "test", "-c", fmt.Sprintf(`\copy doe from '%s' delimiter '\t' csv header;`, os.Args[1])}
	v, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		panic(string(v))
	}
	println("Ok")
}

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

	commandLine()

	conn, err := pgx.Connect(context.Background(), "postgres://csce315904_4user:Helpme012@csce-315-db.engr.tamu.edu/csce315904_4db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	//Close the connection
	defer conn.Close(context.Background())

	var sum int32

	// Send the query to the server. The returned rows MUST be closed
	// before conn can be used again.
	rows, err := conn.Query(context.Background(), "select generate_series(1,$1)", 100)
	if err != nil {
		fmt.Println(err)
	}
	// Iterate through the result set
	for rows.Next() {
		var n int32
		err = rows.Scan(&n)
		if err != nil {
			fmt.Println(err)
		}
		sum += n
	}

	fmt.Println(sum)

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		fmt.Println(rows.Err())
	}

	rows.Close()

	rows, err = conn.Query(context.Background(), "SELECT EXISTS( SELECT * FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'TeamMembers')")
	if err != nil {
		fmt.Println(err)
	}
	// Iterate through the result set
	for rows.Next() {
		if err != nil {
			fmt.Println(err)
		}
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		fmt.Println(rows.Err())
	}

	fmt.Println(rows)

	rows.Close()

	// No errors found - do something with sum

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
}
