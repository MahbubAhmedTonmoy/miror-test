package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
}

func greeter() {
	fmt.Println("hey there mod user")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> welcome to golang </h1>"))
}
