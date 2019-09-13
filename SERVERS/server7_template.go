package main
import (
    "fmt"
    "net/http"
    "html/template"
)
type ViewData struct{
 
    Title string
    Message string
}
func main() {
      
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
 
        data := ViewData{
            Title: "World Cup",
            Message: "FIFA will never regret it",
        }
        tmpl, _ := template.ParseFiles("server6_templates/index.html")
        
        tmpl.Execute(w, data)
    })
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}
