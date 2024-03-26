package server

import (
	"testing"
)

func TestDataBaseConnection(t *testing.T) {
	TestConnection()
	t.Log("success!")
}
