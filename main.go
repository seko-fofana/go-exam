package main

import (
	"fmt"
	"net/http"
)

func main() {
}

type Task struct {
	Description string
	Done        bool
}

var task []Task

func list(rw http.ResponseWriter, _ *http.Request) {
	fmt.Println("list")
}
func done(w http.ResponseWriter, r *http.Request) {
	fmt.Println("done")
}
func add(w http.ResponseWriter, r *http.Request) {
	fmt.Println("add")
}
func main() {

	http.HandleFunc("/", list)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	http.ListenAndServe(":8080", nil)
}
