package tree

// 递归方法
func PreorderInorderTraversal1(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	res = append(res, PreorderInorderTraversal1(root.Left)...)
	res = append(res, PreorderInorderTraversal1(root.Right)...)
	return
}

// 迭代方法
func PreorderInorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for len(stack) >0 || root != nil {
		for root !=nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack) - 1].Right
		stack = stack[:len(stack) - 1]
	}
	return
}
