/*
A left rotation operation on an array shifts each of the array's elements  unit to the left. For example, if  left rotations are performed on array , then the array would become . Note that the lowest index item moves to the highest index in a rotation. This is called a circular array.

Given an array  of  integers and a number, , perform  left rotations on the array. Return the updated array to be printed as a single line of space-separated integers.

Function Description

Complete the function rotLeft in the editor below.

rotLeft has the following parameter(s):

int a[n]: the array to rotate
int d: the number of rotations
Returns

int a'[n]: the rotated array
Input Format

The first line contains two space-separated integers  and , the size of  and the number of left rotations.
The second line contains  space-separated integers, each an .

Constraints

Sample Input

5 4
1 2 3 4 5
Sample Output

5 1 2 3 4
Explanation

When we perform  left rotations, the array undergoes the following sequence of changes:


*/

package main

/*
 * Complete the 'rotLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER d
 */

//Easy way: using Go features as slices and subslices
// func rotLeft(a []int32, d int32) []int32 {
//     return append(a[d:], a[:d]...)
// }

// Hard way: implementing from scratch
func rotLeft(a []int32, d int32) []int32 {
	n := int32(len(a))

	reverse(a, 0, d-1)
	reverse(a, d, n-1)
	reverse(a, 0, n-1)

	return a
}

func reverse(arr []int32, startIndex int32, endIndex int32) {
	for i, j := startIndex, endIndex; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
