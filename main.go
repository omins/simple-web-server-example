package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Student struct {
	Age  int
	Name string
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/foo", fooHandler)
	mux.HandleFunc("/student", studentHandler)
	return mux
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	name := queryValues.Get("name")
	id, _ := strconv.Atoi(queryValues.Get("id"))
	fmt.Fprintf(w, "Hello, %s with id %d", name, id)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func studentHandler(w http.ResponseWriter, r *http.Request) {
	student := Student{20, "John"}
	data, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func main() {
	http.ListenAndServe(":8080", MakeWebHandler())
}
