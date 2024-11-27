package leetcode

/*
给你一个整数 n 和一个二维整数数组 queries。
有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。
queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。
返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。
*/
func ShortestDistanceAfterQueries(n int, queries [][]int) []int {
	res := make([]int, 0)
	g := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		g[i] = append(g[i], i+1)
	}
	for i := 0; i < len(queries); i++ {
		g[queries[i][0]] = append(g[queries[i][0]], queries[i][1])
		res = append(res, bfs(g, n))
	}
	return res
}

func bfs(g map[int][]int, n int) int {
	q := make([]int, 0)
	vis := make(map[int]bool)
	q = append(q, 0)
	vis[0] = true
	dist := make([]int, n)
	dist[0] = 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			for _, child := range g[node] {
				if !vis[child] {
					q = append(q, child)
					vis[child] = true
					dist[child] = dist[node] + 1
				}
			}
		}
		q = q[size:]
	}
	return dist[n-1]
}

func myMin3243(a, b int) int {
	if a < b {
		return a
	}
	return b
}
