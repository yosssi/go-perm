package perm

import (
	"fmt"
	"testing"
)

var intsMap = make(map[int][]int)

func init() {
	addIntsMap(0)
	addIntsMap(1)
	addIntsMap(2)
	addIntsMap(3)
	addIntsMap(4)
	addIntsMap(5)
	addIntsMap(10)
	addIntsMap(11)
}

func TestInt_nil(t *testing.T) {
	fmt.Println(Int(nil))
}

func TestInt_len0(t *testing.T) {
	fmt.Println(Int(intsMap[0]))
}

func TestInt_len1(t *testing.T) {
	fmt.Println(Int(intsMap[1]))
}

func TestInt_len2(t *testing.T) {
	fmt.Println(Int(intsMap[2]))
}

func TestInt_len3(t *testing.T) {
	fmt.Println(Int(intsMap[3]))
}

func TestInt_len4(t *testing.T) {
	fmt.Println(Int(intsMap[4]))
}

func TestInt_len5(t *testing.T) {
	// Refrain from printing the result because it has a lot of values.
	Int(intsMap[5])
}

func TestInt_len10(t *testing.T) {
	Int(intsMap[10])
}

func TestInt_len11(t *testing.T) {
	Int(intsMap[11])
}

func genInts(l int) []int {
	ints := make([]int, l, l)

	for i := 0; i < l; i++ {
		ints[i] = i + 1
	}

	return ints
}

func addIntsMap(i int) {
	intsMap[i] = genInts(i)
}
