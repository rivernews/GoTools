package GoTools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

// FetchOption - args for method `Fetch`
type FetchOption struct {
	QueryParams         map[string]string
	PostData            interface{}
	Headers             map[string][]string
	URL                 string
	Method              string
	DisableHumanMessage bool
	responseStore       interface{}
}

func logRequest(req *http.Request) {
	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		SimpleLogger("ERROR", err.Error())
	}
	SimpleLogger("VERBOSE", fmt.Sprintf("REQUEST:\n%s\nREQUEST END", bytes.TrimSpace(reqDump)))
}

func logResponse(res *http.Response) {
	resDump, err := httputil.DumpResponse(res, false)
	if err != nil {
		SimpleLogger("ERROR", err.Error())
	}
	SimpleLogger("VERBOSE", fmt.Sprintf("RESPONSE:\n%s\nRESPONSE END", bytes.TrimSpace(resDump)))
}

// Fetch - convenient method to make request with querystring and post data
func Fetch(option FetchOption) ([]byte, string, error) {
	requestURL, _ := url.Parse(option.URL)

	// prepare querystring
	params := url.Values{}
	if option.QueryParams != nil {
		for k, v := range option.QueryParams {
			params.Add(k, v)
		}
	}
	requestURL.RawQuery = params.Encode()

	// prepare post data
	var postData io.Reader
	if option.PostData != nil {
		postData = bytes.NewReader(AsJsonBytes(option.PostData))
	} else {
		postData = nil
	}

	// prepare headers
	headers := map[string][]string{}
	if option.Headers != nil {
		for k, v := range option.Headers {
			headers[k] = v
		}
	}

	// append request config and make request
	req, _ := http.NewRequest(option.Method, option.URL, postData)
	req.Header = headers
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", `Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.27 Safari/537.36`)
	}
	logRequest(req)

	client := &http.Client{}
	res, fetchErr := client.Do(req)

	var bytesContent []byte
	if fetchErr != nil {
		SimpleLogger("WARN", "Fetch error:"+fetchErr.Error())
	} else {
		briefMsg := fmt.Sprintf("%d %s, contentLength: %d", res.StatusCode, option.URL, res.ContentLength)
		if res.StatusCode >= 400 {
			SimpleLogger("WARN", briefMsg)
		} else {
			SimpleLogger("DEBUG", briefMsg)
		}
		bytesContent, _ = ioutil.ReadAll(res.Body)
	}
	if res != nil {
		logResponse(res)
		// No need to do `defer res.Body.Close()`
		// it's handled by client.Do already
		// https://stackoverflow.com/a/68851335/9814131
		res.Body.Close()
	} else {
		SimpleLogger("WARN", "Response object is nil")
	}

	// log response
	var responseMessage strings.Builder
	if !option.DisableHumanMessage {
		responseMessage.WriteString("Response:\n``` ")

		if bytesContent != nil {
			responseMessage.WriteString(string(bytesContent))
		} else {
			responseMessage.WriteString("(Empty content)")
		}

		responseMessage.WriteString(" ```\n")

		responseMessage.WriteString("Any error:\n``` ")
		if fetchErr != nil {
			responseMessage.WriteString("🔴 ")
			responseMessage.WriteString(fetchErr.Error())
		} else {
			responseMessage.WriteString("🟢 No error")
		}
		responseMessage.WriteString(" ```\n")
	}

	if fetchErr != nil {
		return bytesContent, responseMessage.String(), fetchErr
	}

	// parse response into struct if an interface is provided
	if option.responseStore != nil {
		unmarshalJSONErr := json.Unmarshal(bytesContent, option.responseStore)
		if unmarshalJSONErr != nil {
			Logger("INFO", "Failed to parse JSON response=", string(bytesContent))
			return bytesContent, responseMessage.String(), unmarshalJSONErr
		}
	}

	return bytesContent, responseMessage.String(), nil
}
