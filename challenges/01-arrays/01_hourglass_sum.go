/*
Given a  2D array, , an hourglass is a subset of values with indices falling in the following pattern:

a b c
  d
e f g
There are  hourglasses in a  array. The  is the sum of the values in an hourglass. Calculate the hourglass sum for every hourglass in , then print the  hourglass sum.

Example


-9 -9 -9  1 1 1
 0 -9  0  4 3 2
-9 -9 -9  1 2 3
 0  0  8  6 6 0
 0  0  0 -2 0 0
 0  0  1  2 4 0
The  hourglass sums are:

-63, -34, -9, 12,
-10,   0, 28, 23,
-27, -11, -2, 10,
  9,  17, 25, 18
The highest hourglass sum is  from the hourglass beginning at row , column :

0 4 3
  1
8 6 6
Note: If you have already solved the Java domain's Java 2D Array challenge, you may wish to skip this challenge.

Function Description

Complete the function  with the following parameter(s):

: a 2-D array of integers
Returns

: the maximum hourglass sum
Input Format

Each of the  lines of inputs  contains  space-separated integers .

Constraints

Sample Input

1 1 1 0 0 0
0 1 0 0 0 0
1 1 1 0 0 0
0 0 2 4 4 0
0 0 0 2 0 0
0 0 1 2 4 0
Sample Output

19
Explanation

 contains the following hourglasses:

image

The hourglass with the maximum sum () is:

2 4 4
  2
1 2 4
*/

package main

import (
	"math"
)

func hourglassSum(arr [][]int32) int32 {
	var hourglassSums []int32

	// Iterate 2D arrays
	for row := range arr {
		for col := range arr[row] {
			// Avoid index out of bounds error. Exclusive for 6x6 arrays and 3x3 hourglass problem
			if row <= 3 && col <= 3 {
				hourglassSums = append(hourglassSums, getHourGlassSum(arr, row, col))
			}
		}
	}

	// Getting lower possible number that int32 could be. Avoiding errors when the sum is a negative number
	maxSum := int32(math.MinInt32)
	for _, val := range hourglassSums {
		if val > maxSum {
			maxSum = val
		}
	}

	return maxSum
}

func getHourGlassSum(arr [][]int32, i int, j int) int32 {
	var sum int32

	// Easy 'for loop' to design a hourglass
	// First, add all the top and bottom rows
	for col := range 3 {
		sum += arr[i][j+col]
		sum += arr[i+2][j+col]
	}

	// After, add the middle row
	sum += arr[i+1][j+1]

	return sum
}
