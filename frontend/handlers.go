package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// when request is root, send index.html
	// otherwise, send the file

	path := r.URL.Path[len("/"):]

	if path == "" {
		log.Println("(rootHandler) The Main page access has been occurred.")

		source, err := ioutil.ReadFile("./webRoot/resources/index.html")
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Println("(rootHandler) ", err)
			return
		}
		fmt.Fprint(w, string(source))

	} else {
		// send requested file
		source, err := ioutil.ReadFile("./webRoot/" + path)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Println("(rootHandler) ", err)
			return
		}
		fmt.Fprint(w, string(source))
		// log.Println("(handler) The requested file has been sent: ", filename)
	}
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	// cmd := r.URL.Path[len("/db/"):]
	log.Println("(dbHandler) DB Handler is called")

	switch r.Method {
	case "POST":
		fmt.Fprint(w, "(dbHandler) server send message!")
	case "GET":
		fmt.Fprint(w, "(dbHandler) get call has been arrived:"+r.URL.Path)
	default:
		http.NotFound(w, r)
		return
	}
}

func sendDataByMajorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
	}

	// case is only "POST"
	decoder := json.NewDecoder(r.Body)
	var t struct {
		Major string `json:major`
	}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// get subject data by major name
	subjects, err := getDataFromDBByMajor(t.Major)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("(sendDataByMajorHandler) ", err)
		return
	}

	b, _ := json.Marshal(subjects)
	fmt.Fprint(w, string(b))
	log.Println("(sendDataByMajorHandler) The subject data have been sent.")
}

func sendSubjectTableHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.NotFound(w, r)
		return
	}

	// case is only "POST"
	decoder := json.NewDecoder(r.Body)
	var t struct {
		Major string `json:major`
	}
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// get subject data by major name
	subjects, err := getDataFromDBByMajor(t.Major)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("(sendSubjectTableHandler) ", err)
		return
	}

	// get html source of table
	tableSource, err := drawSubjectTable(subjects)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("(sendSubjectTableHandler) ", err)
		return
	}
	fmt.Fprint(w, tableSource)
	log.Println("(sendSubjectTableHandler) The subject table html code has been sent to browser.")
}

func sendMajorNameList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, majorNameListJSON)
}
