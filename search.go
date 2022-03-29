package main

import (
	"fmt"
	"math"
	"time"
)

var divideDepth int = 0
var pv = []Move{}
var pvScore = []int{}
var pvMove Move = Move{}

func perft(pos Position, depth int, divide bool) uint64 {
	if depth == 0 {
		return 1
	}

	moves, nrMoves := generateMoves(pos, false)

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
	pv = make([]Move, 100)
	pvScore = make([]int, 100)
	pvMove = Move{}

	for i := 0; i < 100; i++ {
		pvScore[i] = -32000
	}

	maxThinkTime := 5.0
	if pos.fullMoveCounter >= 50 {
		maxThinkTime = 90.0
	} else if pos.fullMoveCounter >= 40 {
		maxThinkTime = 75.0
	} else if pos.fullMoveCounter >= 30 {
		maxThinkTime = 60.0
	} else if pos.fullMoveCounter >= 20 {
		maxThinkTime = 45.0
	} else if pos.fullMoveCounter >= 10 {
		maxThinkTime = 30.0
	} else if pos.fullMoveCounter >= 5 {
		maxThinkTime = 15.0
	}

	start := time.Now()
	latestDepth := 1
	for depth := 1; depth < 100; depth++ {
		latestDepth = depth

		before := time.Now()
		search(&pos, depth, -32000, 32000)

		fullDuration := time.Since(start).Seconds()
		thisDuration := time.Since(before).Seconds()

		if fullDuration > maxThinkTime || fullDuration+(thisDuration*4) > maxThinkTime {
			break
		}
	}

	/*
		fmt.Printf("info depth %d pv ", depth)
		for i := depth - 1; i >= 0; i-- {
			fmt.Printf("%s ", pv[i].toStr())
		}*/

	fmt.Println("Used depth", latestDepth)
	pvMove = pv[latestDepth-1]

	if !TESTING {
		fmt.Printf("bestmove %s\n", pvMove.toStr())
	}

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
	mp.generateMoves(*pos, false)

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
	mp.generateMoves(*pos, true)

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
