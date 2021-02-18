package main

func main() {
	initBitboards()
	initMagics()
	initMoveGen()

	var pos Position

	pos.load(STARTING_POS)
	uciLoop(&pos)

}
