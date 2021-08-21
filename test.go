package main

func main() {
	SimpleLogger("INFO", "Starting")

	_, res, _ := Fetch(FetchOption{
		Method: "GET",
		URL: "https://api.shaungc.com",
	})

	SimpleLogger("INFO", "Res is " + res)
}
