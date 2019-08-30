package main

import (
	"fmt"
	"net/http"
)

func http_page(page string, html_source string) {
	http.HandleFunc(page, func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, html_source)})
}

func main() {
	http_page("/", "server3HTML/index.html")
	http_page("/about", "server3HTML/about.html")
	http_page("/contact", "server3HTML/contact.html")

	fmt.Println("Server is listening on localhost: 8181...")
	http.ListenAndServe(":8181", nil)
	



}
