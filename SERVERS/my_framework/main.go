package my_framework

import (
	"net/http"
)


func Routing_html(page string, html_source string) {
	http.HandleFunc(page, func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, html_source)})
}

//SHOULD BE make 
//go install 
//in the folder my_frameworks and then in the file you should reference as
//"SERVERS/my_framework"
//my_framework.Routing_html("/", "server3HTML/index.html")
