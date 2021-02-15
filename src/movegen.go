package main

func generateMoves(pos Position) [MAX_MOVES_IN_POS]Move {
	var moves [MAX_MOVES_IN_POS]Move
	var index int = 0

	generatePawnMoves(pos, &moves, &index)

	return moves
}

func generatePawnMoves(pos Position, moves *[MAX_MOVES_IN_POS]Move, index *int) {
	piece := pos.turn | PIECE_PAWN
	pieceCount := pos.getPieceCount(piece)

	for i := 0; i < pieceCount; i++ {
		origSquare := pos.pieceList[piece][i]
		rank := getRankOf(origSquare)
		file := getFileOf(origSquare)
		oppositeSide := pos.turn ^ 1
		isPromotion := (rank == RANK_7 && pos.turn == SIDE_WHITE) || (rank == RANK_2 && pos.turn == SIDE_BLACK)

		// Captures and EP
		pawnCaptureDir1 := DIR_NE
		pawnCaptureDir2 := DIR_NW

		if pos.turn == SIDE_BLACK {
			pawnCaptureDir1 = DIR_SE
			pawnCaptureDir2 = DIR_SW
		}

		pawnCapture1Square := origSquare + pawnCaptureDir1
		pawnCapture1Piece := pos.getPieceOn(pawnCapture1Square)
		if (pawnCapture1Square == pos.epSquare) || (file < FILE_H && pawnCapture1Piece != NO_PIECE && getColorOf(pawnCapture1Piece) == oppositeSide) {

			if isPromotion {
				for promotePiece := PROMOTION_PIECE_ROOK; promotePiece <= PROMOTION_PIECE_QUEEN; promotePiece++ {
					move := Move{}
					move.createWithPromotion(origSquare, pawnCapture1Square, promotePiece)
					moves[*index] = move
					*index++
				}
			} else {
				pawnCapture1Move := Move{}
				pawnCapture1Move.create(origSquare, pawnCapture1Square)
				moves[*index] = pawnCapture1Move
				*index++
			}
		}

		pawnCapture2Square := origSquare + pawnCaptureDir2
		pawnCapture2Piece := pos.getPieceOn(pawnCapture2Square)
		if (pawnCapture2Square == pos.epSquare) || (file > FILE_A && pawnCapture2Piece != NO_PIECE && getColorOf(pawnCapture2Piece) == oppositeSide) {
			if isPromotion {
				for promotePiece := PROMOTION_PIECE_ROOK; promotePiece <= PROMOTION_PIECE_QUEEN; promotePiece++ {
					move := Move{}
					move.createWithPromotion(origSquare, pawnCapture2Square, promotePiece)
					moves[*index] = move
					*index++
				}
			} else {
				pawnCapture2Move := Move{}
				pawnCapture2Move.create(origSquare, pawnCapture2Square)

				moves[*index] = pawnCapture2Move
				*index++
			}
		}

		// Pawn pushes
		pawnPushDir := DIR_N

		if pos.turn == SIDE_BLACK {
			pawnPushDir = DIR_S
		}

		// Single push
		singlePushSquare := origSquare + pawnPushDir
		if pos.getPieceOn(singlePushSquare) == NO_PIECE {
			if isPromotion {
				for promotePiece := PROMOTION_PIECE_ROOK; promotePiece <= PROMOTION_PIECE_QUEEN; promotePiece++ {
					move := Move{}
					move.createWithPromotion(origSquare, singlePushSquare, promotePiece)
					moves[*index] = move
					*index++
				}
			} else {
				singlePushMove := Move{}
				singlePushMove.create(origSquare, singlePushSquare)

				moves[*index] = singlePushMove
				*index++
			}

			// Double push
			if (rank == RANK_2 && pos.turn == SIDE_WHITE) || (rank == RANK_7 && pos.turn == SIDE_BLACK) {
				doublePushSquare := origSquare + (pawnPushDir * 2)
				if pos.getPieceOn(doublePushSquare) == NO_PIECE {
					doublePushMove := Move{}
					doublePushMove.create(origSquare, doublePushSquare)

					moves[*index] = doublePushMove
					*index++
				}
			}
		}
	}
}
