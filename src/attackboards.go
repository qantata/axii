package main

func getPawnAttacks(square int, color int) uint64 {
	return pawnAttacksBB[square][color]
}

func getRookAttacks(square int, occupancy uint64) uint64 {
	occupancy &= rookMasks[square]
	occupancy *= rookMagics[square]
	occupancy >>= 64 - rookRelevantBits[square]

	return rookAttacks[square][occupancy]
}

func getKnightAttacks(square int) uint64 {
	return knightAttacksBB[square]
}

func getBishopAttacks(square int, occupancy uint64) uint64 {
	occupancy &= bishopMasks[square]
	occupancy *= bishopMagics[square]
	occupancy >>= 64 - bishopRelevantBits[square]

	return bishopAttacks[square][occupancy]
}

func getQueenAttacks(square int, occupancy uint64) uint64 {
	return getBishopAttacks(square, occupancy) | getRookAttacks(square, occupancy)
}

func getKingAttacks(square int) uint64 {
	return kingAttacksBB[square]
}

func getAttackersBBToSquare(square int, occupancy uint64, pos Position) uint64 {
	whitePawns := getPawnAttacks(square, SIDE_BLACK) & pos.getPiecesByTypeAndColor(PIECE_PAWN, SIDE_WHITE)
	blackPawns := getPawnAttacks(square, SIDE_WHITE) & pos.getPiecesByTypeAndColor(PIECE_PAWN, SIDE_BLACK)
	rooksQueens := getRookAttacks(square, occupancy) & (pos.getPiecesByType(PIECE_ROOK) | pos.getPiecesByType(PIECE_QUEEN))
	knights := getKnightAttacks(square) & pos.getPiecesByType(PIECE_KNIGHT)
	bishopsQueens := getBishopAttacks(square, occupancy) & (pos.getPiecesByType(PIECE_BISHOP) | pos.getPiecesByType(PIECE_QUEEN))
	kings := getKingAttacks(square) & pos.getPiecesByType(PIECE_KING)

	return whitePawns | blackPawns | rooksQueens | knights | bishopsQueens | kings
}
