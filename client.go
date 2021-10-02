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
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conn, err := pgxpool.Connect(context.Background(), "postgres://csce315904_4user:Helpme012@csce-315-db.engr.tamu.edu/csce315904_4db")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()

	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	var sum int32

	// Send the query to the server. The returned rows MUST be closed
	// before conn can be used again.
	rows, err := conn.Query(context.Background(), "select generate_series(1,$1)", 10)
	if err != nil {
		log.Fatal("Error")
	}

	// rows.Close is called by rows.Next when all rows are read
	// or an error occurs in Next or Scan. So it may optionally be
	// omitted if nothing in the rows.Next loop can panic. It is
	// safe to close rows multiple times.
	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var n int32
		err = rows.Scan(&n)
		if err != nil {
			log.Fatal("Error")
		}
		sum += n
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		log.Fatal("Error")
	}

	// No errors found - do something with sum
}
