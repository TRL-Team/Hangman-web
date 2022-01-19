package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// ┌────────────────────────────────────────────────────────────┐
// │ Globals				             						│
// └────────────────────────────────────────────────────────────┘

var state State // Sauvegarde l'état du jeu global (permet de garder les données entre les requêtes)

// ┌────────────────────────────────────────────────────────────┐
// │ Structs				             						│
// └────────────────────────────────────────────────────────────┘
// Représentation de l'état d'une lettre dans le jeu :
// Letter {Value : "a", Used : false} => La lettre a n'a pas été jouée.
// Letter {Value : "b", Used : true} => La lettre b a été jouée.

type Letter struct {
	Value string
	Used  bool
}

// State représente l'état global du jeu  :
type State struct {
	CompleteWord string   // C'est la solution du pendu en clair : Exemple :
	Letters      []Letter // Liste des 26 lettres
	Errors       int      // Nb d'erreurs
	CurrentWord  []string // état du mot en cours de partie : Exemple :
	GameOver     string   // Msg de fin de partie modifié selon l'état du jeu : "" (en cours), gagné, perdu
}

// ┌────────────────────────────────────────────────────────────┐
// │ Route handlers			             						│
// └────────────────────────────────────────────────────────────┘

func homeHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/index.html"))
	page.Execute(w, state)
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	page := template.Must(template.ParseFiles("views/game.html"))

	// ┌────────────────────────────────┐
	// │ Initiliaze the game			│
	// └────────────────────────────────┘
	switch r.Method {
	// On met en place ce qui va nous servir pour le jeu
	case http.MethodGet:
		word := getNewWord()                            // mot choisi aléatoirement pour le jeu
		state.CompleteWord = word                       // On sauvegarde le mot complet dans l'état global
		state.Letters = initializeLetters()             // On sauvegarde la liste des lettres non utilisées dans l'état global
		state.Errors = 0                                // On sauvegarde le nombre d'erreur dans l'état global
		state.CurrentWord = initializeCurrentWord(word) // On sauvegarde la version "cryptée" du mot dans l'état global
		state.GameOver = ""                             // On sauvegarde le message de fin de jeu dans l'état global
		page.Execute(w, state)

	// ┌────────────────────────────────┐
	// │ Read the letter sent by player │
	// └────────────────────────────────┘
	case http.MethodPost:
		r.ParseForm()
		letter := r.FormValue("letter")
		isError := true

		// Replace "_" with the letter from the player, if found
		for i, v := range state.CompleteWord {
			if string(v) == letter {
				isError = false
				state.CurrentWord[i] = letter
			}
			// If all letters from the word have been checked, the letter has not been found isError stays "true"
		}
		// Je désactive la lettre qui a été jouée en parcourant mon tableau d'alphabet
		for i, v := range state.Letters {
			if v.Value == letter {
				state.Letters[i] = Letter{Value: v.Value, Used: true}
				break
			}
		}

		if isError {
			state.Errors++
		}

		// quand on appelle la fonction, on appelle ce qu'elle renvoie... ;)
		switch isGameOver(state.CurrentWord, state.Errors) {
		case 2:
			for i, v := range state.Letters {
				state.Letters[i] = Letter{Value: v.Value, Used: true}
			}
			state.GameOver = "You lose! Game over"
			state.CurrentWord = getCompleteWord(state.CompleteWord)

		case 1:
			for i, v := range state.Letters {
				state.Letters[i] = Letter{Value: v.Value, Used: true}
			}
			state.GameOver = "You win! Game over"
		}

		page.Execute(w, state)
	}
}

// ┌────────────────────────────────────────────────────────────┐
// │ Utilities				             						│
// └────────────────────────────────────────────────────────────┘

/*
Récupère une ligne aléatoirement dans le fichier en .txt

Exemple :
"carotte"
*/
func getNewWord() string {

}

/*
Génère une liste de 26 lettres non utilisées.

Exemple :
[
	{Value : "a", Used: false},
	{Value : "b", Used: false},
	...,
	{Value : "z", Used: false}
]
*/
func initializeLetters() []Letter {}

/*
Génère les _ correspondants au mot /!\ Attention si espaces

Exemple :
word = Fifa
return => _ _ _ _
*/
func initializeCurrentWord(word string) []string {

	var pendu []rune
	for i := 0; i < len(bot); i++ {
		pendu = append(pendu, '_')
	}

	for i := 0; i < len(tab); i++ {
		pendu[tab[i]] = bot[tab[i]]
	}

	for i := 0; i < len(pendu); i++ {
		fmt.Printf("%c ", pendu[i])
	}

}

/*
Renvoie le mot sous forme d'un tableau de string

Exemple :
word = fifa
return => ["f","i","f","a"]
*/
func getCompleteWord(word string) []string {}

/*
Renvoie 0 si partie en cours, 1 si c'est gagné, 2 si perdu

Exemple :
word = ["f","_","f","a"]
errors = 3
return 0

Exemple 2 :
word = ["f","_","f","a"]
errors = 6
return 2

Exemple 3 :
word = ["f","i","f","a"]
errors = 2
return 1
*/
func isGameOver(word []string, errors int) int {}

// ┌────────────────────────────────────────────────────────────┐
// │ Main					             						│
// └────────────────────────────────────────────────────────────┘

func main() {

	// ┌────────────────────────────────┐
	// │ Serve static files				│
	// └────────────────────────────────┘
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static", fileserver))

	// ┌────────────────────────────────┐
	// │ Routes							│
	// └────────────────────────────────┘
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/play", playHandler)

	// ┌────────────────────────────────┐
	// │ Start the server				│
	// └────────────────────────────────┘
	http.ListenAndServe(":8080", nil)
}
