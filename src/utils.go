package main

func getRankOf(square int) int {
	return square >> 3
}

func getFileOf(square int) int {
	return square & 7
}

func getColorOf(piece int) int {
	return piece >> 4
}
