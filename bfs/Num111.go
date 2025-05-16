package bfs

// *
// * Definition for a binary tree node.

func minDepth(root *TreeNode) int {
	var length int
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	q = append(q, root)
	length = 1
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			t := q[0]
			//出队列
			q = q[1:]
			if t.Left == nil && t.Right == nil {
				return length
			}
			if t.Left != nil {
				q = append(q, t.Left)
			}
			if t.Right != nil {
				q = append(q, t.Right)
			}
		}
		length++
	}
	return length
}
