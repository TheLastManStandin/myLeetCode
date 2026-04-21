package main

func maximumAmount(coins [][]int) int {
	rows := len(coins)
	cols := len(coins[0])
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			currentCoins := coins[i][j]
			if i == 0 && j == 0 {
				matrix[i][j] = currentCoins
			} else if i == 0 {
				matrix[i][j] = matrix[i][j-1] + currentCoins
			} else if j == 0 {
				matrix[i][j] = matrix[i-1][j] + currentCoins
			} else {
				matrix[i][j] = max(matrix[i-1][j], matrix[i][j-1]) + currentCoins
			}

		}
	}

	return matrix[rows-1][cols-1]
}

//func main() {
//	a := [][]int{
//		{0, 1, -1},
//		{1, -2, 3},
//		{2, -3, 4},
//	}
//
//	fmt.Println(maximumAmount(a))
//}
