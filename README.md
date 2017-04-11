# httpkit
A golang library that contains a simple http request method called `HttpRequest`.

## Function


```go
func HttpRequest(method string, url string, headers map[string]string, payload []byte) (*http.Response, error)
```

## Parameters

**`method`:** The HTTP method to be used.

**`url`:** The request URL.

**`headers`:** A string-string map of the headers to be sent in the request.

**`payload`:** The request body.



## Sample Call

```go

	resp, err := httpkit.HttpRequest("POST", "https://graph.facebook.com/v2.6/me/messages?access_token="+os.Getenv("PAGE_ACCESS_TOKEN"), headers, jsonPayload)
	if err != nil {
		log.Println("Error sending reply" + err.Error())

		return
	}
	// Close the Body after using.
	defer resp.Body.Close()

```
