package comb

import (
	"sync"
)

func PermutationOrig(max int) [][]int {
	if max == 1 {
		return [][]int{{1}}
	}
	out := [][]int{}
	inner := PermutationOrig(max - 1)
	for _, v := range inner {
		for i := 0;i<=len(v);i++ {
			out = append(out, buildNext(v, i, max))
		}
	}
	return out
}

func CPermutationOrig(max int) <-chan []int {
	out := make(chan []int)
	var innerFunc func(int)
	innerFunc = func(m int) {
		if m == 1 {
			out <- []int{1}
		}
		inner := innerFunc(m-1)
		for _, v := range inner {
			for i := 0;i<len(v);i++ {
				out = append(out, buildNext(v,i,m))
			}
		}
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		innerFunc(max)
	}()
	go func() {
		wg.Wait()
		close(out)
	}()
	return out	
}

func buildNext(v []int, i int, max int) []int {
	next := []int{}
	next = append(next, v[0:i]...)
	next = append(next, max)
	next = append(next, v[i:]...)
	return next
}

func contains(s []int, i int) bool {
	for _, v := range s {
		if v == i {
			return true
		}
	}
	return false
}

func Permutation(max int) [][]int {
	return PermutationPartial(max, max)
}

func PermutationPartial(max int, pick int) [][]int {
	val := make([]int, pick)
	out := [][]int{}
	var visit func(int, int)
	visit = func(k int, depth int) {
		depth--
		val[depth] = k
		if depth == 0 {
			next := make([]int, len(val))
			copy(next, val)
			out = append(out, next)
		} else {
			for i := 1; i <= max; i++ {
				if !contains(val, i) {
					visit(i, depth)
				}
			}
		}
		val[depth] = 0
	}

	for i := 1; i <= max; i++ {
		visit(i, pick)
	}
	return out
}

func PermutationPartialParallel(max int, pick int) <-chan []int {
	out := make(chan []int)
	var visit func(int, int, []int)
	visit = func(k int, depth int, val []int) {
		depth--
		val[depth] = k
		if depth == 0 {
			next := make([]int, len(val))
			copy(next, val)
			out <- next
		} else {
			for i := 1; i <= max; i++ {
				if !contains(val, i) {
					visit(i, depth, val)
				}
			}
		}
		val[depth] = 0
	}

	var wg sync.WaitGroup	
	go func() {
		for i := 1;i<=max;i++ {
			wg.Add(1)
			go func(start int) {
				defer wg.Done()
				visit(start, pick, make([]int,pick))
			}(i)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func Combination(max int) [][]int {
	out := [][]int{}
	return out
}

func CombinationPartial(max int, pick int) [][]int {
	out := [][]int{}
	return out
}
