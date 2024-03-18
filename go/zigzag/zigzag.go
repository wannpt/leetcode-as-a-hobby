package zigzag

import "strings"

func Convert(s string, numRows int) string {
	tmp := make([]string, numRows)
	rowPos := 0
	down := true
	for i := 0; i < len(s); i++ {
		tmp[rowPos] += string(s[i])

		if down {
			rowPos++
		} else {
			rowPos--
		}

		if rowPos == 0 || rowPos == numRows-1 {
			down = !down
		}

	}

	return strings.Join(tmp, "")
}
