package Week_03

// 最终结果
var result [][]int

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used[]bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(nums))
		copy(tmp, pathNums)
		result = append(result, tmp)
		return
	}

	for i:=0; i<len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backtrack(nums,pathNums,used)
			pathNums = pathNums[:len(pathNums) -1]
			used[i] = false
		}
	}
}

func permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	result = [][]int{}
	backtrack(nums, pathNums, used)
	return result
}