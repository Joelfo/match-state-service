package utils

import (
	"github.com/joelfo/match-state-service/models/game"
	"github.com/joelfo/match-state-service/models/pieceType"
	"github.com/joelfo/match-state-service/models/teamColors"
)

type PiecePos struct {
	game.Piece
	CurrentPos game.Position
}

func SetupBoard(board *game.Board) {
	whitePieces := getPiecesForTeam(teamColors.White)
	blackPieces := getPiecesForTeam(teamColors.Black)

	allpieces := append(whitePieces, blackPieces...)
	if board == nil {
		return
	}
	for i := range allpieces {
		piece := &allpieces[i]
		(*board).SetPieceAt(&piece.Piece, piece.CurrentPos.X, piece.CurrentPos.Y)
	}
}

func getPiecesForTeam(teamColor string) []PiecePos {
	var yPos int

	if teamColor == teamColors.White {
		yPos = 0
	} else {
		yPos = 7
	}

	var allpieces []PiecePos

	king := PiecePos{CurrentPos: game.Position{X: 4, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Knight}}
	allpieces = append(allpieces, king)

	queen := PiecePos{CurrentPos: game.Position{X: 3, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Queen}}
	allpieces = append(allpieces, queen)

	bishops := []PiecePos{
		PiecePos{CurrentPos: game.Position{X: 2, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Bishop}},
		PiecePos{CurrentPos: game.Position{X: 5, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Bishop}},
	}
	allpieces = append(allpieces, bishops...)

	knights := []PiecePos{
		PiecePos{CurrentPos: game.Position{X: 1, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Knight}},
		PiecePos{CurrentPos: game.Position{X: 6, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Knight}},
	}
	allpieces = append(allpieces, knights...)

	rooks := []PiecePos{
		{CurrentPos: game.Position{X: 0, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Rook}},
		{CurrentPos: game.Position{X: 7, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Rook}},
	}
	allpieces = append(allpieces, rooks...)

	if teamColor == teamColors.Black {
		yPos -= 1
	} else {
		yPos += 1
	}

	var pawns []PiecePos
	for i := 0; i < 8; i++ {
		pawns = append(pawns, PiecePos{CurrentPos: game.Position{X: i, Y: yPos}, Piece: game.Piece{Color: teamColor, Type: pieceType.Pawn}})
	}
	allpieces = append(allpieces, pawns...)

	return allpieces
}
