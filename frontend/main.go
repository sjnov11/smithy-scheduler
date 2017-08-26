package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var currentPath string

func main() {
	// initialize logger
	logFile := logInit()
	defer logFile.Close()

	log.Println("(main) Server Start")

	http.HandleFunc("/", handler)
	http.HandleFunc("/db/", dbHandler)
	http.HandleFunc("/db/getDataByMajor", sendDataByMajorHandler)
	http.HandleFunc("/db/getSubjectTable", sendSubjectTableHandler)
	http.ListenAndServe(":8080", nil)
}

func logInit() *os.File {
	currentPath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}

	// setting logger
	logDir := "log"
	_, err = os.Stat(logDir)
	if err != nil {
		// directory is not exist
		err = os.Mkdir("log", 0755)
		if err != nil {
			panic(err)
		}
	}

	logFile, err := os.Create(currentPath + "/log/" + time.Now().String())
	if err != nil {
		panic(err)
	}

	// print log to console and logFile simultaneously
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// set default logger
	log.SetOutput(multiWriter)

	return logFile
}
