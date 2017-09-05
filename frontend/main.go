package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var currentPath string
var majorNameListJSON string

func main() {
	log.Println("(main) The Server starts.")
	// initialize logger
	logFile := logInit()
	defer logFile.Close()

	log.Println("(main) The Logger has been initialized.")

	majorNameListJSON = majorNameListInit()
	log.Println("(main) The majorNameListJSON has been generated.")

	log.Println("(main) Waiting request... ")

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/db/", dbHandler)
	http.HandleFunc("/db/getDataByMajor", sendDataByMajorHandler)
	http.HandleFunc("/db/getSubjectTable", sendSubjectTableHandler)

	http.HandleFunc("/getMajorNameList", sendMajorNameList)

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

	t := time.Now()
	currentTimeString := fmt.Sprintf("%d-%02d-%02dT%02dh_%02dm_%02ds", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	logFile, err := os.Create(currentPath + "/log/" + currentTimeString)
	if err != nil {
		panic(err)
	}

	// print log to console and logFile simultaneously
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// set default logger
	log.SetOutput(multiWriter)

	return logFile
}

func majorNameListInit() string {
	majorNameList, err := getMajorListFromDB()
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(majorNameList)
	if err != nil {
		panic(err)
	}
	return string(b)
}
