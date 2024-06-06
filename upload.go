package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (s *UpServ) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 500<<20) // 500 MB
	if err := r.ParseMultipartForm(500 << 20); err != nil {
		http.Error(w, "File too big", http.StatusRequestEntityTooLarge)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Get the current time and format it
	currentTime := time.Now().Format("20060102_150405") // YYYYMMDD_HHMMSS

	// Extract the file extension
	extension := filepath.Ext(handler.Filename)
	name := handler.Filename[:len(handler.Filename)-len(extension)]

	// Construct the new filename with date/time postfix
	newFilename := fmt.Sprintf("%s_%s%s", name, currentTime, extension)
	path := filepath.Join(s.UploadPath, newFilename)

	dst, err := os.Create(path)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err = io.Copy(dst, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s", newFilename)
}
