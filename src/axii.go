package main

func main() {
	initBitboards()
	initMagics()

	var pos Position

	pos.load(STARTING_POS)
	uciLoop(&pos)
}
