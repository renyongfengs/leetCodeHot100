package binarySearch

//4. 寻找两个正序数组的中位数
//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//算法的时间复杂度应该为 O(log (m+n)) 。
//链接：https://leetcode.cn/problems/median-of-two-sorted-arrays/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/median-of-two-sorted-arrays/solutions/3026487/tong-su-yi-dong-de-fang-fa-jiang-jie-gol-6kw5/?envType=study-plan-v2&envId=top-100-liked
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	if total%2 == 0 {
		return float64(findN(nums1, nums2, total/2)+findN(nums1, nums2, total/2+1)) / 2
	} else {
		return float64(findN(nums1, nums2, total/2+1))
	}
}

func findN(nums1 []int, nums2 []int, n int) int {
	if len(nums1) == 0 {
		return nums2[n-1]
	}

	if len(nums2) == 0 {
		return nums1[n-1]
	}

	if n == 1 {
		return Min(nums1[0], nums2[0])
	}

	index := n / 2
	i1, i2 := Min(index, len(nums1)), Min(index, len(nums2))

	if nums1[i1-1] < nums2[i2-1] {
		return findN(nums1[i1:], nums2, n-i1)
	} else {
		return findN(nums1, nums2[i2:], n-i2)
	}
}

func Min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
