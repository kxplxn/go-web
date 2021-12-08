package _030102_httptestPackage_

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func DemoTestHTTPCalls() {
	fmt.Println("\n030102 Basics: Package `httptest` — 01 Test HTTP Calls")

	testHTTPCall := func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "This is a test response.")
	}

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "http://testurl.com", nil)
	testHTTPCall(recorder, request)

	response := recorder.Result()
	responseBody, _ := io.ReadAll(response.Body)

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("response body:", string(responseBody))
}
