package _030102_httptestPackage_

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func DemoUseTestServer() {
	fmt.Println("\n030102 Basics: Package `httptest` — 02 Use Test Server")

	testServer := httptest.NewServer(http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.WriteHeader(http.StatusNotFound)
		io.WriteString(responseWriter, "This is a test response.")
	}))
	defer testServer.Close()

	response, _ := http.Get(testServer.URL)
	responseBody, _ := io.ReadAll(response.Body)

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("response body:", string(responseBody))
}
