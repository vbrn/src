package main

import (
	"fmt"
	"net/http"
	"strconv"
	"math/rand"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {http.ServeFile(w, r, "server6_templates/post.html")})
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("username")
		age := r.FormValue("age")
		age_int, _ := strconv.ParseInt(age, 10, 64)//converst string to int
		rand_int := rand.Intn(50)
		
		fmt.Fprintf(w, "Привет %s Возрастом %s лет\n через %v лет тебе будет %v", name, age, rand_int, age_int+int64(rand_int))
	})
	
	fmt.Println("Listening on localhost:8181 ...")
	http.ListenAndServe(":8181", nil)

}
