package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	cbFile, cbFileExists := cb[file]

	if !cbFileExists {
		return 0
	}

	countInFile := 0

	for _, value := range cbFile {
		if value {
			countInFile++
		}
	}

	return countInFile
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	rank--
	if rank < 0 || rank > 7 {
		return 0
	}

	countInRank := 0

	for _, v := range cb {
		if v[rank] {
			countInRank++
		}
	}

	return countInRank
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	squareCount := 0
	for _, v := range cb {
		squareCount += len(v)
	}
	return squareCount
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSquareCount := 0
	for i := range cb {
		occupiedSquareCount += CountInFile(cb, i)
	}
	return occupiedSquareCount
}
