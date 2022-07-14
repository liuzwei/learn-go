package util

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Construct(ary []any) *TreeNode {
	if len(ary) == 0 {
		return nil
	}
	root := &TreeNode{
		ary[0].(int), nil, nil,
	}

	queue := []*TreeNode{root}

	i := 1
	for i < len(ary) {
		current := queue[0]
		queue = queue[1:]
		// 取两个值
		if i < len(ary) && ary[i] != nil {
			current.Left = &TreeNode{ary[i].(int), nil, nil}
			queue = append(queue, current.Left)
		}
		i++
		if i < len(ary) && ary[i] != nil {
			current.Right = &TreeNode{ary[i].(int), nil, nil}
			queue = append(queue, current.Right)
		}
		i++
	}
	return root
}
