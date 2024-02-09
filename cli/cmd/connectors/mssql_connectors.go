package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/microsoft/go-mssqldb"
)

func mssqlCredsConnection(addr string, unam string, pwrd string, db string, port string) (*sql.DB, error) {
	var connString string = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", addr, unam, pwrd, port, db)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Failed to open connection: ", err.Error())
	}
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn, err
}

// Can be used with Entra or U&P... still yet to be tested
func mssqlStringConnection(connString string) (*sql.DB, error) {
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Failed to open connection: ", err.Error())
	}
	ctx := context.Background()
	err = conn.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	return conn, err
}

func closeDBconnection(conn *sql.DB) {
	conn.Close()
}
