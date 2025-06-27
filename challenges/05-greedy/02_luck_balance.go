/*
Lena is preparing for an important coding competition that is preceded by a number of sequential preliminary contests. Initially, her luck balance is 0. She believes in "saving luck", and wants to check her theory. Each contest is described by two integers,  and :

 is the amount of luck associated with a contest. If Lena wins the contest, her luck balance will decrease by ; if she loses it, her luck balance will increase by .
 denotes the contest's importance rating. It's equal to  if the contest is important, and it's equal to  if it's unimportant.
If Lena loses no more than  important contests, what is the maximum amount of luck she can have after competing in all the preliminary contests? This value may be negative.

Example



Contest		L[i]	T[i]
1		5	1
2		1	1
3		4	0
If Lena loses all of the contests, her will be . Since she is allowed to lose  important contests, and there are only  important contests, she can lose all three contests to maximize her luck at .

If , she has to win at least  of the  important contests. She would choose to win the lowest value important contest worth . Her final luck will be .

Function Description

Complete the luckBalance function in the editor below.

luckBalance has the following parameter(s):

int k: the number of important contests Lena can lose
int contests[n][2]: a 2D array of integers where each  contains two integers that represent the luck balance and importance of the  contest
Returns

int: the maximum luck balance achievable
Input Format

The first line contains two space-separated integers  and , the number of preliminary contests and the maximum number of important contests Lena can lose.
Each of the next  lines contains two space-separated integers,  and , the contest's luck balance and its importance rating.

Constraints

Sample Input

STDIN       Function
-----       --------
6 3         n = 6, k = 3
5 1         contests = [[5, 1], [2, 1], [1, 1], [8, 1], [10, 0], [5, 0]]
2 1
1 1
8 1
10 0
5 0
Sample Output

29
Explanation

There are  contests. Of these contests,  are important and she cannot lose more than  of them. Lena maximizes her luck if she wins the  important contest (where ) and loses all of the other five contests for a total luck balance of .
*/

package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'luckBalance' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. 2D_INTEGER_ARRAY contests
 */

func luckBalance(k int32, contests [][]int32) int32 {
	// Write your code here
	var importantContests []int
	var maxLuckBalance int32

	for _, contest := range contests {
		luckValue := contest[0]
		isImportant := contest[1] == 1

		if isImportant {
			importantContests = append(importantContests, int(luckValue))
		} else {
			maxLuckBalance += luckValue
		}
	}

	sort.Ints(importantContests)

	for i := range importantContests {
		loaseableElements := len(importantContests) - int(k)
		if i >= loaseableElements {
			maxLuckBalance += int32(importantContests[i])
			continue
		}
		maxLuckBalance -= int32(importantContests[i])
	}

	return maxLuckBalance
}

func main() {
	contests := make([][]int32, 6)

	contests[0] = []int32{5, 1}
	contests[1] = []int32{2, 1}
	contests[2] = []int32{1, 1}
	contests[3] = []int32{8, 1}
	contests[4] = []int32{10, 0}
	contests[5] = []int32{5, 0}

	fmt.Println("Expected 29, Received:", luckBalance(3, contests))
}
