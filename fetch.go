package GoTools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	postDataBuffer := new(bytes.Buffer)
	if option.PostData != nil {
		postDataMap := option.PostData
		json.NewEncoder(postDataBuffer).Encode(postDataMap)
	} else {
		postDataMap := map[string]string{}
		json.NewEncoder(postDataBuffer).Encode(postDataMap)
	}

	// prepare headers
	headers := map[string][]string{}
	if option.Headers != nil {
		for k, v := range option.Headers {
			headers[k] = v
		}
	}

	// append request config and make request
	req, _ := http.NewRequest(option.Method, requestURL.String(), postDataBuffer)
	req.Header = headers
	client := &http.Client{}
	res, fetchErr := client.Do(req)

	var bytesContent []byte
	if fetchErr == nil {
		defer res.Body.Close()
		bytesContent, _ = ioutil.ReadAll(res.Body)
	} else {
		SimpleLogger("WARN", "Fetch error:" + fetchErr.Error())
	}

	// log response
	var responseMessage strings.Builder
	if !option.DisableHumanMessage {
		responseMessage.WriteString("Response:\n```\n")

		if bytesContent != nil {
			responseMessage.WriteString(string(bytesContent))
		} else {
			responseMessage.WriteString("Empty content")
		}

		responseMessage.WriteString("\n```\nAny error:\n```\n")
		if fetchErr != nil {
			responseMessage.WriteString("🔴 ")
			responseMessage.WriteString(fetchErr.Error())
		} else {
			responseMessage.WriteString("🟢 No error")
		}
		responseMessage.WriteString("\n```\n")
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