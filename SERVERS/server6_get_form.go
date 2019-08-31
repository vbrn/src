//localhost:8181/user?name=Sam&age=21

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		age := r.URL.Query().Get("age")
		fmt.Fprintf(w, "Имя :%s Возраст :%s", name, age)
	})
	fmt.Println("Server is listening on localhost:8181 ...")
	http.ListenAndServe(":8181", nil)
}
