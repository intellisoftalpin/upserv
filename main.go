package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	uploadPath := os.Getenv("UPLOAD_PATH")
	username := os.Getenv("BASIC_AUTH_USERNAME")
	password := os.Getenv("BASIC_AUTH_PASSWORD")

	if uploadPath == "" || username == "" || password == "" {
		fmt.Println("Environment variables UPLOAD_PATH, BASIC_AUTH_USERNAME, and BASIC_AUTH_PASSWORD must be set")
		os.Exit(1)
	}

	server := &UpServ{
		UploadPath: uploadPath,
		Username:   username,
		Password:   password,
	}

	http.Handle("/", server.BasicAuth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			server.UploadHandler(w, r)
		} else if r.Method == http.MethodGet {
			server.ListFilesHandler(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
