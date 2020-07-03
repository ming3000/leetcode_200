package trie

//确定数据结构，由于每个数据都能接26个字母，所以下面肯定是个切片，而不是pre,next这样的格式
//设置tag来标记完整单词，辨别属性
type Trie struct {
	next  [26]*Trie
	isEnd bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{next: [26]*Trie{}, isEnd: false}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = new(Trie)
		} // if>>
		this = this.next[v-'a']
	} // for>
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			return false
		} // if>>
		this = this.next[v-'a']
	} // for>
	return this.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return false
		} // if>>
		this = this.next[v-'a']
	} //for>
	return this.isEnd
}
