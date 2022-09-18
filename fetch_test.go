package GoTools

import (
	"testing"
)

func TestFetch(t *testing.T) {
	SimpleLogger("INFO", "Testing Fetch()...")

	Fetch(FetchOption{
		Method: "GET",
		URL:    "https://checkip.amazonaws.com",
	})

	Fetch(FetchOption{
		Method: "GET",
		URL:    "https://www.google.com",
	})

	SendSlackMessage("Golang Fetch test - POST slack")
}
