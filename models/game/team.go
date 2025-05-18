package game

type Team struct {
	Pieces         [16]Piece
	CapturedPieces []Piece
	Color          int
}
