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

## How to test

1. Create file with name convention `<...>_test.go`
1. Import `"testing"`, then writing the function as `func TestFetch(t *testing.T) ...`
1. Run:

```sh
export SLACK_WEBHOOK_URL=...
DEBUG=true go test
```

*Do not change package name to `main` and run `go run fetch_test.go`, use `go test` instead*.

## How to publish

```sh
// commit and push to main first

VERSION=v0.1.6
git tag ${VERSION}
git push origin ${VERSION}
GOPROXY=proxy.golang.org go list -m github.com/rivernews/GoTools@${VERSION}
```
