package interfaces

type Board interface {
	GetPieceAt(x, y int) *Movable
}
