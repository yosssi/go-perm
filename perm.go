// Package perm generates and returns permutation slices.
package perm

// Int generates and returns permutation slices which have int elements.
func Int(src []int) [][]int {
	srcLen := len(src)

	pSlicesLen := factorial(srcLen)

	pSlices := make([][]int, pSlicesLen, pSlicesLen)

	pSlicesIdx := 0

	genInt(src, srcLen, 0, make(map[int]struct{}), make([]int, srcLen, srcLen), pSlices, &pSlicesIdx)

	return pSlices
}

// genInt generates permutation slices which have int elements.
func genInt(src []int, srcLen int, recurCnt int, usedIndices map[int]struct{}, work []int, pSlices [][]int, pSlicesIdx *int) {
	for i, v := range src {
		if _, ok := usedIndices[i]; ok {
			continue
		}

		work[recurCnt] = v

		if recurCnt == srcLen-1 {
			tmp := make([]int, srcLen, srcLen)
			copy(tmp, work)
			pSlices[*pSlicesIdx] = tmp
			*pSlicesIdx++
			return
		}

		usedIndices[i] = struct{}{}

		genInt(src, srcLen, recurCnt+1, usedIndices, work, pSlices, pSlicesIdx)

		delete(usedIndices, i)
	}
}

// factorial returns a factorial of n.
func factorial(n int) int {
	if n <= 0 {
		return 1
	}
	return n * factorial(n-1)
}
