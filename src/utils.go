package main

func getRankOf(square int) int {
	return square >> 3
}

func getFileOf(square int) int {
	return square & 7
}

func getColorOf(piece int) int {
	// White is 0, black is 8
	return piece >> 4
}

func getTypeOf(piece int) int {
	/*
		Pieces are from 1-6 with possible black color bit
		at the 4th bit, and this removes that to get the
		piece type
	*/
	return piece & 7
}

func getBBOfSquare(square int) uint64 {
	return squareBB[square]
}
