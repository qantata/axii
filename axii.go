package main

var TESTING = false

func main() {
	initPosition()
	initBitboards()
	initMagics()
	initMoveGen()

	var pos Position
	pos.load(STARTING_POS)

	uciLoop(&pos)
}
