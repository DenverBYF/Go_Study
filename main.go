package main

import (
	"example.com/m/v2/structture"
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func getPath(head *structture.Tree, path string, res *[]string) {
	path += strconv.Itoa(head.Val.(int))
	if head.Left == nil && head.Right == nil {
		*res = append(*res, path[1:])
		return
	}
	if head.Left != nil {
		getPath(head.Left, path, res)
	}
	if head.Right != nil {
		getPath(head.Right, path, res)
	}
}

func validStrings(n int) []string {
	head := &structture.Tree{
		Val: 1,
	}

	res := make([]string, 0)
	l := make([]*structture.Tree, 0)
	l = append(l, head)
	for i := 0; i < n; i++ {
		ll := len(l)
		for j := 0; j < ll; j++ {
			if l[j].Val == 1 {
				l[j].Left = &structture.Tree{Val: 1}
				l[j].Right = &structture.Tree{Val: 0}
				l = append(l, l[j].Left)
				l = append(l, l[j].Right)
			} else {
				l[j].Left = &structture.Tree{Val: 1}
				l = append(l, l[j].Left)
			}
		}
		l = l[ll:]
	}

	getPath(head, "", &res)
	return res
}

/*
给你一个仅由数字组成的字符串 s，在最多交换一次 相邻 且具有相同 奇偶性 的数字后，返回可以得到的
字典序最小的字符串。
如果两个数字都是奇数或都是偶数，则它们具有相同的奇偶性。例如，5 和 9、2 和 4 奇偶性相同，而 6 和 9 奇偶性不同。
*/
func getSmallestString(s string) string {
	for i := 0; i < len(s)-1; i++ {
		iNow, _ := strconv.Atoi(string(s[i]))
		iNext, _ := strconv.Atoi(string(s[i+1]))
		if iNow > iNext && (iNow+iNext)%2 == 0 {
			return s[:i] + string(s[i+1]) + string(s[i]) + s[i+2:]
		}
	}
	return s
}

func kthSmallest(root *TreeNode, k int) int {
	i := 0
	st := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(st) > 0 {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		node := st[len(st)-1]
		i++
		if i == k {
			return node.Val
		}
		st = st[:len(st)-1]
		cur = node.Right
	}
	return -1
}

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		l := len(q)
		res = append(res, q[l-1].Val)
		for i := 0; i < l; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		q = q[l:]
	}
	return res
}

func flatten(root *TreeNode) {
	linkList := make([]*TreeNode, 0)
	linkList = append(linkList, root)
	helper := make([]*TreeNode, 0)
	for len(linkList) > 0 {
		node := linkList[len(linkList)-1]
		linkList = linkList[:len(linkList)-1]
		helper = append(helper, node)
		if node.Right != nil {
			linkList = append(linkList, node.Right)
		}
		if node.Left != nil {
			linkList = append(linkList, node.Left)
		}
	}
	for i := 0; i < len(helper)-1; i++ {
		helper[i].Left = nil
		helper[i].Right = helper[i+1]
	}
}

func flattenV2(root *TreeNode) {
	for root != nil {
		if root.Left == nil {
			root = root.Right
		} else {
			leftSub := root.Left
			for leftSub.Right != nil {
				leftSub = leftSub.Right
			}

			leftSub.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			root = root.Right
		}
	}
}

func change(s []int) {
	s = append(s, 3)
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1)
	s2 = append(s2, 5, 6, 7)
	fmt.Println(s1)

	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)

	slice := make([]int, 5, 5)
	slice[0] = 1
	slice[1] = 2
	change(slice)
	fmt.Println(slice)
	change(slice[0:2])
	fmt.Println(slice)

	x := make([]int, 2, 10)
	_ = x[6:10] //1
	_ = x[6:]   //2
	_ = x[2:]   //3
}
