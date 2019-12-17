package main

//给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。
//
//请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
//
//你可以假设 nums1 和 nums2 不会同时为空。
//
//示例 1:
//
//nums1 = [1, 3]
//nums2 = [2]
//
//则中位数是 2.0
//示例 2:
//
//nums1 = [1, 2]
//nums2 = [3, 4]
//
//则中位数是 (2 + 3)/2 = 2.5
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

const MaxInt = int(^uint(0) >> 1)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var l1, l2, i1, i2, v1, v2, tmp, mid1, mid2, lc int
	var res float64
	l1 = len(nums1)
	l2 = len(nums2)
	mid2 = (l2 + l1) / 2
	if (l1+l2)%2 == 0 {
		mid1 = mid2 - 1
	} else {
		mid1 = mid2
	}
	for i1 <= l1 || i2 <= l2 {
		if i1 < l1 {
			v1 = nums1[i1]
		} else {
			v1 = MaxInt
		}
		if i2 < l2 {
			v2 = nums2[i2]
		} else {
			v2 = MaxInt
		}
		if v1 < v2 {
			tmp = v1
			i1++
		} else {
			tmp = v2
			i2++
		}
		if lc == mid1 {
			res = float64(tmp)
		}
		if lc == mid2 {
			res = (res + float64(tmp)) / 2
			break
		}
		lc++
	}
	return res
}
