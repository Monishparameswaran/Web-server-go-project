package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		fmt.Fprintf(w, "error invalid path")
		return
	}
	if r.Method != "GET" {
		fmt.Fprintf(w, "invalid method")
		return
	}
	fmt.Fprintf(w, "Hello user")
	return
}
func formhandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "the error occured %s", err)
	}
	name := r.FormValue("name")
	addr := r.FormValue("address")
	fmt.Fprintf(w, "hi %s and your are at the location %s", name, addr)

}
func main() {

	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("server started its process")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
