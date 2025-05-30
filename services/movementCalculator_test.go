package services

import (
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
	return &DefaultMovementCalculator{}, nil
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

func TestGetRookPossibleMovements(t *testing.T) {
	calculator, err := getMovementCalculator()

	if err != nil {
		t.Fatalf("Error during Movement Calculator instantiation: %f", err)
	}

	board := getSetUpBoard(t)
	if board == nil {
		t.Errorf("Board is nil")
		return
	}
	whiteRook := (*board).GetPieceAt(0, 0)
	possibleMovements := calculator.GetRookPossibleMovements(whiteRook, board)

	if len(possibleMovements) > 0 {
		t.Errorf("Rook possible moviments in board just set up are wrong")
	}

	board = game.NewEmptyBoard()
	board.SetPieceAt(whiteRook, 2, 3)

	whiteBishop := &game.Piece{Color: teamColors.White, Type: pieceType.Bishop}
	board.SetPieceAt(whiteBishop, 2, 0)

	blackPawn := &game.Piece{Color: teamColors.Black, Type: pieceType.Pawn}
	board.SetPieceAt(blackPawn, 4, 3)

	whiteKnight := &game.Piece{Color: teamColors.White, Type: pieceType.Knight}
	board.SetPieceAt(whiteKnight, 1, 3)

	possibleMovements = calculator.GetRookPossibleMovements(whiteRook, board)

	for _, movement := range possibleMovements {
		if movement.FinalPos.Y != 3 && movement.FinalPos.X != 2 {
			t.Errorf("Rook does not respect rook movements limitation")
		}
		if movement.FinalPos.Y == 3 && movement.FinalPos.X == 0 {
			t.Errorf("Rook can't move through white pieces.")
		}
		if movement.FinalPos.Y == 3 && movement.FinalPos.X > 4 {
			t.Errorf("Rook can't move through black pieces.")
		}
	}

	if len(possibleMovements) < 8 {
		t.Errorf("Rook calculated movements are less than real possible movements.")
	} else if len(possibleMovements) > 8 {
		t.Errorf("Rook calculated movements are more than real possible movements.")
	}

}

func TestGetBishopPossibleMovements(t *testing.T) {
	calculator, err := getMovementCalculator()

	if err != nil {
		t.Fatalf("Error during Movement Calculator instantiation: %f", err)
	}

	board := game.NewEmptyBoard()

	whiteBishop := &game.Piece{Type: pieceType.Bishop, Color: teamColors.White}
	board.SetPieceAt(whiteBishop, 3, 5)

	whiteRook := &game.Piece{Type: pieceType.Rook, Color: teamColors.White}
	board.SetPieceAt(whiteRook, 1, 3)

	blackQueen := &game.Piece{Type: pieceType.Queen, Color: teamColors.Black}
	board.SetPieceAt(blackQueen, 5, 7)

	blackRook := &game.Piece{Type: pieceType.Bishop, Color: teamColors.Black}
	board.SetPieceAt(blackRook, 6, 2)

	possibleMovements := calculator.GetBishopPossibleMovements(whiteBishop, board)

	if len(possibleMovements) != 8 {
		t.Errorf("Bishop possible movements are not in the same number as real possible movements.")
	}

	currentWhiteBishopPos := game.Position{X: 3, Y: 5}
	for _, possibleMovement := range possibleMovements {
		if possibleMovement.FinalPos.X == currentWhiteBishopPos.X || possibleMovement.FinalPos.Y == currentWhiteBishopPos.Y {
			t.Errorf("Bishop is not able to move vertically or horizontally.")
		}
		if (possibleMovement.FinalPos.X == 0 && possibleMovement.FinalPos.Y == 2) || (possibleMovement.FinalPos.X == 7 && possibleMovement.FinalPos.Y == 7) {
			t.Errorf("Bishop can not move through other pieces")
		}
	}
}
