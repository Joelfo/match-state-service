package services

import (
	"errors"
	"testing"

	"github.com/joelfo/match-state-service/models/game"
	"github.com/joelfo/match-state-service/models/pieceType"
	"github.com/joelfo/match-state-service/models/teamColors"
	"github.com/joelfo/match-state-service/utils"
)

func getSetUpBoard(t *testing.T) *game.Board {
	t.Helper()
	board := game.NewEmptyBoard()
	utils.SetupBoard(board)
	return board
}

func getMovementCalculator() (MovementCalculator, error) {
	return nil, errors.New("Not implemented yet")
}

func TestGeralRulesOnPieceMovements(t *testing.T) {
	calculator, err := getMovementCalculator()

	if err != nil {
		t.Fatalf("Error during Movement Calculator instantiation: %f", err)
	}

	whiteRook := &game.Piece{Color: teamColors.White}
	board := game.NewEmptyBoard()
	board.SetPieceAt(whiteRook, 3, 4)

	whiteBishop := &game.Piece{Color: teamColors.White, Type: pieceType.Bishop}
	board.SetPieceAt(whiteBishop, 3, 1)

	blackPawn := &game.Piece{Color: teamColors.Black, Type: pieceType.Pawn}
	board.SetPieceAt(blackPawn, 5, 4)

	possibleMovements := calculator.GetRookPossibleMovements(whiteRook, board)

	for _, movement := range possibleMovements {
		if movement.CapturedPiece == whiteBishop {
			t.Errorf("White piece can not capture another white piece.")
		}
		if movement.FinalPos.X == 3 && movement.FinalPos.Y == 1 {
			t.Errorf("White piece can not move into another white piece position.")
		}
		if movement.FinalPos.X == 5 && movement.FinalPos.Y == 4 && movement.CapturedPiece != blackPawn {
			t.Errorf("White piece can not move to a black piece position without capturing it.")
		}
	}
}
