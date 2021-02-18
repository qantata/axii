package main

import (
	"fmt"
)

var divideDepth int = 0
var pv = []Move{}

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

func searchGo(pos Position) {
	depth := 5
	pv = make([]Move, depth)

	search(&pos, depth, -32000, 32000)

	/*
		fmt.Printf("info depth %d pv ", depth)
		for i := depth - 1; i >= 0; i-- {
			fmt.Printf("%s ", pv[i].toStr())
		}*/

	fmt.Printf("bestmove %s\n", pv[depth-1].toStr())
}

func search(pos *Position, depth int, alpha int, beta int) int {
	if depth == 0 {
		return pos.evaluate()
	}

	bestScore := -32000
	moves, nrMoves := generateMoves(*pos)

	for i := 0; i < nrMoves; i++ {
		move := moves[i]
		if !pos.isMoveLegal(move, true) {
			continue
		}

		score := -search(pos, depth-1, -beta, -alpha)
		pos.unmakeMove(move)

		if score >= beta {
			return score
		}

		if score > bestScore {
			bestScore = score

			if score > alpha {
				alpha = score
				pv[depth-1] = move
			}
		}
	}

	return bestScore
}
