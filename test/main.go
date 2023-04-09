package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	result = make([]int, 2)
	var indexMap map[int]int
	indexMap = make(map[int]int, len(nums))

	for i := range nums {
		index, exist := indexMap[target-nums[i]]
		if exist {
			result[0] = index
			result[1] = i
			return result
		} else {
			indexMap[nums[i]] = i
		}
	}

	return result
}

func twoSum0(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 9
	var result = twoSum(nums, target)
	for i := 0; i < 2; i++ {
		fmt.Printf("%d ", result[i])
	}

	fmt.Println()

	var result0 = twoSum0(nums, target)
	for i := 0; i < 2; i++ {
		fmt.Printf("%d ", result0[i])
	}
}
