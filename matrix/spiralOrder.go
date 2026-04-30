package matrix

//54. 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//链接：https://leetcode.cn/problems/spiral-matrix/description/?envType=study-plan-v2&envId=top-100-liked
//讲解：https://leetcode.cn/problems/spiral-matrix/solutions/275716/shou-hui-tu-jie-liang-chong-bian-li-de-ce-lue-kan-/?envType=study-plan-v2&envId=top-100-liked
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	result := make([]int, 0)
	top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for top < bottom && left < right {
		//上边
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		//右边
		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right])
		}
		//下边
		for i := right; i > left; i-- {
			result = append(result, matrix[bottom][i])
		}
		//	左边
		for i := bottom; i > top; i-- {
			result = append(result, matrix[i][left])
		}
		top++
		bottom--
		left++
		right--
	}

	if left == right {
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][left])
		}
	} else if top == bottom {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
	}

	return result
}
