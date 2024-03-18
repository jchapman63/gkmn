package server

import "testing"

// container must be running locally for this to work ? or maybe not ?
// this test passes when the container isn't running ? PGADMIN will not connect
func TestDataBaseConnection(t *testing.T) {
	TestConnection()
}
