// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"time"
// )

// func main() {
// 	words, err := os.Open("./files/words.txt")
// 	if err != nil {
// 		log.Fatal("y'a r", err)
// 	}
// 	motsDuFichier := bufio.NewScanner(words)
// 	var mots []string
// 	for motsDuFichier.Scan() {
// 		mots = append(mots, motsDuFichier.Text())
// 	}
// 	rand.Seed(time.Now().UTC().UnixNano())
// 	chiffre := rand.Intn(len(mots))
// 	fmt.Println(mots[chiffre])
// }
