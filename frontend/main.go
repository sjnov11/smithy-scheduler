package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/db/", dbHandler)
	http.HandleFunc("/db/getDataByMajor", sendDataByMajorHandler)
	http.HandleFunc("/db/getSubjectTable", sendSubjectTableHandler)
	http.ListenAndServe(":8080", nil)
}
