package cn
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
    // 保证不是最后一个节点，说明node.next != nil

	// 删除当前节点的值，不好操作，可以删除后一节点的值，把后一节点的值赋值给当前节点

	// 后一节点值赋值到当前节点
	node.Val = node.Next.Val
	// 删除后一几点
	node.Next = node.Next.Next
}
//leetcode submit region end(Prohibit modification and deletion)
