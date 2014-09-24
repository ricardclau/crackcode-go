package stringarrays

func rotate90Matrix(matrix [][]int) [][]int {
	dim := len(matrix)
	for i := 0; i < dim/2; i++ {
		// save top
		tmp := make([]int, dim)
		for j := i; j < dim-i-1; j++ {
			tmp[j] = matrix[i][j]
			// left -> top
			matrix[i][j] = matrix[dim-j-1][i]
			// bottom -> left
			matrix[dim-j-1][i] = matrix[dim-i-1][dim-j-1]
			// right -> bottom
			matrix[dim-i-1][dim-j-1] = matrix[j][dim-i-1]
			// top -> right
			matrix[j][dim-i-1] = tmp[j]
		}
	}

	return matrix
}
