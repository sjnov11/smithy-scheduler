package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[len("/"):]

	var source []byte
	var err error

	// when request is root, send index.html
	// otherwise, send the file
	if r.URL.Path == "/" {
		source, err = ioutil.ReadFile("./index.html")
	} else {
		source, err = ioutil.ReadFile("./" + filename)
	}

	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, string(source))
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	// cmd := r.URL.Path[len("/db/"):]
	fmt.Println("DB Handler is called")

	switch r.Method {
	case "POST":
		fmt.Fprint(w, "server send message!")
	case "GET":
		fmt.Fprint(w, "get call has been arrived:"+r.URL.Path)
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
		return
	}

	b, _ := json.Marshal(subjects)
	fmt.Fprint(w, string(b))
}
