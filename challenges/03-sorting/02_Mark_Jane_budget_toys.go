/*

Mark and Jane are very happy after having their first child. Their son loves toys, so Mark wants to buy some. There are a number of different toys lying in front of him, tagged with their prices. Mark has only a certain amount to spend, and he wants to maximize the number of toys he buys with this money. Given a list of toy prices and an amount to spend, determine the maximum number of gifts he can buy.

Note Each toy can be purchased only once.

Example


The budget is  units of currency. He can buy items that cost  for , or  for  units. The maximum is  items.

Function Description

Complete the function maximumToys in the editor below.

maximumToys has the following parameter(s):

int prices[n]: the toy prices
int k: Mark's budget
Returns

int: the maximum number of toys
Input Format

The first line contains two integers,  and , the number of priced toys and the amount Mark has to spend.
The next line contains  space-separated integers

Constraints




A toy can't be bought multiple times.

Sample Input

7 50
1 12 5 111 200 1000 10
Sample Output

4
Explanation

He can buy only  toys at most. These toys have the following prices: .

*/

package main

/*
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

//Easy solution:
//  import (
//     "sort"
// )

// func maximumToys(prices []int32, k int32) int32 {
//     intPrices := make([]int, len(prices))
//     for i, v := range prices {
//         intPrices[i] = int(v)
//     }

//     sort.Ints(intPrices)

//     var toysCount int32
//     for _, val := range intPrices {
//         if int32(val) < k {
//             k -= int32(val)
//             toysCount++
//         } else {
//             break
//         }
//     }

//     return toysCount
// }

// Hard solution: implementing quicksort manually
func maximumToys(prices []int32, k int32) int32 {
	recursiveQuick(&prices, 0, len(prices)-1)

	var toysCount int32
	for _, val := range prices {
		if val < k {
			k -= val
			toysCount++
		} else {
			break
		}
	}

	return toysCount
}

func recursiveQuick(a *[]int32, left int, right int) {
	if left < right {
		pi := partition(a, left, right)
		recursiveQuick(a, left, pi-1)
		recursiveQuick(a, pi+1, right)
	}
}

func partition(a *[]int32, left int, right int) int {
	pivot := (*a)[right] //dumb pivot choice
	i := left - 1

	for j := left; j < right; j++ {
		if (*a)[j] <= pivot {
			i++
			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		}
	}

	(*a)[i+1], (*a)[right] = (*a)[right], (*a)[i+1]
	return i + 1
}
