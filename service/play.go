package service

import (
	"net/http"
	"jokenpo-game-service/model"
	"jokenpo-game-service/util"
)

func Play(writer http.ResponseWriter, request *http.Request) {


	player1 := model.Player{Motion : request.FormValue("player1")}.FillCode()
	player2 := model.Player{Motion : request.FormValue("player2")}.FillCode()

	if codeMotionsAreNotValid(player1.CodeMotion, player2.CodeMotion) {
		http.Error(writer, "Some motion was invalid. Try again", http.StatusInternalServerError)
		return
	}

	round := model.Round{Player1: player1, Player2: player2}
	round.Start()

	jsonBytes := util.ConvertObjectToJson(round, writer)
	util.WriteJsonResponse(writer, jsonBytes)
}

func codeMotionsAreNotValid(codeMotionPlayer1 int, codeMotionPlayer2 int) bool {
	if codeMotionPlayer1 == 0 || codeMotionPlayer2 == 0 {
		return true
	}

	return false;
}