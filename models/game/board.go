package game

type Board struct {
	positions [8][8]*BoardPosition
}

func (board Board) GetPieceAt(x, y int) *Piece {
	if x > 8 || y > 8 {
		return nil
	}

	position := board.positions[y][x]

	return position.Piece
}

func (board *Board) SetPieceAt(piece *Piece, x, y int) {
	if x > 8 || y > 8 {
		return
	}

	position := board.positions[y][x]

	position.Piece = piece
}

func NewEmptyBoard() *Board {
	board := &Board{}
	for i := 0; i < 8; i++ {
		var row [8]*BoardPosition
		for j := 0; j < 8; j++ {
			row[j] = &BoardPosition{
				Position: &Position{
					Y: i, X: j,
				},
			}
		}
		board.positions[i] = row
	}
	return board
}

func (board *Board) GetPiecePosition(piece *Piece) *Position {
	for _, row := range board.positions {
		for _, cell := range row {
			if cell.Piece == piece {
				return cell.Position
			}
		}
	}
	return nil
}

func (board *Board) GetMovementTo(piece *Piece, x, y int) *Movement {

	if x > 7 || y > 7 || x < 0 || y < 0 {
		return nil
	}

	pieceAtPos := (*board).GetPieceAt(x, y)
	if pieceAtPos != nil && pieceAtPos.Color == piece.Color {
		return nil
	}

	return &Movement{CapturedPiece: pieceAtPos, FinalPos: Position{X: x, Y: y}}
}
