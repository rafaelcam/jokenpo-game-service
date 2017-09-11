package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"jokenpo-game-service/service"
	"fmt"
)

/*
/ 	Example call: http://localhost:3000/api/v1/play?player1=PAPER;player2=STONE
*/
func main() {
	router := mux.NewRouter().StrictSlash(true)

	sub := router.PathPrefix("/api/v1").Subrouter()
	sub.Methods("GET").Path("/play").HandlerFunc(service.Play)

	fmt.Println("Listening in http://localhost:3000/ ...")
	log.Fatal(http.ListenAndServe(":3000", router))
}