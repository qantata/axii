package main

import "fmt"

var divideDepth int = 0

func perft(pos Position, depth int, divide bool) uint64 {
	if depth == 0 {
		return 1
	}

	moves, nrMoves := generateMoves(pos)

	var nrNodes uint64 = 0
	for i := 0; i < nrMoves; i++ {
		move := moves[i]

		if !pos.isMoveLegal(move, true) {
			continue
		}

		nodes := perft(pos, depth-1, divide)
		pos.unmakeMove(move)

		if divide && depth == divideDepth {
			fmt.Printf("Move: %s\tNodes: %d\n", move.toStr(), nodes)
		}

		nrNodes += nodes
	}

	return nrNodes
}

func divide(pos Position, depth int) uint64 {
	divideDepth = depth
	return perft(pos, depth, true)
}
