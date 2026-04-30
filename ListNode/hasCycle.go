package ListNode

//141. 环形链表
//给你一个链表的头节点 head ，判断链表中是否有环。
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
//注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
//如果链表中存在环 ，则返回 true 。 否则，返回 false 。
//链接：https://leetcode.cn/problems/linked-list-cycle/description/?envType=study-plan-v2&envId=top-100-liked

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head // 乌龟和兔子同时从起点出发
	for fast != nil && fast.Next != nil {
		slow = slow.Next      // 乌龟走一步
		fast = fast.Next.Next // 兔子走两步
		if fast == slow {     // 兔子追上乌龟（套圈），说明有环
			return true
		}
	}
	return false // 访问到了链表末尾，无环
}
