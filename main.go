package main

import (
	"net/http"
	"text/template"
)

type Firstpage struct {
	Title string
}
type Gamepage struct {
	Title    string
	subtitle string
}

func main() {

	template := template.Must(template.ParseFiles("./index.html"))

	css := http.FileServer(http.Dir("./css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	img := http.FileServer(http.Dir("./images"))
	http.Handle("/images/", http.StripPrefix("/images", img))

	js := http.FileServer(http.Dir("./js"))
	http.Handle("/js/", http.StripPrefix("/js", js))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		page := Firstpage{
			Title: "Hagman Games",
		}
		template.Execute(w, page)
	})

	http.ListenAndServe(":80", nil)

}

func game() {
	template2 := template.Must(template.ParseFiles("./jeux.html"))

	http.HandleFunc("/jeux", func(w http.ResponseWriter, r *http.Request) {
		page2 := Gamepage{
			Title:    "Play",
			subtitle: "Vous êtes dans le mode ...",
		}
		template2.Execute(w, page2)
	})
}
