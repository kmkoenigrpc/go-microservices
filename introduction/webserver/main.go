package main

import (
	"net/http"
)

// this is an example
// ordinarily want to name the package the same name as the folder containing this main.go
// in this case it would be nominal to name the package 'webserver'

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world!"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
