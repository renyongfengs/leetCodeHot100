package heap

//215. 数组中的第K个最大元素
//给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//链接：https://leetcode.cn/problems/kth-largest-element-in-an-array/description/?envType=study-plan-v2&envId=top-100-liked

func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 || k > len(nums) {
		return 0
	}
	return findKthLargestQuickSort(nums, 0, len(nums)-1, len(nums)-k)
}

func findKthLargestQuickSort(arr []int, left, right, k int) int {
	if left < right {
		result := quickPartition(arr, left, right)
		if k > result[0] && k < result[1] {
			return arr[k]
		} else if k <= result[0] {
			return findKthLargestQuickSort(arr, left, result[0], k)
		} else if k >= result[1] {
			return findKthLargestQuickSort(arr, result[1], right, k)
		}

	}
	return arr[left]
}

func quickPartition(arr []int, left, right int) []int {
	mid := left + (right-left)>>1
	target := arr[mid]
	lBorder, rBorder := left-1, right+1
	for i := left; i < rBorder; {
		if arr[i] < target {
			lBorder++
			swap(arr, i, lBorder)
			i++
		} else if arr[i] > target {
			rBorder--
			swap(arr, i, rBorder)
		} else {
			i++
		}
	}

	return []int{lBorder, rBorder}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
