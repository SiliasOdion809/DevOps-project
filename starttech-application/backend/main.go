package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Backend API is running 🚀")
}

func main() {
	http.HandleFunc("/", healthHandler)

	fmt.Println("Server running on port 8080")

	http.ListenAndServe(":8080", nil)
}