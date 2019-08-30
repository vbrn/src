package main

import (
	"fmt"
	"net/http"
	"SERVERS/my_framework"
)



func main() {
	my_framework.Routing_html("/", "server3HTML/index.html")
	my_framework.Routing_html("/about", "server3HTML/about.html")
	my_framework.Routing_html("/contact", "server3HTML/contact.html")

	fmt.Println("Server is listening on localhost: 8283...")
	http.ListenAndServe(":8283", nil)
	



}
