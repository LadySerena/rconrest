package main

import (
	"github.com/LadySerena/rconrest/pkg/players"
	"log"
	"net/http"
)

func main() {
	playerHandler := players.Players
	log.Println("listening on 8080")
	http.HandleFunc("/players/", playerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
