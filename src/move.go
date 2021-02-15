package main

type Move struct {
	/*
		The move is stored in the first 16 bits of the integer
		Bits 0-5 store the origin square
		Bits 6-11 store the destination square
		Bits 12-13 store promotion piece type (rook, knight, bishop, queen)
		Bits 14-15 store move flag (normal, promotion, ep, castling)
	*/
	move int
}

// Get origin square of the move
func (move Move) orig() int {
	return move.move & 63
}

// Get destination square of the move
func (move Move) dest() int {
	return (move.move >> 6) & 63
}

func (move *Move) create(orig int, dest int) {
	move.move = (dest << 6) + orig
}

func (move *Move) createWithPromotion(orig int, dest int, promotionPiece int) {
	// 1 = promotion
	move.move = (1 << 14) + (promotionPiece << 12) + (dest << 6) + orig
}
