/*
The height of a binary tree is the number of edges between the tree's root and its furthest leaf. For example, the following binary tree is of height :

image
Function Description

Complete the getHeight or height function in the editor. It must return the height of a binary tree as an integer.

getHeight or height has the following parameter(s):

root: a reference to the root of a binary tree.
Note -The Height of binary tree with single node is taken as zero.

# Input Format

The first line contains an integer , the number of nodes in the tree.
Next line contains  space separated integer where th integer denotes node[i].data.

Note: Node values are inserted into a binary search tree before a reference to the tree's root node is passed to your function. In a binary search tree, all nodes on the left branch of a node are less than the node value. All values on the right branch are greater than the node value.

# Constraints

# Output Format

Your function should return a single integer denoting the height of the binary tree.

# Sample Input

image

# Sample Output

3
Explanation

The longest root-to-leaf path is shown below:

image

There are  nodes in this path that are connected by  edges, meaning our binary tree's .
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func getHeight(root *BinaryTreeNode) int {
	if root == nil {
		return -1
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	fmt.Println(getHeight(root))
}
