package connections

import (
	"testing"
)

// All test functions are set up to use the databases in the docker compose.
func testMssqlCredsDBConnection(t *testing.T) {
	var addr string = "localhost"
	var unam string = "sa"
	var pwrd string = "ThisIsASecur3Pass!"
	var db string = "students"
	var port string = "1401"
	conn, err := mssqlCredsDBConnection(addr, unam, pwrd, db, port)

	if conn.Stats().OpenConnections < 1 || err != nil {
		t.Fatalf(`mssqlCredsDBConnection(addr,unam,pwrd,db,port)= %s Unable to make database connection`, "yeet")
	}

}
