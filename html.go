package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func (s *UpServ) ListFilesHandler(w http.ResponseWriter, _ *http.Request) {
	files, err := filepath.Glob(filepath.Join(s.UploadPath, "*.apk"))
	if err != nil {
		http.Error(w, "Error listing files", http.StatusInternalServerError)
		return
	}

	type FileInfo struct {
		Name    string
		ModTime time.Time
	}

	var fileInfos []FileInfo
	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			continue
		}
		fileInfos = append(fileInfos, FileInfo{
			Name:    filepath.Base(file),
			ModTime: info.ModTime(),
		})
	}

	// Sort files by modification time in descending order
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime.After(fileInfos[j].ModTime)
	})

	tmpl, err := template.New("files").Parse(CHTMLTemplate)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, fileInfos)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}
