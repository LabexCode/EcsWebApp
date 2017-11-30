package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, nil)
}
func main() {
	http.HandleFunc("/", IndexHandler)
	fmt.Println(http.ListenAndServe(":8888", nil))
}
