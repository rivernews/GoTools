package GoTools

import (
	"testing"
)

func TestFetch(t *testing.T) {
	SimpleLogger("INFO", "Testing Fetch()...")

	_, res, _ := Fetch(FetchOption{
		Method: "GET",
		URL:    "https://checkip.amazonaws.com",
	})

	SimpleLogger("INFO", "Fetch test result - "+res)

	_, res, _ = Fetch(FetchOption{
		Method: "GET",
		URL:    "https://www.google.com",
	})

	SimpleLogger("INFO", "Fetch test result - "+res)
}
