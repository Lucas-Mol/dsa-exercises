/*
You are given an array and you need to find number of tripets of indices  such that the elements at those indices are in geometric progression for a given common ratio  and .

Example


There are  and  at indices  and . Return .

Function Description

Complete the countTriplets function in the editor below.

countTriplets has the following parameter(s):

int arr[n]: an array of integers
int r: the common ratio
Returns

int: the number of triplets
Input Format

The first line contains two space-separated integers  and , the size of  and the common ratio.
The next line contains  space-seperated integers .

Constraints

Sample Input 0

4 2
1 2 2 4
Sample Output 0

2
Explanation 0

There are  triplets in satisfying our criteria, whose indices are  and

Sample Input 1

6 3
1 3 9 9 27 81
Sample Output 1

6
Explanation 1

The triplets satisfying are index , , , ,  and .

Sample Input 2

5 5
1 5 5 25 125
Sample Output 2

4
Explanation 2

The triplets satisfying are index , , , .

*/

package main

// Complete the countTriplets function below.
func countTriplets(arr []int64, r int64) int64 {
	countLeft := make(map[int64]int64)
	countRight := make(map[int64]int64)

	// Inicialmente contamos todos os elementos como "à direita"
	for _, val := range arr {
		countRight[val]++
	}

	var total int64 = 0

	for _, mid := range arr {
		countRight[mid]-- // vamos processar esse número, tiramos ele da direita

		// Se mid pode ser o termo do meio de uma progressão
		if mid%r == 0 {
			left := mid / r
			right := mid * r

			// Contamos quantas vezes vimos left e ainda restam right
			total += countLeft[left] * countRight[right]
		}

		// Agora esse mid já foi visto
		countLeft[mid]++
	}

	return total
}
