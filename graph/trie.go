package graph

//208. 实现 Trie (前缀树)
//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

type Node struct {
	son [26]*Node //标记每个字符
	end bool      //表示是不是结束字符
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, ch := range word {
		ch -= 'a'
		if cur.son[ch] == nil {
			cur.son[ch] = &Node{} //第一次初始化
		}

		cur = cur.son[ch]
	}
	cur.end = true
}

//find 0:找不到 1:前缀匹配 2:完全匹配
func (this *Trie) find(word string) int {
	cur := this.root
	for _, ch := range word {
		ch -= 'a'
		if cur.son[ch] == nil {
			return 0 //找不到
		}

		cur = cur.son[ch]
	}

	if !cur.end {
		return 1 //不是最后一个字符，前缀匹配
	}

	//完全匹配
	return 2
}

func (this *Trie) Search(word string) bool {
	return this.find(word) == 2
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix) != 0
}
