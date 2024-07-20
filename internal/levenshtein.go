package internal

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func levenshteinDistance(inputStr, compareStr string) int {
	m := len(inputStr)
	n := len(compareStr)

	prevRow := make([]int, n+1)
	currRow := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		prevRow[i] = i
	}

	for j := 1; j < m+1; j++ {
		currRow[0] = j

		for k := 1; k < n+1; k++ {
			if inputStr[j-1] == compareStr[k-1] {
				currRow[k] = prevRow[k-1]
			} else {
				currRow[k] = 1 + min(currRow[k-1], min(prevRow[k], prevRow[k-1]))
			}
		}
		copy(prevRow, currRow)
	}
	return currRow[n]
}
