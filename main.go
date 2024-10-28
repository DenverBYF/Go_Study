package main

import (
	"example.com/m/v2/structture"
)

/*
给你一个包含若干星号 * 的字符串 s 。

在一步操作中，你可以：

选中 s 中的一个星号。
移除星号 左侧 最近的那个 非星号 字符，并移除该星号自身。
返回移除 所有 星号之后的字符串。

输入：s = "leet**cod*e"
输出："lecoe"
解释：从左到右执行移除操作：
- 距离第 1 个星号最近的字符是 "leet**cod*e" 中的 't' ，s 变为 "lee*cod*e" 。
- 距离第 2 个星号最近的字符是 "lee*cod*e" 中的 'e' ，s 变为 "lecod*e" 。
- 距离第 3 个星号最近的字符是 "lecod*e" 中的 'd' ，s 变为 "lecoe" 。
不存在其他星号，返回 "lecoe" 。
*/
func removeStars(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "*" {
			if len(res) == 0 {
				continue
			}
			res = res[0 : len(res)-1]
		} else {
			res += string(s[i])
		}
	}
	return res
}

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。

找出该数组中满足其总和大于等于 target 的长度最小的
子数组

	[numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
*/
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	start := 0
	end := 0
	total := nums[0]
	for start < len(nums) {
		if total >= target {
			// 符合要求，记录窗口大小
			if res == 0 || (end-start+1 < res) {
				res = end - start + 1
			}
			// 移动左边窗口
			total -= nums[start]
			start++
			continue
		} else {
			// 移动右边窗口
			if end < len(nums)-1 {
				// 能够移动右边的
				end++
				total += nums[end]
			} else {
				// 到底了，可以直接返回
				break
			}
		}
	}
	return res
}

/*
给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0)
	index1 := 0
	index2 := 0
	for index1 < m && index2 < n {
		if nums1[index1] < nums2[index2] {
			res = append(res, nums1[index1])
			index1++
		} else {
			res = append(res, nums2[index2])
			index2++
		}
	}
	if index1 < m {
		res = append(res, nums1[index1:]...)
	}
	if index2 < n {
		res = append(res, nums2[index2:]...)
	}
	copy(nums1, res)
}

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。
*/
func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := make(map[int32]int32)
	for _, c := range magazine {
		magazineMap[c] += 1
	}

	for _, c := range ransomNote {
		if magazineMap[c] > 0 {
			magazineMap[c] -= 1
		} else {
			return false
		}
	}
	return true
}

/*
给你两个整数数组 nums1 和 nums2，长度分别为 n 和 m。同时给你一个正整数 k。
如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。
返回 优质数对 的总数。
*/
func numberOfPairs(nums1 []int, nums2 []int, k int) int64 {
	return 0
}

/*
在本问题中，有根树指满足以下条件的 有向 图。该树只有一个根节点，所有其他节点都是该根节点的后继。该树除了根节点之外的每一个节点都有且只有一个父节点，而根节点没有父节点。
输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。附加的边包含在 1 到 n 中的两个不同顶点间，这条附加的边不属于树中已存在的边。
结果图是一个以边组成的二维数组 edges 。 每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是 vi 的一个父节点。
返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。
*/
func findRedundantDirectedConnection(edges [][]int) []int {
	return nil
}

func main() {
	val := structture.Link{}
	val.Val = 1
	val.Next = &structture.Link{Val: 2, Next: &structture.Link{Val: 3, Next: &structture.Link{Val: 4}}}
	res := val.Reserve()
	res.Print()
}
