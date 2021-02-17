package main

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
