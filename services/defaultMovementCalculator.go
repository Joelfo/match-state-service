package services

import "github.com/joelfo/match-state-service/models/game"

type DefaultMovementCalculator struct {
}

var boardSize = 8

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

func (calculator *DefaultMovementCalculator) GetBishopPossibleMovements(bishop *game.Piece, board *game.Board) []*game.Movement {
	currentPos := board.GetPiecePosition(bishop)

	var possibleMovements []*game.Movement

	directions := []struct{ dx, dy int }{
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	for _, direction := range directions {
		for i := 1; i < boardSize; i++ {
			xPos := currentPos.X + direction.dx*i
			yPos := currentPos.Y + direction.dy*i

			if xPos > boardSize || yPos > boardSize || xPos < 0 || yPos < 0 {
				break
			}

			movement := board.GetMovementTo(bishop, xPos, yPos)

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
