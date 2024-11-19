package leetcode

import "math"

/*
图像平滑器 是大小为 3 x 3 的过滤器，用于对图像的每个单元格平滑处理，平滑处理后单元格的值为该单元格的平均灰度。
每个单元格的  平均灰度 定义为：该单元格自身及其周围的 8 个单元格的平均值，结果需向下取整。（即，需要计算蓝色平滑器中 9 个单元格的平均值）。
如果一个单元格周围存在单元格缺失的情况，则计算平均灰度时不考虑缺失的单元格（即，需要计算红色平滑器中 4 个单元格的平均值）。
*/
func imageSmoother(img [][]int) [][]int {
	res := make([][]int, len(img))
	for i := 0; i < len(img); i++ {
		res[i] = make([]int, len(img[i]))
	}

	calc := func(i, j int) (int, int) {
		if i < 0 || j >= len(img[0]) || i >= len(img) || j < 0 {
			return 0, 0
		}
		return img[i][j], 1
	}

	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			sum, total := 0, 0
			for k := i - 1; k <= i+1; k++ {
				for l := j - 1; l <= j+1; l++ {
					vs, vt := calc(k, l)
					sum += vs
					total += vt
				}
			}
			res[i][j] = int(math.Floor(float64(sum) / float64(total)))
		}
	}
	return res
}
