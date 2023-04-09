package main

import "fmt"

//import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums1 int整型一维数组
 * @param m int整型
 * @param nums2 int整型一维数组
 * @param n int整型
 * @return int整型一维数组
 */
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	// write code here
	nums := make([]int, 0, len(nums1))
	var mm, nn = 0, 0
	for {
		if mm < m && nn < n {
			if nums1[mm] < nums2[nn] {
				nums = append(nums, nums1[mm])
				mm++
			} else {
				nums = append(nums, nums2[nn])
				nn++
			}
		} else {
			for mm < m {
				nums = append(nums, nums1[mm])
				mm++
			}
			for nn < n {
				nums = append(nums, nums2[nn])
				nn++
			}
			break
		}
	}
	return nums
}

func merge1(nums1 []int, m int, nums2 []int, n int) []int {
	var index = len(nums1) - 1
	m--
	n--
	for {
		if m >= 0 && n >= 0 {
			if nums1[m] >= nums2[n] {
				nums1[index] = nums1[m]
				m--
			} else {
				nums1[index] = nums2[n]
				n--
			}
			index--
		} else {
			for m >= 0 {
				nums1[index] = nums1[m]
				m--
			}
			for n >= 0 {
				nums1[index] = nums2[n]
				n--
			}
			break
		}
	}
	return nums1
}

func main() {
	s1 := []int{1, 2, 3, 0, 0, 0}
	s2 := []int{2, 4, 5}
	s3 := merge(s1, 3, s2, 3)
	fmt.Println(s3)

	s4 := merge1(s1, 3, s2, 3)
	fmt.Println(s4)
}
