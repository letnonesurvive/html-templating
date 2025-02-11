package main

import "net/http"

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
