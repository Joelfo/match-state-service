package services

import "github.com/joelfo/match-state-service/models/game"

type DefaultMovementCalculator struct {
}

func (calculator *DefaultMovementCalculator) GetRookPossibleMovements(rook *game.Piece, board *game.Board) []*game.Movement {
	currentPos := board.GetPiecePosition(rook)

	var possibleMovements []*game.Movement

	possibleDirections := []int{1, -1}

	for _, dxDirection := range possibleDirections {
		for dx := 1; dx < 8; dx++ {
			posX := currentPos.X + dx*dxDirection
			if posX < 0 || posX > 7 {
				break
			}

			movement := board.GetMovementTo(rook, posX, currentPos.Y)

			if movement == nil {
				break
			}

			possibleMovements = append(possibleMovements, movement)

			if movement.CapturedPiece != nil {
				break
			}
		}
	}

	for _, dyDirection := range possibleDirections {
		for dy := 1; dy < 8; dy++ {
			posY := currentPos.Y + dy*dyDirection

			if posY < 0 || posY > 8 {
				break
			}

			movement := board.GetMovementTo(rook, currentPos.X, posY)

			if movement == nil {
				break
			}

			possibleMovements = append(possibleMovements, movement)

			if movement.CapturedPiece != nil {
				break
			}
		}
	}

	return possibleMovements
}
