package leetcode

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for key, value := range nums {
		fmt.Println("key", key)
		fmt.Println("value", value)
		if idx, ok := m[target-value]; ok {
			return []int{idx, key}
		}
		m[value] = key
	}

	return nil
}
