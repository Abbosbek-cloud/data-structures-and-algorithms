package main

/* 53. Maximum Subarray problem
 *
 * Given an integer array nums,
 * find the subarray with the
 * largest sum, and return its sum
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4]
 * Output: 6
 * Explanation: The subarray [4,-1,2,1] has the largest sum 6.
 */

type Solution struct {}

type ISolution interface {
	MaxSubarrayBad(nums []int) int 
	MaxSubarrayGood(nums []int) int 
	MaxSubarrayBest(nums []int) int 
}

func NewSolution() ISolution {
	return &Solution{}
}

/*
 * 1. Bad Solution: Brute Force Approach
 * 
 * This approach uses nested loops to calculate 
 * the sum of all possible subarrays.
 *
 * Complexity: 
 *
 * Time Complexity: O(n^3)
 * Space Complexity: O(1)
*/

func (s Solution) MaxSubarrayBad(nums []int) int {
    maxSum := nums[0]
    n := len(nums)
    
    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            currentSum := 0
            for k := i; k <= j; k++ {
                currentSum += nums[k]
            }
            if currentSum > maxSum {
                maxSum = currentSum
            }
        }
    }
    
    return maxSum
}

/*
 * 2. Good Solution: Optimized Brute Force
 * 
 * This approach eliminates the innermost loop 
 * by calculating the subarray sum on the fly.
 *
 * Complexity: 
 *
 * Time Complexity: O(n^2)
 * Space Complexity: O(1)
*/

func (s Solution) MaxSubarrayGood(nums []int) int {
    maxSum := nums[0]
    n := len(nums)
    
    for i := 0; i < n; i++ {
        currentSum := 0
        for j := i; j < n; j++ {
            currentSum += nums[j]
            if currentSum > maxSum {
                maxSum = currentSum
            }
        }
    }
    
    return maxSum
}

/*
 * 3. Best Solution: Kadane's Algorithm
 * 
 * Kadane's algorithm is the most efficient approach, using dynamic
 * programming to find the maximum subarray sum in a single pass.
 *
 * Complexity: 
 *
 * Time Complexity: O(n)
 * Space Complexity: O(1)
*/

func (s Solution) MaxSubarrayBest(nums []int) int {
    maxSum := nums[0]
    currentSum := nums[0]
    
    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum+nums[i])
        maxSum = max(maxSum, currentSum)
    }
    
    return maxSum
}

// Utility function for maximum of two integers
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/*
 * |------------------------|-----------------|------------------|--------------------------------------|
 * | Approach               | Time Complexity | Space Complexity | Notes                                |
 * |------------------------|-----------------|------------------|--------------------------------------|
 * | Brute Force            | O(n^3)          | O(1)             | Extremely slow for large inputs.     |
 * |------------------------|-----------------|------------------|--------------------------------------|
 * | Optimized Brute Force  | O(n^2)          | O(1)             | Better but still inefficient.        |
 * |------------------------|-----------------|------------------|--------------------------------------|
 * | Kadane's Algorithm     | O(n)            | O(1)             | Best approach; works in linear time. |
 * |------------------------|-----------------|------------------|--------------------------------------|
 * 
 */


