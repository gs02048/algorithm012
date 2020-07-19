package Week_02

import "sort"

/*
	给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

	示例:

	输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
	输出:
	[
	  ["ate","eat","tea"],
	  ["nat","tan"],
	  ["bat"]
	]
	说明：

	所有输入均为小写字母。
	不考虑答案输出的顺序。
 */
func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string]int)
	for _,str := range strs{
		byteS := []byte(str)
		sort.Slice(byteS, func(i, j int) bool {
			return byteS[i] < byteS[j]
		})
		newStr := string(byteS)
		if idx,ok := m[newStr];!ok{
			m[newStr] = len(res)
			res = append(res,[]string{str})
		}else{
			res[idx] = append(res[idx],str)
		}
	}
	return res
}