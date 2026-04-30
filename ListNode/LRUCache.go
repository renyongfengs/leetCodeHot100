package ListNode

//146. LRU 缓存
//请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

//链接：https://leetcode.cn/problems/lru-cache/description/?envType=study-plan-v2&envId=top-100-liked
//讲解：https://leetcode.cn/problems/lru-cache/?envType=study-plan-v2&envId=top-100-liked

type LND struct {
	key, val  int
	pre, next *LND
}
type LRUCache struct {
	cap   int
	dummy *LND
	cache map[int]*LND
}

func Constructor(capacity int) LRUCache {
	dummy := new(LND)
	dummy.next = dummy
	dummy.pre = dummy

	return LRUCache{
		cap:   capacity,
		dummy: dummy,
		cache: map[int]*LND{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	//删除节点，并把节点移到首部
	this.remove(node)
	this.pushFront(node)
	return node.val

}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cache[key]
	if ok {
		node.val = value
		this.remove(node)
	} else {
		//不存在，要插入
		//判断是否超出容量大小
		if len(this.cache) >= this.cap {
			//删除不活跃的节点（尾节点）
			tmp := this.dummy.pre
			delete(this.cache, tmp.key)
			this.remove(tmp)
		}

		node = &LND{key, value, nil, nil}
		this.cache[key] = node
	}
	//把节点移到首部
	this.pushFront(node)
}

//remove 删除节点
func (this *LRUCache) remove(node *LND) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

//pushFront 把节点插入头部(不破坏环)
func (this *LRUCache) pushFront(node *LND) {
	node.pre = this.dummy
	node.next = this.dummy.next

	node.pre.next = node
	node.next.pre = node
}
