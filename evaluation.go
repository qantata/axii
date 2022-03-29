package main

var pieceValues = []int{
	0,
	100,
	500,
	300,
	300,
	900,
	900000,
}

var whitePawnSqTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, -20, -20, 10, 10, 5,
	5, -5, -10, 0, 0, -10, -5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, 5, 10, 25, 25, 10, 5, 5,
	10, 10, 20, 30, 30, 20, 10, 10,
	50, 50, 50, 50, 50, 50, 50, 50,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var blackPawnSqTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	50, 50, 50, 50, 50, 50, 50, 50,
	10, 10, 20, 30, 30, 20, 10, 10,
	5, 5, 10, 25, 25, 10, 5, 5,
	0, 0, 0, 20, 20, 0, 0, 0,
	5, -5, -10, 0, 0, -10, -5, 5,
	5, 10, 10, -20, -20, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var whiteKnightSqTable = []int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 0, 10, 15, 15, 10, 0, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

var blackKnightSqTable = []int{
	-50, -40, -30, -30, -30, -30, -40, -50,
	-40, -20, 0, 0, 0, 0, -20, -40,
	-30, 0, 10, 15, 15, 10, 0, -30,
	-30, 5, 15, 20, 20, 15, 5, -30,
	-30, 0, 15, 20, 20, 15, 0, -30,
	-30, 5, 10, 15, 15, 10, 5, -30,
	-40, -20, 0, 5, 5, 0, -20, -40,
	-50, -40, -30, -30, -30, -30, -40, -50,
}

var bishopSqTable = []int{
	-20, -10, -10, -10, -10, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 10, 10, 5, 0, -10,
	-10, 5, 5, 10, 10, 5, 5, -10,
	-10, 0, 10, 10, 10, 10, 0, -10,
	-10, 10, 10, 10, 10, 10, 10, -10,
	-10, 5, 0, 0, 0, 0, 5, -10,
	-20, -10, -10, -10, -10, -10, -10, -20,
}

var whiteRookSqTable = []int{
	0, 0, 0, 5, 5, 0, 0, 0,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	5, 10, 10, 10, 10, 10, 10, 5,
	0, 0, 0, 0, 0, 0, 0, 0,
}

var blackRookSqTable = []int{
	0, 0, 0, 0, 0, 0, 0, 0,
	5, 10, 10, 10, 10, 10, 10, 5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	-5, 0, 0, 0, 0, 0, 0, -5,
	0, 0, 0, 5, 5, 0, 0, 0,
}

var queenSqTable = []int{
	-20, -10, -10, -5, -5, -10, -10, -20,
	-10, 0, 0, 0, 0, 0, 0, -10,
	-10, 0, 5, 5, 5, 5, 0, -10,
	-5, 0, 5, 5, 5, 5, 0, -5,
	0, 0, 5, 5, 5, 5, 0, -5,
	-10, 5, 5, 5, 5, 5, 0, -10,
	-10, 0, 5, 0, 0, 0, 0, -10,
	-20, -10, -10, -5, -5, -10, -10, -20,
}

var whiteKingSqTable = []int{
	20, 30, 10, 0, 0, 10, 30, 20,
	20, 20, 0, 0, 0, 0, 20, 20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
}

var blackKingSqTable = []int{
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-30, -40, -40, -50, -50, -40, -40, -30,
	-20, -30, -30, -40, -40, -30, -30, -20,
	-10, -20, -20, -20, -20, -20, -20, -10,
	20, 20, 0, 0, 0, 0, 20, 20,
	20, 30, 10, 0, 0, 10, 30, 20,
}

func evaluatePosition(pos Position) int {
	if pos.isInsufficientMaterial() {
		return 0
	}

	whiteValue := 0
	blackValue := 0

	// White
	whitePawns := pos.getPiecesByTypeAndColor(PIECE_PAWN, SIDE_WHITE)
	for whitePawns != 0 {
		square := bbPopLSB(&whitePawns)
		whiteValue += pieceValues[PIECE_PAWN] + whitePawnSqTable[square]
	}

	whiteRooks := pos.getPiecesByTypeAndColor(PIECE_ROOK, SIDE_WHITE)
	for whiteRooks != 0 {
		square := bbPopLSB(&whiteRooks)
		whiteValue += pieceValues[PIECE_ROOK] + whiteRookSqTable[square]
	}

	whiteKnights := pos.getPiecesByTypeAndColor(PIECE_KNIGHT, SIDE_WHITE)
	for whiteKnights != 0 {
		square := bbPopLSB(&whiteKnights)
		whiteValue += pieceValues[PIECE_KNIGHT] + whiteKnightSqTable[square]
	}

	whiteBishops := pos.getPiecesByTypeAndColor(PIECE_BISHOP, SIDE_WHITE)
	for whiteBishops != 0 {
		square := bbPopLSB(&whiteBishops)
		whiteValue += pieceValues[PIECE_BISHOP] + bishopSqTable[square]
	}

	whiteQueens := pos.getPiecesByTypeAndColor(PIECE_QUEEN, SIDE_WHITE)
	for whiteQueens != 0 {
		square := bbPopLSB(&whiteQueens)
		whiteValue += pieceValues[PIECE_QUEEN] + queenSqTable[square]
	}

	whiteKing := pos.getPiecesByTypeAndColor(PIECE_KING, SIDE_WHITE)
	for whiteKing != 0 {
		square := bbPopLSB(&whiteKing)
		whiteValue += pieceValues[PIECE_KING] + whiteKingSqTable[square]
	}

	// Black
	blackPawns := pos.getPiecesByTypeAndColor(PIECE_PAWN, SIDE_BLACK)
	for blackPawns != 0 {
		square := bbPopLSB(&blackPawns)
		blackValue += pieceValues[PIECE_PAWN] + blackPawnSqTable[square]
	}

	blackRooks := pos.getPiecesByTypeAndColor(PIECE_ROOK, SIDE_BLACK)
	for blackRooks != 0 {
		square := bbPopLSB(&blackRooks)
		blackValue += pieceValues[PIECE_ROOK] + blackRookSqTable[square]
	}

	blackKnights := pos.getPiecesByTypeAndColor(PIECE_KNIGHT, SIDE_BLACK)
	for blackKnights != 0 {
		square := bbPopLSB(&blackKnights)
		blackValue += pieceValues[PIECE_KNIGHT] + blackKnightSqTable[square]
	}

	blackBishops := pos.getPiecesByTypeAndColor(PIECE_BISHOP, SIDE_BLACK)
	for blackBishops != 0 {
		square := bbPopLSB(&blackBishops)
		blackValue += pieceValues[PIECE_BISHOP] + bishopSqTable[square]
	}

	blackQueens := pos.getPiecesByTypeAndColor(PIECE_QUEEN, SIDE_BLACK)
	for blackQueens != 0 {
		square := bbPopLSB(&blackQueens)
		blackValue += pieceValues[PIECE_QUEEN] + queenSqTable[square]
	}

	blackKing := pos.getPiecesByTypeAndColor(PIECE_KING, SIDE_BLACK)
	for blackKing != 0 {
		square := bbPopLSB(&blackKing)
		blackValue += pieceValues[PIECE_KING] + blackKingSqTable[square]
	}

	result := whiteValue - blackValue

	if pos.turn == SIDE_WHITE {
		return result
	} else {
		return result * -1
	}
}
