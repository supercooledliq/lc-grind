type trieNode struct {
	child map[rune]*trieNode
	isEnd bool
}

type WordDictionary struct {
	root *trieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &trieNode{child: make(map[rune]*trieNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root

	for _, ch := range word {
		if _, ok := node.child[ch]; !ok {
			node.child[ch] = &trieNode{child: make(map[rune]*trieNode)}
		}
		node = node.child[ch]
	}

	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, pos int, node *trieNode) bool {
	if len(word) == 0 || pos > len(word)-1 {
		return node.isEnd
	}

	// Two case:
	// Case1: If the current character is "." then we need to dfs over all the child nodes
	// Case2: If the current character is not "." then we simply continue
	currentChar := rune(word[pos])
	switch {
	case currentChar == '.':
		// dfs over all the children of this node
		for _, next := range node.child {
			if this.dfs(word, pos+1, next) {
				return true
			}
		}
		return false
	default:
		_, ok := node.child[currentChar]
		if !ok {
			return false
		}
		return this.dfs(word, pos+1, node.child[currentChar])
	}

	return true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
