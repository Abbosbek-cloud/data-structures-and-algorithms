package main

/* 54. Spiral Matrix
 *
 * Given an m x n matrix, return all elements of the matrix in spiral order.
 *
 * Example:
 *
 * Input: matrix = [[1, 2, 3],
 *                  [4, 5, 6],
 *                  [7, 8, 9]]
 *
 * Output: [1, 2, 3, 6, 9, 8, 7, 4, 5]
 */

type Solution struct {}

type ISolution interface {
	SpiralOrderBad(matrix [][]int) []int
	SpiralOrderGood(matrix [][]int) []int
	SpiralOrderBest(matrix [][]int) []int
}

func NewSolution() ISolution {
	return &Solution{}
}

/*
 * 1. Bad Solution: Use Visited Matrix
 * 
 * This approach uses an auxiliary visited matrix of the same size as the input
 * to mark which elements have already been visited. It follows the spiral order
 * by checking direction and boundaries at each step.
 *
 * Complexity:
 * Time Complexity: O(m * n) — Visiting each element once.
 * Space Complexity: O(m * n) — For the visited matrix.
 */

func (s Solution) SpiralOrderBad(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // Right, Down, Left, Up
	result := []int{}
	x, y, dir := 0, 0, 0

	for len(result) < m*n {
		if !visited[x][y] {
			result = append(result, matrix[x][y])
			visited[x][y] = true
		}

		nextX, nextY := x+directions[dir][0], y+directions[dir][1]
		if nextX < 0 || nextY < 0 || nextX >= m || nextY >= n || visited[nextX][nextY] {
			dir = (dir + 1) % 4
		}
		x, y = x+directions[dir][0], y+directions[dir][1]
	}

	return result
}

/*
 * 2. Good Solution: Layer by Layer Traversal
 *
 * This approach traverses the matrix layer by layer, starting from the outermost
 * layer and moving inward. It uses simple loops to traverse the rows and columns
 * of each layer.
 *
 * Complexity:
 * Time Complexity: O(m * n) — Visiting each element once.
 * Space Complexity: O(1) — No extra space except the result array.
 */

func (s Solution) SpiralOrderGood(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := []int{}
	top, bottom, left, right := 0, m-1, 0, n-1

	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			// Traverse bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// Traverse left column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

/*
 * 3. Best Solution: Optimized Layer Traversal
 *
 * This approach is similar to the "Good Solution" but removes unnecessary boundary
 * checks and organizes the traversal logic for simplicity and clarity.
 *
 * Complexity:
 * Time Complexity: O(m * n) — Visiting each element once.
 * Space Complexity: O(1) — No extra space except the result array.
 */

func (s Solution) SpiralOrderBest(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)
	top, bottom, left, right := 0, m-1, 0, n-1

	for top <= bottom && left <= right {
		// Traverse top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Traverse bottom row if still in bounds
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Traverse left column if still in bounds
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}

/*
 * |------------------------|-----------------|------------------|-------------------------------------|
 * | Approach               | Time Complexity | Space Complexity | Notes                               |
 * |------------------------|-----------------|------------------|-------------------------------------|
 * | Use Visited Matrix     | O(m * n)        | O(m * n)         | Inefficient due to extra space.     |
 * |------------------------|-----------------|------------------|-------------------------------------|
 * | Layer by Layer         | O(m * n)        | O(1)             | Efficient and simple.               |
 * |------------------------|-----------------|------------------|-------------------------------------|
 * | Optimized Layer        | O(m * n)        | O(1)             | Most elegant and clear solution.    |
 * |------------------------|-----------------|------------------|-------------------------------------|
 */
