package main

import (
	"fmt"
	"math"
)

var DE_BRUJIN_TABLE_FORWARD = [64]int{
	SQ_A1, SQ_H6, SQ_B1, SQ_A8, SQ_A7, SQ_D4, SQ_C1, SQ_E8,
	SQ_B8, SQ_B7, SQ_B6, SQ_F5, SQ_E4, SQ_A3, SQ_D1, SQ_F8,
	SQ_G7, SQ_C8, SQ_D5, SQ_E7, SQ_C7, SQ_C6, SQ_F3, SQ_E6,
	SQ_G5, SQ_A5, SQ_F4, SQ_H3, SQ_B3, SQ_D2, SQ_E1, SQ_G8,
	SQ_G6, SQ_H7, SQ_C4, SQ_D8, SQ_A6, SQ_E5, SQ_H2, SQ_F7,
	SQ_C5, SQ_D7, SQ_E3, SQ_D6, SQ_H4, SQ_G3, SQ_C2, SQ_F6,
	SQ_B4, SQ_H5, SQ_G2, SQ_B5, SQ_D3, SQ_G4, SQ_B2, SQ_A4,
	SQ_F2, SQ_C3, SQ_A2, SQ_E2, SQ_H1, SQ_G1, SQ_F1, SQ_H8,
}

var squareBB [64]uint64

func initBitboards() {
	for square := SQ_A1; square <= SQ_H8; square++ {
		squareBB[square] = uint64(math.Pow(2, float64(square)))
	}

	fmt.Println("Initialized bitboards")
}

func bbCountBits(bb uint64) int {
	count := 0

	// Pop bits until bitboard is empty
	for bb > 0 {
		count++

		// Reset LS1B
		bb &= (bb - 1)
	}

	return count
}

/*
	Returns the index of the least significant bit in a bitboard (0-63)
	See https://www.chessprogramming.org/BitScan
*/
func bbGetLSBIndex(bb uint64) int {
	var debruijn64 uint64 = 0x03f79d71b4cb0a89
	return DE_BRUJIN_TABLE_FORWARD[((bb^(bb-1))*debruijn64)>>58]
}

func bbGetLS1BIndex(bb uint64) int {
	if bb != 0 {
		return bbCountBits((bb & -bb) - 1)
	}

	return -1
}

func bbPopLSB(bb *uint64) int {
	square := bbGetLSBIndex(*bb)
	*bb &= *bb - 1
	return square
}

func bbGetSquareBit(bb uint64, square int) uint64 {
	return bb & (1 << square)
}

func bbSetSquareBit(bb *uint64, square int) {
	*bb ^= (1 << square)
}

func bbPopSquareBit(bb *uint64, square int) {
	bit := bbGetSquareBit(*bb, square)

	if bit != 0 {
		*bb ^= (1 << square)
	}
}

func bbPrint(bb uint64) {
	var result string = ""
	square := SQ_A8

	for rank := RANK_8; rank >= RANK_1; rank-- {
		for file := FILE_A; file <= FILE_H; file++ {
			if (bb & squareBB[square]) != 0 {
				result += "x "
			} else {
				result += ". "
			}

			square++
		}

		result += "\n"
		square -= 16
	}

	fmt.Printf("%s\n", result)
}
