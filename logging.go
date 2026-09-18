package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func startServer(handler http.Handler) {
	logFile, err := setupLogging()
	if err != nil {
		log.Fatalf("failed to set up logging: %v", err)
	}
	defer logFile.Close()

	log.Println("starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}

func setupLogging() (*os.File, error) {
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(
		filepath.Join("logs", "app.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return nil, err
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)

	return logFile, nil
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *statusRecorder) Write(body []byte) (int, error) {
	if r.status == 0 {
		r.status = http.StatusOK
	}

	return r.ResponseWriter.Write(body)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &statusRecorder{ResponseWriter: w}

		next.ServeHTTP(recorder, r)

		if recorder.status == 0 {
			recorder.status = http.StatusOK
		}
		log.Printf("%s %s %d %s", r.Method, r.URL.RequestURI(), recorder.status, time.Since(start))
	})
}
