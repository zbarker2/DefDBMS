package cmd

import (
	"testing"
)

// All test functions are set up to use the databases in the docker compose.
func testMssqlCredsDBConnection(t *testing.T) {
	var addr string = "localhost"
	var unam string = "sa"
	var pwrd string = "ThisIsASecur3Pass!"
	var db string = "students"
	var port string = "1433"
	conn, err := mssqlCredsConnection(addr, unam, pwrd, db, port)

	if conn.Stats().OpenConnections < 1 || err != nil {
		t.Fatalf(`mssqlCredsDBConnection(addr,unam,pwrd,db,port)= %s`, "Not Connected")
	}

}
func testBadMssqlStringConnection(t *testing.T) {
	var connString = "yeet"
	conn, err := mssqlStringConnection(connString)

	if conn.Stats().OpenConnections > 0 || err == nil {
		t.Fatalf(`mssqlCredsDBConnection(connString)= %s`, " Should not be Connected")
	}
}
func testGoodStringConnection(t *testing.T) {
	var connString = "server=localhost;user id=sa;password=ThisIsASecur3Pass!;port=1433;database=students"
	conn, err := mssqlStringConnection(connString)

	if conn.Stats().OpenConnections < 1 || err != nil {
		t.Fatalf(`mssqlCredsDBConnection(addr,unam,pwrd,db,port)= %s`, "Not Connected")
	}
}
