package interfaces

import "joel/onlineChess/matchState/models/game"

type Movable interface {
	Move(pos game.Position) (game.Position, error)
	GetPossibleMovements(board Board) []game.Movement
}
