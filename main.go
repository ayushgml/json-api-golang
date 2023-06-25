package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/user", makeHttpHandler(handleGetUserById))
	http.ListenAndServe(":3000", nil)
}

