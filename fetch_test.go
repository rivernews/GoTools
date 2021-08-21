package GoTools

import (
	"testing"
)

func TestFetch(t *testing.T) {
	SimpleLogger("INFO", "Starting")

	_, res, _ := Fetch(FetchOption{
		Method: "GET",
		URL: "https://google.com",
	})

	SimpleLogger("INFO", "Res is " + res)
}
