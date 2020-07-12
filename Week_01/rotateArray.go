package Week_01

func reverse(nums []int) {  // 翻转数组
	i,j := 0, len(nums)-1
	for i<j {
		nums[i],nums[j] = nums[j],nums[i]
		i++
		j--
	}
}

func Rotate(nums []int, k int)  {
	k = k%len(nums)
	reverse(nums[:len(nums)-k])
	reverse(nums[len(nums)-k:])
	reverse(nums)
}