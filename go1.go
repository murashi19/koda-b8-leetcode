package main

func twoSum(nums []int, target int) []int {
	var result = make(map[int]int)
	for i, num := range nums {
		x := target - num
		if index, found := result[x]; found {
			return []int{index, i}
		}
		result[num] = i
	}
	return nil
}
