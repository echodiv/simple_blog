package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost port=5432 dbname=dev_db sslmode=disable user=dev_user password=secret12345"
	}
	os.Exit(m.Run())
}
