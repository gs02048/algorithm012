package Week_02

import (
	"sort"
)
/*
	给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

	示例 1:

	输入: s = "anagram", t = "nagaram"
	输出: true
	示例 2:

	输入: s = "rat", t = "car"
	输出: false
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}
	count := make(map[uint8]int,len(s))
	for i := 0;i<len(s);i++{
		count[s[i]]++
		count[t[i]]--
	}
	for _,v := range count{
		if v != 0{
			return false
		}
	}
	return true
}

func isAnagramSort(s string,t string) bool{
	if len(s) != len(t){
		return false
	}
	byteS := []byte(s)
	byteT := []byte(t)

	sort.Slice(byteS,func(i,j int) bool{
		return byteS[i] < byteS[j]
	})

	sort.Slice(byteT,func(i,j int)bool{
		return byteT[i] < byteT[j]
	})
	for i:=0;i<len(s);i++{
		if byteT[i] != byteS[i]{
			return false
		}
	}
	return true
}