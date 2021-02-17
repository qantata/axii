package main

import "fmt"

/*
	Credit for the algorithms to maksimKorzh on github and "Chess Programming" on YouTube
	Youtube video: https://www.youtube.com/watch?v=4ohJQ9pCkHI
	Reference link: https://github.com/maksimKorzh/chess_programming/blob/master/src/magics/magics.c
*/

// Masks
var bishopMasks [64]uint64
var rookMasks [64]uint64

// Attacks
var bishopAttacks [64]map[uint64]uint64
var rookAttacks [64]map[uint64]uint64

var bishopRelevantBits = []uint64{
	6, 5, 5, 5, 5, 5, 5, 6,
	5, 5, 5, 5, 5, 5, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 9, 9, 7, 5, 5,
	5, 5, 7, 7, 7, 7, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5,
	6, 5, 5, 5, 5, 5, 5, 6,
}

var rookRelevantBits = []uint64{
	12, 11, 11, 11, 11, 11, 11, 12,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	11, 10, 10, 10, 10, 10, 10, 11,
	12, 11, 11, 11, 11, 11, 11, 12,
}

func maskBishopAttacks(square int) uint64 {
	// Attack bitboard
	var attacks uint64 = 0

	var file int
	var rank int

	// Init target files & ranks
	var tr = square / 8
	var tf = square % 8

	// Generate attacks
	for rank, file = tr+1, tf+1; rank <= 6 && file <= 6; rank, file = rank+1, file+1 {
		attacks |= (uint64(1) << (rank*8 + file))
	}

	for rank, file = tr+1, tf-1; rank <= 6 && file >= 1; rank, file = rank+1, file-1 {
		attacks |= (uint64(1) << (rank*8 + file))
	}

	for rank, file = tr-1, tf+1; rank >= 1 && file <= 6; rank, file = rank-1, file+1 {
		attacks |= (uint64(1) << (rank*8 + file))
	}

	for rank, file = tr-1, tf-1; rank >= 1 && file >= 1; rank, file = rank-1, file-1 {
		attacks |= (uint64(1) << (rank*8 + file))
	}

	return attacks
}

func maskRookAttacks(square int) uint64 {
	// Attack bitboard
	var attacks uint64 = 0

	var file int
	var rank int

	// Init target files & ranks
	var tr = square / 8
	var tf = square % 8

	// Generate attacks
	for rank = tr + 1; rank <= 6; rank++ {
		attacks |= (1 << (rank*8 + tf))
	}

	for rank = tr - 1; rank >= 1; rank-- {
		attacks |= (1 << (rank*8 + tf))
	}

	for file = tf + 1; file <= 6; file++ {
		attacks |= (1 << (tr*8 + file))
	}

	for file = tf - 1; file >= 1; file-- {
		attacks |= (1 << (tr*8 + file))
	}

	return attacks
}

func setOccupancy(index int, nrBitsInMask int, attackMask uint64) uint64 {
	var occupancy uint64 = 0

	for count := 0; count < nrBitsInMask; count++ {
		square := bbGetLS1BIndex(attackMask)

		bbPopSquareBit(&attackMask, square)

		// Make sure occupancy is on the board
		if (index & (1 << count)) != 0 {
			occupancy |= (1 << square)
		}
	}

	return occupancy
}

func bishopAttacksOnTheFly(square int, block uint64) uint64 {
	// Attack bitboard
	var attacks uint64 = 0

	var file int
	var rank int

	// Init target files & ranks
	var tr = square / 8
	var tf = square % 8

	for rank, file = tr+1, tf+1; rank <= 7 && file <= 7; rank, file = rank+1, file+1 {
		attacks |= (1 << (rank*8 + file))
		if (block & (1 << (rank*8 + file))) != 0 {
			break
		}
	}

	for rank, file = tr+1, tf-1; rank <= 7 && file >= 0; rank, file = rank+1, file-1 {
		attacks |= (1 << (rank*8 + file))
		if (block & (1 << (rank*8 + file))) != 0 {
			break
		}
	}

	for rank, file = tr-1, tf+1; rank >= 0 && file <= 7; rank, file = rank-1, file+1 {
		attacks |= (1 << (rank*8 + file))
		if (block & (1 << (rank*8 + file))) != 0 {
			break
		}
	}

	for rank, file = tr-1, tf-1; rank >= 0 && file >= 0; rank, file = rank-1, file-1 {
		attacks |= (1 << (rank*8 + file))
		if (block & (1 << (rank*8 + file))) != 0 {
			break
		}
	}

	return attacks
}

func rookAttacksOnTheFly(square int, block uint64) uint64 {
	// Attack bitboard
	var attacks uint64 = 0

	var file int
	var rank int

	// Init target files & ranks
	var tr = square / 8
	var tf = square % 8

	// Generate attacks
	for rank = tr + 1; rank <= 7; rank++ {
		attacks |= (1 << (rank*8 + tf))
		if (block & (1 << (rank*8 + tf))) != 0 {
			break
		}
	}

	for rank = tr - 1; rank >= 0; rank-- {
		attacks |= (1 << (rank*8 + tf))
		if (block & (1 << (rank*8 + tf))) != 0 {
			break
		}
	}

	for file = tf + 1; file <= 7; file++ {
		attacks |= (1 << (tr*8 + file))
		if (block & (1 << (tr*8 + file))) != 0 {
			break
		}
	}

	for file = tf - 1; file >= 0; file-- {
		attacks |= (1 << (tr*8 + file))
		if (block & (1 << (tr*8 + file))) != 0 {
			break
		}
	}

	return attacks
}

func initSliderAttacks(isBishop bool) {
	for square := SQ_A1; square <= SQ_H8; square++ {
		// Init bishop & rook masks
		bishopMasks[square] = maskBishopAttacks(square)
		rookMasks[square] = maskRookAttacks(square)

		// Current mask
		var mask uint64

		if isBishop {
			mask = maskBishopAttacks(square)
		} else {
			mask = maskRookAttacks(square)
		}

		// Count attack mask bits
		bitCount := bbCountBits(mask)

		// Occupancy variations count
		occupancyVariationsCount := (1 << bitCount)

		for count := 0; count < occupancyVariationsCount; count++ {
			occupancy := setOccupancy(count, bitCount, mask)

			// Init occupancies, magic index & attacks
			if isBishop {
				magicIndex := occupancy * bishopMagics[square] >> (64 - bishopRelevantBits[square])
				bishopAttacks[square][magicIndex] = bishopAttacksOnTheFly(square, occupancy)
			} else {
				magicIndex := occupancy * rookMagics[square] >> (64 - rookRelevantBits[square])
				rookAttacks[square][magicIndex] = rookAttacksOnTheFly(square, occupancy)
			}
		}
	}
}

func getBishopAttacks(square int, occupancy uint64) uint64 {
	occupancy &= bishopMasks[square]
	occupancy *= bishopMagics[square]
	occupancy >>= 64 - bishopRelevantBits[square]

	return bishopAttacks[square][occupancy]
}

func getRookAttacks(square int, occupancy uint64) uint64 {
	occupancy &= rookMasks[square]
	occupancy *= rookMagics[square]
	occupancy >>= 64 - rookRelevantBits[square]

	return rookAttacks[square][occupancy]
}

func initMagics() {
	for square := SQ_A1; square <= SQ_H8; square++ {
		bishopAttacks[square] = make(map[uint64]uint64)
		rookAttacks[square] = make(map[uint64]uint64)
	}

	initSliderAttacks(true)
	initSliderAttacks(false)

	fmt.Println("Initialized magics")
}
