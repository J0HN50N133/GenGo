package slices

import (
	"fmt"
	"testing"
	"time"
)

func remove[A comparable](x A, l []A) []A {
	return Filter(func(y A) bool { return x != y }, l)
}

func permutations[A comparable](s []A) [][]A {
	if len(s) == 0 {
		return [][]A{[]A{}}
	}
	return ConcatMap(func(x A) [][]A {
		return Map(func(p []A) []A { return append(p, x) }, permutations(remove(x, s)))
	}, s)
}

func TestPermutations(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	p := permutations(a)
	if len(p) != 3628800 {
		t.Error("Permutations Implements Error")
	}
}

/*
isSafe p ps = not (elem p ps || sameDiag p ps)
    where
        sameDiag p ps = any (\(dist, q) -> abs (p - q) == dist) $ zip [1..] ps
*/
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func safe(k int, positions []int) bool {
	for i, p := range positions {
		//if p == k || p+i == k+len(positions) || p-i == k-len(positions) {
		if p == k || abs(p-k) == i+1 {
			return false
		}
	}
	return true
}

func sequence(a, b int) []int {
	if a > b {
		return []int{}
	}
	result := make([]int, b-a+1)
	for i := range result {
		result[i] = a + i
	}
	return result
}

func positions(k, n int) [][]int {
	if k == 0 {
		return [][]int{[]int{}}
	}
	conf := positions(k-1, n)
	return ConcatMap(func(p int) [][]int {
		//return Map(func(ps []int) []int { return append(ps, p) },
		return Map(func(ps []int) []int { return append([]int{p}, ps...) },
			Filter(func(ps []int) bool { return safe(p, ps) }, conf))
	}, sequence(1, n))
}

func Queen(boardSize int) [][]int {
	return positions(boardSize, boardSize)
}

func TestQueen(t *testing.T) {
	begin := time.Now()
	//ForEach(func(l []int) { fmt.Println(l) }, Queen(8))
	fmt.Println(len(Queen(8)))
	fmt.Println(time.Since(begin))
}
