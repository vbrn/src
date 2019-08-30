package main

import ("fmt"; "net/http" ; "io")


func index(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type", 
        "text/html",
    )
    io.WriteString(
        res, 
        `<doctype html>
<html>
    <head>
        <title>Main</title>
    </head>
    <body>
    	Function main<br>
        <a href="/hello">Get Hello</a>
    </body>
</html>`,
    )
}



func hello(res http.ResponseWriter, req *http.Request) {
    res.Header().Set(
        "Content-Type", 
        "text/html",
    )
    io.WriteString(
        res, 
        `<doctype html>
<html>
    <head>
        <title>Hello Zoomie world</title>
    </head>
    <body>
        Hello ZOOMIE World!<br>
        <a href="/">Back to Home</a>
    </body>
</html>`,
    )
}


func main() {
    http.HandleFunc("/", index)
    http.HandleFunc("/hello", hello)
    fmt.Println("Listening on port 127.0.0.1:9000")
    http.ListenAndServe(":9000", nil)
}
