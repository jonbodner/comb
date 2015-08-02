package comb

import (
	"github.com/jonbodner/sets"
	"sync"
)

func PermutationOrig(max int) [][]int {
	if max == 1 {
		return [][]int{{1}}
	}
	out := [][]int{}
	inner := PermutationOrig(max - 1)
	for _, v := range inner {
		for i := 0; i <= len(v); i++ {
			out = append(out, buildNext(v, i, max))
		}
	}
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

func CPermutation(max int) <-chan []int {
	return CPermutationPartial(max, max)
}

func CPermutationPartial(max int, pick int) <-chan []int {
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
		for i := 1; i <= max; i++ {
			wg.Add(1)
			go func(start int) {
				defer wg.Done()
				visit(start, pick, make([]int, pick))
			}(i)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func Combination(max int, num int) []sets.IntSet {
	v := make([]int, max)
	for k := 0; k < max; k++ {
		v[k] = k
	}
	return combInner(sets.IntSet{}, v, num)
}

func combInner(used sets.IntSet, vals []int, toKeep int) []sets.IntSet {
	if len(vals)+len(used) < toKeep {
		return []sets.IntSet{}
	}
	out := []sets.IntSet{}
	for i :=0;i<len(vals);i++ {
		start := used.Copy()
		start.Add(vals[i])
		if len(start) == toKeep {
			out = append(out, start)
		} else {
			out = append(out, combInner(start,vals[i+1:],toKeep)...)
		}
	}

	return out
}

func CombinationConc(max int, num int) []sets.IntSet {
	v := make([]int, max)
	for k := 0; k < max; k++ {
		v[k] = k
	}
	c := make(chan sets.IntSet)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		combInnerConc(sets.IntSet{}, v, num,c)
	}()
	go func() {
		defer close(c)
		wg.Wait()
	}()
	out := []sets.IntSet{}
	for s := range c {
		out = append(out, s)
	}
	return out
}

func combInnerConc(used sets.IntSet, vals []int, toKeep int, c chan sets.IntSet) {
	if len(vals)+len(used) < toKeep {
		return
	}
	var wg sync.WaitGroup
	for i :=0;i<len(vals);i++ {
		start := used.Copy()
		start.Add(vals[i])
		if len(start) == toKeep {
			c <- start
		} else {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				combInnerConc(start,vals[i+1:],toKeep, c)
			}(i)
		}
	}
	wg.Wait()
}

func CombinationOrig(max int, pick int) []sets.IntSet {
	val := sets.IntSet{}
	out := []sets.IntSet{}
	var visit func(int, int)
	visit = func(k int, depth int) {
		depth--
		val.Add(k)
		if depth == 0 {
			out = append(out, val.Copy())
		} else {
			for i := k + 1; i <= max; i++ {
				if !val.Contains(i) {
					visit(i, depth)
				}
			}
		}
		val.Remove(k)
	}

	for i := 1; i <= max; i++ {
		visit(i, pick)
	}
	return out
}
