package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	parameter
	answer
}

type parameter struct {
	nums   []int
	target int
}

type answer struct {
	ans []int
}

func Test_Question1(t *testing.T) {
	ques := []question1{
		{
			parameter{[]int{2,7,11,15}, 9},
			answer{[]int{0, 1}},
		},

		{
			parameter{[]int{3, 2, 4}, 6},
			answer{[]int{1, 2}},
		},

		{
			parameter{[]int{3,3}, 6},
			answer{[]int{0, 1}},
		},
	}

	fmt.Printf("--------Leetcode Question No.1 Test--------\n")

	for _, que := range ques {
		par :=  que.parameter
		fmt.Printf("Input:%v   Output:%v\n", par, twoSum(par.nums, par.target))
	}
}