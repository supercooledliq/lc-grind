type trieNode struct {
    // child nodes map[rune]*trieNode
    child map[rune]*trieNode
    // is this the end of the node
    isEnd bool
}

type Trie struct {
    root *trieNode
}


func Constructor() Trie {
    return Trie{root: &trieNode{child: make(map[rune]*trieNode)}}
}


func (t *Trie) Insert(word string)  {
    // Get the root node of the current trie 
    node := t.root

    // Iterate over the input word
    for _, ch := range word {
        // Check if a trie node exists against this character
        if _, exists := node.child[ch]; !exists {
            // The node for this character does not exists so create it
            // it also means that we might be seeing this character for the 
            // first time in the trie 
            node.child[ch] = &trieNode{child: make(map[rune]*trieNode)}
        //    if word == "abc" {
        //        fmt.Printf("Debug> Word: %s, NewChar: %+v", word, string(ch))
        //    }
        }
        // Now move to this newly created node
        // What is important here to understand is that each edge is represented
        // by a character in the input string 
        node = node.child[ch]
    }
    node.isEnd = true
}


func (t *Trie) Search(word string) bool {
    // Get the root node of the current trie 
    node := t.root
    
    // Iterate over the word 
    for _, ch := range word {
        // if node.isEnd {
        //     if word == "abc" {
        //         fmt.Println("returning on 51")
        //     }
        //     return false 
        // }

        // Lets say the word is "hello"
        if _, exists := node.child[ch]; !exists {
            // if word == "abc"{
            //     fmt.Println("returning on 58")
            // }
            return false 
        }       
        
        // Move to the next node now
        node = node.child[ch]
    }   

    return node.isEnd
    // if node.isEnd {
    //     return true
    // }
    // return false
}


func (t *Trie) StartsWith(prefix string) bool {
    node := t.root 

    // Lets say we have a word called "hello" already 
    // and we need to search for word "hell", then we need 
    // again start looking for node with character h, then e, and so on. 
    for _, ch := range prefix {
        if _, exists := node.child[ch]; !exists {
            return false 
        }
        node = node.child[ch]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
