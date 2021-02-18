package main

import (
	"fmt"
	"math"
)

var divideDepth int = 0
var pv = []Move{}
var pvScore = []int{}

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
	depth := 6
	pv = make([]Move, depth)
	pvScore = make([]int, depth)

	for i := 0; i < depth; i++ {
		pvScore[i] = -32000
	}

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
		return quiesce(pos, alpha, beta)
		//return pos.evaluate()
	}

	// 50 move rule
	if pos.state.rule50 >= 100 {
		return 0
	}

	bestScore := -32000
	mp := Movepick{}
	mp.generateMoves(*pos)

	nrLegalMoves := 0
	isInCheck := pos.isInCheck()

	for {
		move := mp.getNextMove()
		if move.move == 0 {
			break
		}

		if !pos.isMoveLegal(move, true) {
			continue
		}

		nrLegalMoves++

		score := -search(pos, depth-1, -beta, -alpha)
		pos.unmakeMove(move)

		if score > pvScore[depth-1] {
			pvScore[depth-1] = score
			pv[depth-1] = move
		}

		if score >= beta {
			return score
		}

		if score > bestScore {
			bestScore = score

			if score > alpha {
				alpha = score
			}
		}
	}

	if nrLegalMoves == 0 {
		// Checkmate
		if isInCheck {
			// Award checkmates that lead to it quicker
			return -30000 - int(math.Abs(float64(1-depth)))
		} else { // Stalemate
			return 0
		}
	}

	return bestScore
}

func quiesce(pos *Position, alpha int, beta int) int {
	eval := pos.evaluate()
	if eval >= beta {
		return beta
	}

	if eval > alpha {
		alpha = eval
	}

	mp := Movepick{}
	mp.generateMoves(*pos)

	for {
		move := mp.getNextMove()
		if move.move == 0 {
			break
		}

		if !pos.isMoveLegal(move, true) {
			continue
		}

		score := -quiesce(pos, -beta, -alpha)
		pos.unmakeMove(move)

		if score >= beta {
			return beta
		}

		if score > alpha {
			alpha = score
		}
	}

	return alpha
}
