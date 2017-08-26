package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// when request is root, send index.html
	// otherwise, send the file
	if r.URL.Path == "/" {
		// send index.html
		// TODO: 사용자가 루트에 접속할 때마다 main page를 새로 그린다. 서버 main함수가 실행될 때 한 번만 그리고 요청은 그냥 html코드 보내줘도 될듯.
		err := writeMainPageHTML(w)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Println("(handler) ", err)
			return
		}
	} else {
		// send requested file
		filename := r.URL.Path[len("/"):]
		source, err := ioutil.ReadFile("./" + filename)
		if err != nil {
			http.Error(w, err.Error(), 500)
			log.Println("(handler) ", err)
			return
		}
		fmt.Fprint(w, string(source))
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
}

func sendSubjectTableHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Println("(sendSubjectTableHandler) ", err)
		return
	}

	// get html source of table
	tableSource, err := drawSubjectTable(subjects)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("(sendSubjectTableHandler) ", err)
	}
	fmt.Fprint(w, tableSource)
}
