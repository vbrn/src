package main


/* IN TEMPLATES
    eq: возвращает true, если два значения равны
    ne: возвращает true, если два значения НЕ равны
    le: возвращает true, если первое значение меньше или равно второму
    gt: возвращает true, если первое значение больше второго
    ge: возвращает true, если первое значение больше или равно второму
Кроме того, есть ряд операторов, которые аналогичны логическим операторам:
    and: возвращает true, если два выражения равны true
    or: возвращает true, если хотя бы одно из двух выражений равно true
    not: возвращает true, если выражение возвращает false
EXAMPLEs:  {{if gt .Age 20}} Старикашка{{end}} Старикашка если больше 20
 {{if or (gt 2 1) (lt 5 7)}} <p>Первый вариант</p> {{else}} <p>Второй вариант</p>{{end}}
*/

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
)

type ViewData struct {
	Hour_now int
	Title string
	Users []string
	Users_full[]User
}

type User struct{
    Name string
    Age int
    Gay bool
}

func main() {
	data := ViewData{
	Hour_now: time.Now().Hour(),
	Title: "Users list",
	Users: []string{"tom", "bob", "sam",},
	Users_full : []User{
            User{Name: "Tom", Age: 32, Gay: false},
            User{Name: "Miscevige", Age: 23, Gay: true},
            User{Name: "Hubbard", Age: 15, Gay: true},},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("server8_templates/index.html")
		tmpl.Execute(w, data)
	})
	
	fmt.Println("Server is listening on :8181")
	http.ListenAndServe(":8181", nil)

}
