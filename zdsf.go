package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var motString []string

func main() {
	c := getCompleteWord()
	initializeLetters(c)
}

func getNewWord() string {
	words, err := os.Open("./files/words.txt")
	if err != nil {
		log.Fatal("y'a r", err)
	}
	motsDuFichier := bufio.NewScanner(words)
	var mots []string
	for motsDuFichier.Scan() {
		mots = append(mots, motsDuFichier.Text())
	}
	rand.Seed(time.Now().UTC().UnixNano())
	chiffre := rand.Intn(len(mots))
	return (mots[chiffre])
}

func getCompleteWord() []string {
	mot := []rune(getNewWord())
	var motString []string
	for i := 0; i < len(mot); i++ {
		var conver string
		conver = string(mot[i] - 32)
		motString = append(motString, conver)
	}
	return motString
}
func initializeLetters(words []string) {
	var tabString []string
	for i := 0; i < len(words); i++ {
		tabString = append(tabString, "_ ")
	}

	// for i := 0; i < len(tab); i++ {
	// 	tabString[tab[i]] = words[tab[i]]
	// }
	fmt.Println(tabString)
	// for i := 0; i < len(tabString); i++ {
	// 	fmt.Printf(tabString[i])
	// }
}
