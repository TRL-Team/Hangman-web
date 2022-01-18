package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Firstpage struct {
	Valeur string
}

func main() {
	templates := template.Must(template.ParseFiles("./index.html"))
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		Page := Firstpage{
			Valeur: "hangman trl",
		}
		fmt.Println(Page.Valeur)
		templates.Execute(w, Page)
	})
	http.ListenAndServe(":3634", nil)
}

type Secondpage struct {
	valeurs string
}

func facile() {

	templates := template.Must(template.ParseFiles("./indexx.html"))
	fs := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		Pages := Secondpage{
			valeurs: "hangman trl",
		}
		templates.Execute(w, Pages)
	})

	http.ListenAndServe(":3634", nil)
}
