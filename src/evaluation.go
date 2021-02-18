package main

var pieceValues = []int{
	0,
	1,
	5,
	3,
	3,
	9,
	99999999,
}

func evaluatePosition(pos Position) int {
	if pos.isInsufficientMaterial() {
		return 0
	}

	result := 0

	result += pos.getPieceCount(WHITE_PAWN) * pieceValues[PIECE_PAWN]
	result += pos.getPieceCount(WHITE_ROOK) * pieceValues[PIECE_ROOK]
	result += pos.getPieceCount(WHITE_KNIGHT) * pieceValues[PIECE_KNIGHT]
	result += pos.getPieceCount(WHITE_BISHOP) * pieceValues[PIECE_BISHOP]
	result += pos.getPieceCount(WHITE_QUEEN) * pieceValues[PIECE_QUEEN]

	result -= pos.getPieceCount(BLACK_PAWN) * pieceValues[PIECE_PAWN]
	result -= pos.getPieceCount(BLACK_ROOK) * pieceValues[PIECE_ROOK]
	result -= pos.getPieceCount(BLACK_KNIGHT) * pieceValues[PIECE_KNIGHT]
	result -= pos.getPieceCount(BLACK_BISHOP) * pieceValues[PIECE_BISHOP]
	result -= pos.getPieceCount(BLACK_QUEEN) * pieceValues[PIECE_QUEEN]

	if pos.turn == SIDE_WHITE {
		return result
	} else {
		return result * -1
	}
}
