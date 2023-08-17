package main

import (
	"fmt"
	"net/http"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a GET request")
}

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a POST request")
}

func handlePutRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a PUT request")
}

func handleDeleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a DELETE request")
}

func main() {
	http.HandleFunc("/get", handleGetRequest)
	http.HandleFunc("/post", handlePostRequest)
	http.HandleFunc("/put", handlePutRequest)
	http.HandleFunc("/delete", handleDeleteRequest)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
