package leetcode

func twoSum(nums []int, target int) []int {
	hTable := map[int]int{}
	for i, x := range nums {
	    if p, ok := hTable[target-x]; ok {
		return []int{p, i}
	    }
	    hTable[x] = i
	}
	return nil
}