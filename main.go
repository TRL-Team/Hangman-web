package main

import (
	"net/http"
	"text/template"
)

type Firstpage struct {
	Title string
}

type gamePage struct {
	Title    string
	Subtitle string
}

func main() {

	template := template.Must(template.ParseFiles("./index.html"))

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	img := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images", img))

	js := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js", js))

	http.HandleFunc("/jeu", game)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Firstpage{
			Title: "Hangman Games",
		}
		template.Execute(w, page)
	})

	http.ListenAndServe(":80", nil)

}

func game(w http.ResponseWriter, r *http.Request) {
	template2 := template.Must(template.ParseFiles("./easy.html"))

	page2 := gamePage{
		Title:    "Play",
		Subtitle: "Coucou",
	}
	template2.Execute(w, page2)
}
