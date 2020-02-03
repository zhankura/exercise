package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	trees := make([]*TreeNode, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		result = append(result, cur.Val)
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			result = append(result, cur.Val)
		}
		if cur.Right != nil {
			cur = cur.Right
			trees = append(trees, cur)
		}

	}
	return result
}

func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	trees := make([]*TreeNode, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		if cur.Right != nil {
			trees = append(trees, cur.Right)
		}
		if cur.Left != nil {
			trees = append(trees, cur.Left)
		}
		result = append(result, cur.Val)
	}
	return result
}

func postorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	cur := root
	treeFirst := make([]*TreeNode, 0)
	treeSecond := make([]*TreeNode, 0)
	treeFirst = append(treeFirst, cur)
	for len(treeFirst) != 0 {
		cur = treeFirst[len(treeFirst)-1]
		treeFirst = treeFirst[:len(treeFirst)-1]
		treeSecond = append(treeSecond, cur)
		if cur.Left != nil {
			treeFirst = append(treeFirst, cur.Left)
		}
		if cur.Right != nil {
			treeFirst = append(treeFirst, cur.Right)
		}
	}
	for len(treeSecond) != 0 {
		cur = treeSecond[len(treeSecond)-1]
		treeSecond = treeSecond[:len(treeSecond)-1]
		result = append(result, cur.Val)
	}
	return result
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	cur := root
	end := root
	nextEnd := root
	trees := make([]*TreeNode, 0)
	partResult := make([]int, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append([]*TreeNode{cur.Left}, trees...)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append([]*TreeNode{cur.Right}, trees...)
		}
		if cur == end {
			end = nextEnd
			result = append(result, partResult)
			partResult = partResult[0:0]
		}
	}
	return result
}

func partBalance(root *TreeNode, height int) (int, bool) {
	if root == nil {
		return height, true
	}
	leftHeight, leftBalance := partBalance(root.Left, height+1)
	rightHeight, rightBalance := partBalance(root.Right, height+1)
	ifBalance := leftBalance && rightBalance
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		ifBalance = false
	}
	return max(rightHeight, leftHeight), ifBalance
}

func isBalanced(root *TreeNode) bool {
	_, ifBalance := partBalance(root, 0)
	return ifBalance
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	cur := root
	end := root
	nextEnd := root
	trees := make([]*TreeNode, 0)
	partResult := make([]int, 0)
	trees = append(trees, cur)
	for len(trees) != 0 {
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append([]*TreeNode{cur.Left}, trees...)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append([]*TreeNode{cur.Right}, trees...)
		}
		if cur == end {
			end = nextEnd
			result = append([][]int{partResult}, result...)
			partResult = make([]int, 0)
		}
	}
	return result
}

func isValidBST(root *TreeNode) bool {
	result := make([]int, 0)
	if root == nil {
		return true
	}
	trees := make([]*TreeNode, 0)
	cur := root
	trees = append(trees, cur)
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		result = append(result, cur.Val)
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			result = append(result, cur.Val)
		}
		if cur.Right != nil {
			trees = append(trees, cur.Right)
			cur = cur.Right
		}
	}
	for i := 0; i < len(result)-1; i++ {
		if result[i] >= result[i+1] {
			return false
		}
	}
	return true
}

func kthSmallest2(root *TreeNode, k int) int {
	trees := make([]*TreeNode, 0)
	trees = append(trees, root)
	cur := root
	for len(trees) != 0 {
		for cur.Left != nil {
			trees = append(trees, cur.Left)
			cur = cur.Left
		}
		cur = trees[len(trees)-1]
		trees = trees[:len(trees)-1]
		if k > 1 {
			k -= 1
		} else {
			return cur.Val
		}
		for cur.Right == nil && len(trees) != 0 {
			cur = trees[len(trees)-1]
			trees = trees[:len(trees)-1]
			if k > 1 {
				k -= 1
			} else {
				return cur.Val
			}
		}
		if cur.Right != nil {
			cur = cur.Right
			trees = append(trees, cur)
		}
	}
	return -1
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	trees := make([]*TreeNode, 0)
	trees = append(trees, root)
	end, nextEnd := root, root
	partResult := make([]int, 0)
	earse := false
	for len(trees) != 0 {
		cur := trees[0]
		trees = trees[1:]
		partResult = append(partResult, cur.Val)
		if cur.Left != nil {
			nextEnd = cur.Left
			trees = append(trees, cur.Left)
		}
		if cur.Right != nil {
			nextEnd = cur.Right
			trees = append(trees, cur.Right)
		}
		if cur == end {
			end = nextEnd
			if earse {
				partResult = reverse(partResult)
			}
			result = append(result, partResult)
			partResult = make([]int, 0)
			earse = !earse
		}
	}
	return result
}

func partBuildTree1(preorder []int, inMap map[int]int, left int, right int, index *int) *TreeNode {
	if left > right {
		return nil
	}
	newNode := new(TreeNode)
	number := preorder[*index]
	*index += 1
	cur := inMap[number]
	leftNode := partBuildTree1(preorder, inMap, left, cur-1, index)
	rightNode := partBuildTree1(preorder, inMap, cur+1, right, index)
	newNode.Val, newNode.Left, newNode.Right = number, leftNode, rightNode
	return newNode
}

func buildTree1(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	}
	inMap := make(map[int]int)
	for index, value := range inorder {
		inMap[value] = index
	}
	a := 0
	root := partBuildTree1(preorder, inMap, 0, n-1, &a)
	return root
}

func partBuildTree(postorder []int, inMap map[int]int, left int, right int, index *int) *TreeNode {
	if left > right {
		return nil
	}
	newNode := new(TreeNode)
	number := postorder[*index]
	*index -= 1
	cur := inMap[number]
	rightNode := partBuildTree(postorder, inMap, cur+1, right, index)
	leftNode := partBuildTree(postorder, inMap, left, cur-1, index)
	newNode.Val, newNode.Left, newNode.Right = number, leftNode, rightNode
	return newNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(postorder)
	if n == 0 {
		return nil
	}
	inMap := make(map[int]int)
	for index, value := range inorder {
		inMap[value] = index
	}
	a := n - 1
	root := partBuildTree(postorder, inMap, 0, n-1, &a)
	return root
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	cur := root
	temp := cur
	for {
		if cur == nil {
			break
		}
		if cur.Val < val {
			temp = cur
			cur = cur.Right
		} else {
			temp = cur
			cur = cur.Left
		}
	}
	newNode := new(TreeNode)
	newNode.Val, newNode.Left, newNode.Right = val, nil, nil
	if temp == nil {
		return newNode
	}
	if temp.Val > val {
		temp.Left = newNode
	} else {
		temp.Right = newNode
	}
	return root
}
