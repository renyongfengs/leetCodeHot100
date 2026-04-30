package ListNode

import (
	"container/heap"
)

//23. 合并 K 个升序链表
//给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。
//链接：https://leetcode.cn/problems/merge-k-sorted-lists/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/merge-k-sorted-lists/solutions/2384305/liang-chong-fang-fa-zui-xiao-dui-fen-zhi-zbzx/?envType=study-plan-v2&envId=top-100-liked

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) <= 0 {
		return nil
	}
	//筛选出头节点不等于nil
	ls := make(ListNodeList, 0)
	for _, node := range lists {
		if node != nil {
			ls = append(ls, node)
		}
	}
	heap.Init(&ls)

	dump := new(ListNode)
	cur := dump

	for ls.Len() > 0 {
		pop := heap.Pop(&ls).(*ListNode)
		if pop.Next != nil {
			heap.Push(&ls, pop.Next)
		}

		cur.Next = pop
		cur = pop
	}

	return dump.Next
}

type ListNodeList []*ListNode

func (l ListNodeList) Len() int {
	return len(l)
}

// Less 小根堆
func (l ListNodeList) Less(i, j int) bool {
	return l[i].Val < l[j].Val
}
func (l ListNodeList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l *ListNodeList) Push(x interface{}) {
	*l = append(*l, x.(*ListNode))
}
func (l *ListNodeList) Pop() interface{} {
	old := *l
	out := old[len(old)-1]
	*l = old[:len(old)-1]
	return out
}

//递归
func mergeKLists2(lists []*ListNode) *ListNode {
	length := len(lists)
	if length == 0 {
		return nil
	} else if length == 1 {
		return lists[0]
	}

	left := mergeKLists2(lists[:length/2])
	right := mergeKLists2(lists[length/2:])

	return merge2Lists(left, right)
}

func merge2Lists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dump := &ListNode{}
	cur := dump

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}

		cur = cur.Next
	}

	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}

	return dump.Next

}
