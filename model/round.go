package model

type Round struct {
	Player1  Player `json:"player1"`
	Player2   Player `json:"player2"`
	Winner string `json:"winner"`
}

/*
/ 1 - Paper; 2 - Stone; 3 - Scissors;
/ Subtract the motion code from the player 1 of the player 2
/
/ If result is -1 or 2 Player1 Win
/ If result is 1 or -2 Player2 Win
/
*/
func (round *Round) Start() {
	if round.Player1.CodeMotion == round.Player2.CodeMotion {
		round.Winner = "Draw"
		return
	}

	result := round.Player1.CodeMotion - round.Player2.CodeMotion
	if(result == -1 || result == 2) {
		round.Winner = "Player 1"
	} else {
		round.Winner = "Player 2"
	}
}