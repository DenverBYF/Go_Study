package leetcode

/*
给定一个字符串 s ，请你找出其中不含有重复字符的最长子串的长度。
pwwkew wke
dvdf vdf

tmmzuxt mzuxt
*/
func LengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	res := 0
	left, right := 0, 1
	vis := make(map[byte]bool)
	vis[s[0]] = true
	for right < len(s) && left < len(s) {
		if !vis[s[right]] {
			vis[s[right]] = true
			right++
		} else {
			res = myMax3(res, right-left)
			for left < right && s[left] != s[right] {
				vis[s[left]] = false
				left++
			}
			left++
			right++
		}
	}
	res = myMax3(res, right-left)
	return res
}

func myMax3(a, b int) int {
	if a > b {
		return a
	}
	return b
}
