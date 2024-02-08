package server

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"qarwett/internal/config"
	"qarwett/pkg/logger/sl"
	"strings"
)

type Server struct {
	config *config.Config
	log    *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) *Server {
	return &Server{
		config: cfg,
		log:    log,
	}
}

func (s *Server) Run() {
	log := s.log.With(slog.String("service", "app.Server"))

	http.HandleFunc("/calendars/", func(w http.ResponseWriter, r *http.Request) {
		log.Info("request handled", slog.String("url", r.URL.String()))
		parts := strings.Split(r.URL.String(), "/")
		filename := parts[len(parts)-1]

		filePath := fmt.Sprintf("./calendars/%s.ics", filename)

		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Info("file not found", slog.String("path", filePath))
			http.Error(w, "Calendar not found", http.StatusNotFound)
			return
		}

		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			log.Error("error while reading file", slog.String("path", filePath))
			http.Error(w, "Error while reading while", http.StatusInternalServerError)
			return
		}

		headerContentDisposition := fmt.Sprintf("attachment; filename=\"%s\"", filePath)
		w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
		w.Header().Set("Content-Disposition", headerContentDisposition)

		if _, err = w.Write(fileContent); err != nil {
			log.Error("error while sending file", slog.String("path", filePath))
			http.Error(w, "Error while sending file", http.StatusInternalServerError)
			return
		}
	})

	err := http.ListenAndServe(s.config.ICalendar.Server.Addr, nil)
	if err != nil {
		log.Error("server crashes", sl.Err(err))
	}
}
