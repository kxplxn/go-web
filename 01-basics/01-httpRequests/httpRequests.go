package _030101_httpRequests_

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Demo() {
	fmt.Println("\n030101 Basics: Http Requests")

	response, _ := http.Get("https://postman-echo.com/get")
	responseBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println("status code:", response.StatusCode)
	fmt.Println("response body:", string(responseBody))
}
