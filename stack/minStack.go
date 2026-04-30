package stack

import (
	"math"
)

//155. 最小栈
//设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//实现 MinStack 类:
//MinStack() 初始化堆栈对象。
//void push(int val) 将元素val推入堆栈。
//void pop() 删除堆栈顶部的元素。
//int top() 获取堆栈顶部的元素。
//int getMin() 获取堆栈中的最小元素。
//链接；https://leetcode.cn/problems/min-stack/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/min-stack/solutions/2974438/ben-zhi-shi-wei-hu-qian-zhui-zui-xiao-zh-x0g8/?envType=study-plan-v2&envId=top-100-liked

type elem struct {
	val, preMin int
}

type MinStack struct {
	minStack []elem
}

func Constructor() MinStack {
	return MinStack{
		minStack: make([]elem, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.minStack = append(this.minStack, elem{
		val:    val,
		preMin: min(val, this.GetMin()),
	})
}

func (this *MinStack) Pop() {
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.minStack[len(this.minStack)-1].val
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return math.MaxInt
	}
	return this.minStack[len(this.minStack)-1].preMin
}
