/*
Given an array of integers, find the subset of non-adjacent elements with the maximum sum. Calculate the sum of that subset. It is possible that the maximum sum is , the case when all elements are negative.

Example

The following subsets with more than  element exist. These exclude the empty subset and single element subsets which are also valid.

Subset      Sum
[-2, 3, 5]   6
[-2, 3]      1
[-2, -4]    -6
[-2, 5]      3
[1, -4]     -3
[1, 5]       6
[3, 5]       8
The maximum subset sum is . Note that any individual element is a subset as well.


In this case, it is best to choose no element: return .

Function Description

Complete the  function in the editor below.

maxSubsetSum has the following parameter(s):

int arr[n]: an array of integers
Returns
- int: the maximum subset sum

Input Format

The first line contains an integer, .
The second line contains  space-separated integers .

Constraints

Sample Input 0

5
3 7 4 6 5
Sample Output 0

13
Explanation 0

Our possible subsets are  and . The largest subset sum is  from subset

Sample Input 1

5
2 1 5 8 4
Sample Output 1

11
Explanation 1

Our subsets are  and . The maximum subset sum is  from the first subset listed.

Sample Input 2

5
3 5 -7 8 10
Sample Output 2

15
Explanation 2

Our subsets are  and . The maximum subset sum is  from the fifth subset listed.
*/

package main

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return max(0, arr[0])
	}

	dp := make([]int32, n)
	dp[0] = max(0, arr[0])
	dp[1] = max(dp[0], arr[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+arr[i])
	}

	return dp[n-1]
}
