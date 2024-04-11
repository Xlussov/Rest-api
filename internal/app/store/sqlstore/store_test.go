package sqlstore_test

import (
	"os"
	"testing"
)

var (
	dataBaseURL string
)

func TestMain(m *testing.M) {
	dataBaseURL = os.Getenv("DATABASE_URL")

	if dataBaseURL == "" {
		dataBaseURL = "host=localhost dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}