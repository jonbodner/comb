package comb

import (
	"fmt"
	"testing"
)

func validateResults(t *testing.T, k interface{}, v [][]int, result [][]int) {
	if len(result) != len(v) {
		t.Errorf("wrong length of result: %v, %v, %v", k, v, result)
	}
	for _, rv := range result {
		found := false
		for _, rt := range v {
			found = true
			for pos, val := range rv {
				if val != rt[pos] {
					found = false
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			t.Errorf("cannot find an expected entry: %d, %v, %v", k, v, result)
		}
	}
}

var permResults = map[int][][]int{
	1: [][]int{{1}},
	2: [][]int{{2, 1}, {1, 2}},
	3: [][]int{{1, 2, 3}, {1, 3, 2}, {3, 1, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}},
}

func TestPermutation(t *testing.T) {
	for k, v := range permResults {
		result := Permutation(k)
		validateResults(t, k, v, result)
	}
}

var permPartialResults = map[[2]int][][]int{
	[2]int{1, 1}: [][]int{{1}},
	[2]int{2, 2}: [][]int{{2, 1}, {1, 2}},
	[2]int{2, 1}: [][]int{{1}, {2}},
	[2]int{3, 1}: [][]int{{1}, {2}, {3}},
	[2]int{3, 2}: [][]int{{1, 2}, {1, 3}, {3, 1}, {2, 1}, {2, 3}, {3, 2}},
	[2]int{4, 1}: [][]int{{1}, {2}, {3}, {4}},
	[2]int{4, 2}: [][]int{{1, 2}, {2, 1}, {1, 3}, {3, 1}, {1, 4}, {4, 1}, {2, 3}, {3, 2}, {2, 4}, {4, 2}, {3, 4}, {4, 3}},
	[2]int{4, 3}: [][]int{
		{1, 2, 3}, {1, 3, 2}, {3, 1, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1},
		{4, 2, 3}, {4, 3, 2}, {3, 4, 2}, {2, 4, 3}, {2, 3, 4}, {3, 2, 4},
		{1, 4, 3}, {1, 3, 4}, {3, 1, 4}, {4, 1, 3}, {4, 3, 1}, {3, 4, 1},
		{1, 2, 4}, {1, 4, 2}, {4, 1, 2}, {2, 1, 4}, {2, 4, 1}, {4, 2, 1},
	},
}

func TestPermutationPartial(t *testing.T) {
	for k, v := range permPartialResults {
		result := PermutationPartial(k[0], k[1])
		validateResults(t, k, v, result)
	}
}

func TestIncomplete(t *testing.T) {
	out := CPermutationPartial(3, 3)
	for p := range out {
		fmt.Println(p)
	}

	out = CPermutationPartial(3, 2)
	for p := range out {
		fmt.Println(p)
	}

	out = CPermutationPartial(3, 1)
	for p := range out {
		fmt.Println(p)
	}

	out2 := PermutationPartial(3, 3)
	fmt.Println(out2)

	out2 = PermutationPartial(3, 2)
	fmt.Println(out2)

	out2 = PermutationPartial(3, 1)
	fmt.Println(out2)

	out3 := Combination(3, 3)
	fmt.Println(out3)

	out3 = Combination(3, 2)
	fmt.Println(out3)

	out3 = Combination(3, 1)
	fmt.Println(out3)

	out3 = Combination(4, 4)
	fmt.Println(out3)

	out3 = Combination(4, 3)
	fmt.Println(out3)

	out3 = Combination(4, 2)
	fmt.Println(out3)

	out3 = Combination(4, 1)
	fmt.Println(out3)
}
