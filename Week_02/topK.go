package Week_02

import "sort"

/*
	剑指 Offer 40. 最小的k个数
	输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
	示例 1：
	输入：arr = [3,2,1], k = 2
	输出：[1,2] 或者 [2,1]
	示例 2：
	输入：arr = [0,1,2,1], k = 1
	输出：[0]

	限制：
	0 <= k <= arr.length <= 10000
	0 <= arr[i] <= 10000
 */

/**
	先排序，然后取前k个
 */
func getLeastNumbersSort(arr []int, k int) []int {
	sort.Slice(arr,func(i,j int)bool{
		return arr[i] < arr[j]
	})
	return arr[:k]
}

func getLeastNumbers(arr []int,k int) []int{
	lo := 0
	hi := len(arr)

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(arr, i, hi)
	}

	// Pop elements, largest first, into end of data.
	for i := hi - 1; i >= 0; i-- {
		arr[0],arr[i] = arr[i],arr[0]
		siftDown(arr, lo, i)
	}
	return arr[:k]
}

func siftDown(arr []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && arr[child] < arr[child+1] {
			child++
		}
		if arr[root] >= arr[child] {
			return
		}
		arr[root], arr[child] = arr[child],arr[root]
		root = child
	}
}