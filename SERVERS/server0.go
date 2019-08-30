package main

import ("fmt"; "net/http")

type msg string
func (m msg) ServeHTTP(resp http.ResponseWriter, req *http.Request) {fmt.Fprint(resp, m)}

func main() {
	msgHandler := msg("Hello its me and I am server 0")
	fmt.Println("Server is listening on localhost:8181")
	http.ListenAndServe("localhost:8181", msgHandler)
}
