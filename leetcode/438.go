package leetcode

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
abab ab
*/
func findAnagrams(s string, p string) []int {
	vis := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		vis[p[i]]++
	}
	ans := make([]int, 0)
	left, right := 0, 0
	for right < len(s) {
		if vis[s[right]] > 0 {
			vis[s[right]]--
			if right-left+1 == len(p) {
				ans = append(ans, left)
				vis[s[left]]++
				left++
			}
			right++
		} else {
			vis[s[left]]++
			left++
		}
	}

	return ans
}
