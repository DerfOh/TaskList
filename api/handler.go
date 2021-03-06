package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Handler for all requests that aren't on a specific endpoint
func Handler(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "text/html")
	webpage, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(response, fmt.Sprintf("index.html file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}
