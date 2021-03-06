package trie208

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
}

func generateNode() *Node {
	return &Node{
		next: make(map[string]*Node),
	}
}

/** Initialize your data structure here. */
func New() Trie {
	return Trie{
		root: generateNode(),
	}
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	cur := t.root

	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			cur.next[c] = generateNode()
		}
		cur = cur.next[c]
	}

	cur.isWord = true
}

/** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	cur := t.root

	for _, w := range []rune(word) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}

	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root

	for _, w := range []rune(prefix) {
		c := string(w)
		if cur.next[c] == nil {
			return false
		}
		cur = cur.next[c]
	}

	return true
}
