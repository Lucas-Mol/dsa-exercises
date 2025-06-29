/*
You are given pointer to the root of the binary search tree and two values  and . You need to return the lowest common ancestor (LCA) of  and  in the binary search tree.

image
In the diagram above, the lowest common ancestor of the nodes  and  is the node . Node  is the lowest node which has nodes  and  as descendants.

Function Description

Complete the function lca in the editor below. It should return a pointer to the lowest common ancestor node of the two values given.

lca has the following parameters:
- root: a pointer to the root node of a binary search tree
- v1: a node.data value
- v2: a node.data value

Input Format

The first line contains an integer, , the number of nodes in the tree.
The second line contains  space-separated integers representing  values.
The third line contains two space-separated integers,  and .

To use the test data, you will have to create the binary search tree yourself. Here on the platform, the tree will be created for you.

Constraints




The tree will contain nodes with data equal to  and .

Output Format

Return the a pointer to the node that is the lowest common ancestor of  and .

Sample Input

6
4 2 3 1 7 6
1 7
image

 and .

Sample Output

[reference to node 4]

Explanation

LCA of  and  is , the root in this case.
Return a pointer to the node.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lca(root *BinaryTreeNode, v1 int, v2 int) *BinaryTreeNode {
	if v1 > v2 {
		v1, v2 = v2, v1
	}

	for root != nil {
		if root.Value > v2 {
			root = root.Left
		} else if root.Value < v1 {
			root = root.Right
		} else {
			return root
		}
	}

	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadLine()

	input, _, _ := reader.ReadLine()
	inputs := strings.Split(string(input), " ")
	var nums []int

	for _, input := range inputs {
		num, _ := strconv.Atoi(input)
		nums = append(nums, num)
	}

	root := buildBinaryTree(nums)

	input, _, _ = reader.ReadLine()
	inputs = strings.Split(string(input), " ")

	nums = []int{}
	for _, input := range inputs {
		num, _ := strconv.Atoi(input)
		nums = append(nums, num)
	}

	fmt.Println(lca(root, nums[0], nums[1]).Value)
}

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func buildBinaryTree(nums []int) *BinaryTreeNode {
	root := BinaryTreeNode{Value: nums[0]}
	for i := 1; i < len(nums); i++ {
		insert(&root, nums[i])
	}

	return &root
}

func insert(root *BinaryTreeNode, val int) *BinaryTreeNode {
	if root == nil {
		return &BinaryTreeNode{Value: val}
	}

	if val < root.Value {
		root.Left = insert(root.Left, val)
	} else {
		root.Right = insert(root.Right, val)
	}

	return root
}
