package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/calendars/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("INFO: request handled to: %s\n", r.URL.String())
		parts := strings.Split(r.URL.String(), "/")
		filename := parts[len(parts)-1]
		fmt.Println(filename)

		filePath := fmt.Sprintf("./.calendars/%s.ics", filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Printf("INFO: file not found in '%s'\n", filePath)
			http.Error(w, "Calendar not found", http.StatusNotFound)
			return
		}

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Printf("ERROR: error while reading file in '%s'\n", filePath)
			http.Error(w, "Error while reading while", http.StatusInternalServerError)
			return
		}

		headerContentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", filePath)
		w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
		w.Header().Set("Content-Disposition", headerContentDisposition)

		if _, err = w.Write(fileContent); err != nil {
			log.Printf("ERROR: error while sending file from '%s'\n", filePath)
			http.Error(w, "Error while sending file", http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatalln("ERROR: server crashes: ", err.Error())
	}
}
