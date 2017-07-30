package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId, _ := strconv.Atoi(vars["todoId"])

    t := RepoFindTodo(todoId)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusFound)

    if err := json.NewEncoder(w).Encode(t); err != nil {
	panic(err)
    }
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    var todo Todo

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
	panic(err)
    }

    if err := r.Body.Close(); err != nil {
	panic(err)
    }

    if err := json.Unmarshal(body, &todo); err != nil {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(422)
	if err := json.NewEncoder(w).Encode(err); err != nil {
	    panic(err)
	}
    }

    t := RepoCreateTodo(todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
	panic(err)
    }
}