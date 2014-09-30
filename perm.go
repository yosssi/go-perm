// Package perm generates and returns permutation slices.
package perm

// Int generates and returns permutation slices which have int elements.
func Int(src []int) [][]int {
	srcLen := len(src)

	pSlicesLen := factorial(srcLen)

	pSlices := make([][]int, pSlicesLen, pSlicesLen)

	pSlicesIdx := 0

	unusedIndices := make([]int, srcLen, srcLen)

	for i := 0; i < srcLen; i++ {
		unusedIndices[i] = i
	}

	genInt(src, srcLen, 0, unusedIndices, make([]int, srcLen, srcLen), pSlices, &pSlicesIdx)

	return pSlices
}

// genInt generates permutation slices which have int elements.
func genInt(src []int, srcLen int, recurCnt int, unusedIndices []int, work []int, pSlices [][]int, pSlicesIdx *int) {
	for idx, i := range unusedIndices {
		work[recurCnt] = src[i]

		if recurCnt == srcLen-1 {
			tmp := make([]int, srcLen, srcLen)
			copy(tmp, work)
			pSlices[*pSlicesIdx] = tmp
			*pSlicesIdx++
			return
		}

		l := len(unusedIndices)

		tmp := make([]int, l-1, l-1)

		tmpIdx := 0

		for _, n := range unusedIndices[:idx] {
			tmp[tmpIdx] = n
			tmpIdx++
		}

		for _, n := range unusedIndices[idx+1:] {
			tmp[tmpIdx] = n
			tmpIdx++
		}

		genInt(src, srcLen, recurCnt+1, tmp, work, pSlices, pSlicesIdx)
	}
}

// factorial returns a factorial of n.
func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}
