package main

type Movepick struct {
	moves     [MAX_MOVES_IN_POS]Move
	nrMoves   int
	currIndex int
}

func (mp *Movepick) generateMoves(pos Position) {
	moves, nrMoves := generateMoves(pos)
	mp.moves = moves
	mp.nrMoves = nrMoves
	mp.currIndex = 0
}

func (mp *Movepick) getNextMove() Move {
	if mp.currIndex >= mp.nrMoves {
		return Move{}
	}

	bestScore := -32000
	bestIndex := mp.currIndex

	for i := mp.currIndex; i < mp.nrMoves; i++ {
		move := mp.moves[i]

		if move.score > bestScore {
			bestScore = move.score
			bestIndex = i
		}
	}

	if bestIndex != mp.currIndex {
		temp := mp.moves[mp.currIndex]
		mp.moves[mp.currIndex] = mp.moves[bestIndex]
		mp.moves[bestIndex] = temp
	}

	move := mp.moves[mp.currIndex]
	mp.currIndex++
	return move
}
