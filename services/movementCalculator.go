package services

import (
	"github.com/joelfo/match-state-service/models/game"
)

type MovementCalculator interface {
	GetRookPossibleMovements(rook *game.Piece, board *game.Board) []*game.Movement
}
