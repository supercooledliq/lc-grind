type trieNode struct {
    // Each character in the string holds a link to another trie node
    child map[rune]*trieNode
    isEnd bool
}

type trie struct {
    root *trieNode
}

func newTrie() *trie {
    return &trie{root: &trieNode{child: make(map[rune]*trieNode)}}
}

func (t *trie) insert(str string) {
    node := t.root
    // Iterate over the input string
    for _, ch := range str {
        // Check if the the current char has a trie node associated with it
        if _, exists := node.child[ch]; !exists {
            // If it doesn't exists, then create new newe node
            node.child[ch] = &trieNode{child: make(map[rune]*trieNode)}
        }
        // Iterate to the newly created node
        node = node.child[ch]
    }
    node.isEnd = true
}

func (t *trie) longestPrefix() string {
    node := t.root 
    prefix := []rune{}

    // Iterate over all child nodes to get the longest prefix
    for { 
        // If length is more than 1 then we know that we have a split 
        if len(node.child) != 1 || node.isEnd {
            break
        }

        // Iterate to a new node now
        for ch, nextNode := range node.child {
            prefix = append(prefix, ch)
            node = nextNode
        }
    }

    return string(prefix)
}

func longestCommonPrefix(strs []string) string {
    t := newTrie()
    
    for _, str := range strs {
        t.insert(str)
    }

    return t.longestPrefix()
}
