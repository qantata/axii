package main

func main() {
	initBitboards()
	initMagics()

	var p Position

	p.load("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	p.pretty()
}
