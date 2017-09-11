package model

import "jokenpo-game-service/config"

type Player struct {
	Motion  string `json:"motion"`
	CodeMotion int `json:"-"`
}

func (player Player) FillCode() Player {
	if player.Motion == "PAPER" {
		player.CodeMotion = config.PAPER_CODE
	} else if player.Motion == "STONE" {
		player.CodeMotion = config.STONE_CODE
	} else if player.Motion == "SCISSORS" {
		player.CodeMotion = config.SCISSORS_CODE
	}

	return player
}