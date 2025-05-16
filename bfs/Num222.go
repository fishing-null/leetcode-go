package bfs

var cnt int

func countNodes(root *TreeNode) int {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	cnt := 1
	if root == nil {
		return 0
	}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Left != nil {
				q = append(q, cur.Left)
				cnt++
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
				cnt++
			}
		}
	}
	return cnt
}
