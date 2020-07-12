package Week_01

func removeDuplicates(nums []int) int {
	if len(nums) <= 0{
		return 0
	}
	count := 1
	preNum := nums[0]
	for _,num := range nums{
		if preNum != num{
			preNum = num
			nums[count] = num
			count++
		}
	}
	return count
}