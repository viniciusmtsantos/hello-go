package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	u := usuario{"JUANZITO", "juanzitodaTI@gmail.com"}

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
		w.Write([]byte("Hello, World!!"))

	})

	fmt.Println("Escutando na Porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))

}
