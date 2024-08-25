package utils

import (
	"math"

	"github.com/drod21/DishDuel-be/models"
)

func UpdateMMR(winner, loser *models.Restaurant) {
	kFactor := 32.0
	expectedScoreWinner := 1 / (1 + math.Pow(10, float64(loser.MMR-winner.MMR)/400))
	expectedScoreLoser := 1 - expectedScoreWinner

	winner.MMR += int(kFactor * (1 - expectedScoreWinner))
	loser.MMR += int(kFactor * (0 - expectedScoreLoser))
}
