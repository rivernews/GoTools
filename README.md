# GoTools

Go library for common utilities used by cloud applications.

## How to install

```sh
go get -u github.com/rivernews/GoTools
```

## What's included

Send a slack message:

```golang
// Remember to set environment variable `SLACK_WEBHOOK_URL`
SendSlackMessage("Send a slack message")
```

Make HTTP request:

```golang
Fetch(FetchOption{
    URL: "https://example.com",
    Method: "POST",
    PostData: map[string]string{
        "text": message
    },
})
```

Logger with emoji icons:

```golang
// Below prints `ℹ️ INFO: This is a info log`
Logger("INFO", "This", "is", "a", "info log")

// For Logger("ERROR", ...) it will also send a slack message to you
```