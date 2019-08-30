//go get github.com/gorilla/mux

/*   gorilla/context: предназначен для создания глобальных переменных из тела запроса

    gorilla/rpc: представляет реализацию протокола RPC-JSON

    gorilla/websocket: реализует протокол WebSocket

    gorilla/schema: позволяет создавать из значений формы единую структуру

    gorilla/securecookie: позволяет создавать зашифрованные куки, которые применяются при аутентификации

    gorilla/sessions: обеспечивает поддержку сессий

    gorilla/mux: позволяет определять более сложные маршруты, которые могут использовать регулярные выражения

    gorilla/reverse: используется для создания регулярных выражений для маршрутов*/


package main
import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)
 
func productsHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    response := fmt.Sprintf("id=%s", id)
    fmt.Fprint(w, response)
}
func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Index Page")
}
 
func main() {
      
    router := mux.NewRouter()
    router.HandleFunc("/products/{id:[0-9]+}", productsHandler)
    router.HandleFunc("/articles/{id:[0-9]+}", productsHandler)
    router.HandleFunc("/", indexHandler)
    http.Handle("/",router)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}
