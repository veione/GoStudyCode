package main

type Trie struct {
	IsEnd    bool      // 是否是叶子节点
	Children [26]*Trie // 孩子节点
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, i := range word {
		ch := i - 'a'
		if node.Children[ch] == nil {
			node.Children[ch] = &Trie{}
		}
		node = node.Children[ch]
	}
	node.IsEnd = true
}

func (this *Trie) SearchPrefix(prefix string)*Trie {
	node := this
	for _, i := range prefix {
		ch := i - 'a'
		if node.Children[ch] == nil {
			return nil
		}
		node = node.Children[ch]
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.SearchPrefix(word)
	return node != nil && node.IsEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.SearchPrefix(prefix)
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
