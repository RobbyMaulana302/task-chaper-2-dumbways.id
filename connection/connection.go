package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

// fungsi konek ke database
func DatabaseConnect()  {
	
	// postgres://user:password@host:port/dbname
	databaseURL := "postgres://postgres:root@localhost:5432/db_project"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect Database %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Succesfully connected to database!")

}