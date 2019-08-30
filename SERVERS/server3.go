package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "server3HTML/index.html")})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "server3HTML/about.html")})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "server3HTML/contact.html")})

	fmt.Println("Server is listening on localhost: 8282...")
	http.ListenAndServe(":8282", nil)
	



}
