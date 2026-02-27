package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

var mutex sync.Mutex

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	IP        string `json:"ip"`
	Path      string `json:"path"`
	Level     string `json:"level"`
	Message   string `json:"message,omitempty"`
}

const logFilePath = "logs/app.log"
const maxLogSize = 1024 * 1024 // 1MB

func LogRequest(r *http.Request, level string, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	// Prepare log entry
	entry := LogEntry{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		IP:        r.RemoteAddr,
		Path:      r.URL.Path,
		Level:     level,
		Message:   message,
	}

	logData, _ := json.Marshal(entry)
	logData = append(logData, '\n') // add newline

	// Rotate log if size exceeds max
	if info, err := os.Stat(logFilePath); err == nil && info.Size() > maxLogSize {
		os.Rename(logFilePath, fmt.Sprintf("logs/app-%s.log", time.Now().Format("20060102-150405")))
	}

	// Open log file for appending
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer file.Close()

	// Write to file AND console
	multiWriter := io.MultiWriter(file, os.Stdout)
	multiWriter.Write(logData)
}
