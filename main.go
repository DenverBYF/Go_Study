package main

import (
	"example.com/m/v2/structture"
	"fmt"
	"math"
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

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
	if len(energyDrinkA) == 0 || len(energyDrinkB) == 0 {
		return 0
	}
	dp := [2]int64{}
	dp[0] = int64(energyDrinkA[0])
	dp[1] = int64(energyDrinkB[0])
	n := len(energyDrinkA)
	for i := 1; i < n; i++ {
		num1 := dp[0]
		num2 := dp[1]
		dp[0] = Max(num1+int64(energyDrinkA[i]), num2)
		dp[1] = Max(num2+int64(energyDrinkB[i]), num1)
	}
	return Max(dp[0], dp[1])
}

func quickSort(input []int) []int {
	// 快速排序
	if len(input) <= 1 {
		return input
	}
	pivot := input[0]
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 1; i < len(input); i++ {
		if input[i] < pivot {
			left = append(left, input[i])
		} else {
			right = append(right, input[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

func minChanges(n int, k int) int {
	if n == k {
		return 0
	}
	if n < k {
		return -1
	}
	nBinStr := fmt.Sprintf("%b", n)
	kBinStr := fmt.Sprintf("%b", k)
	if len(nBinStr) < len(kBinStr) {
		return -1
	}
	res := 0
	i := len(nBinStr) - 1
	j := len(kBinStr) - 1
	for i >= 0 && j >= 0 {
		if nBinStr[i] == '1' && nBinStr[i] != kBinStr[j] {
			res++
		}
		if nBinStr[i] == '0' && nBinStr[i] != kBinStr[j] {
			return -1
		}
		i--
		j--
	}
	if i >= 0 {
		for ; i >= 0; i-- {
			if nBinStr[i] == '1' {
				res++
			}
		}
	}
	return res
}

func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	m := make(map[[26]int][]string)
	for _, str := range strs {
		hashKey := [26]int{}
		for i := 0; i < len(str); i++ {
			hashKey[str[i]-'a']++
		}
		m[hashKey] = append(m[hashKey], str)
	}

	for _, list := range m {
		res = append(res, list)
	}

	return res
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}

func judgeSquareSum(c int) bool {
	end := int(math.Ceil(math.Sqrt(float64(c))))
	m := make(map[int]bool)
	for i := 0; i <= end; i++ {
		m[i*i] = true
		if m[c-(i*i)] {
			return true
		}
	}
	return false
}

/*
给你两个 正 整数 x 和 y ，分别表示价值为 75 和 10 的硬币的数目。
Alice 和 Bob 正在玩一个游戏。每一轮中，Alice 先进行操作，Bob 后操作。每次操作中，玩家需要拿出价值 总和 为 115 的硬币。如果一名玩家无法执行此操作，那么这名玩家 输掉 游戏。
两名玩家都采取 最优 策略，请你返回游戏的赢家。
*/
func losingPlayer(x int, y int) string {
	i := 0
	for ; ; i++ {
		if x <= 0 || y < 4 {
			break
		}
		x -= 1
		y -= 4
	}
	if i%2 == 0 {
		return "Bob"
	}
	return "Alice"
}

func resultsArray(nums []int, k int) []int {
	res := make([]int, len(nums)-k+1)
	if k == 1 {
		res[0] = nums[0]
	}
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cnt++
		} else {
			cnt = 1
		}
		if i >= k-1 {
			if cnt >= k {
				res[i-k+1] = nums[i]
			} else {
				res[i-k+1] = -1
			}
		}
	}
	return res
}

/*
有一根长度为 n 个单位的木棍，棍上从 0 到 n 标记了若干位置
给你一个整数数组 cuts ，其中 cuts[i] 表示你需要将棍子切开的位置。
你可以按顺序完成切割，也可以根据需要更改切割的顺序。
每次切割的成本都是当前要切割的棍子的长度，切棍子的总成本是历次切割成本的总和。
对棍子进行切割将会把一根木棍分成两根较小的木棍（这两根木棍的长度和就是切割前木棍的长度）。请参阅第一个示例以获得更直观的解释。
返回切棍子的 最小总成本 。
*/
func minCost(n int, cuts []int) int {

	return 0
}

/*
有 n 个气球，编号为0 到 n - 1，每个气球上都标有一个数字，这些数字存在数组 nums 中。
现在要求你戳破所有的气球。戳破第 i 个气球，
你可以获得 nums[i - 1] * nums[i] * nums[i + 1] 枚硬币。
这里的 i - 1 和 i + 1 代表和 i 相邻的两个气球的序号。
如果 i - 1或 i + 1 超出了数组的边界，那么就当它是一个数字为 1 的气球。
求所能获得硬币的最大数量。
*/
func maxCoins(nums []int) int {
	return 0
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
给你一个 二进制 字符串 s 和一个整数 k。
如果一个 二进制字符串 满足以下任一条件，则认为该字符串满足 k 约束：
字符串中 0 的数量最多为 k。
字符串中 1 的数量最多为 k。
返回一个整数，表示 s 的所有满足 k 约束的子字符串的数量。
*/
func countKConstraintSubstrings(s string, k int) int {
	res := 0
	/*for i := 0; i < len(s); i++ {
		val0 := 0
		val1 := 0
		for j := i; j < len(s); j++ {
			if s[j] == '0' {
				val0++
			}
			if s[j] == '1' {
				val1++
			}
			if val0 <= k || val1 <= k {
				res++
			}
			if val0 > k && val1 > k {
				break
			}
		}
	}*/
	val0 := 0
	val1 := 0
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			val0++
		}
		if s[i] == '1' {
			val1++
		}
		for val0 > k && val1 > k {
			if s[left] == '0' {
				val0--
			}
			if s[left] == '1' {
				val1--
			}
			left++
		}
		res += i - left + 1
	}
	return res
}

func main() {
	//fmt.Println(fmt.Sprintf("%b", 44), fmt.Sprintf("%b", 2))
	//language.TestChan()

	/*lRUCache := structture.Constructor(3)
	lRUCache.Put(1, 1)
	lRUCache.Put(2, 2)
	lRUCache.Put(3, 3)
	lRUCache.Put(4, 4)
	fmt.Println(lRUCache.Get(4))
	fmt.Println(lRUCache.Get(3))
	fmt.Println(lRUCache.Get(2))
	fmt.Println(lRUCache.Get(1))
	lRUCache.Put(5, 5)
	fmt.Println(lRUCache.Get(1))
	fmt.Println(lRUCache.Get(2))
	fmt.Println(lRUCache.Get(3))
	fmt.Println(lRUCache.Get(4))
	fmt.Println(lRUCache.Get(5))*/
}
