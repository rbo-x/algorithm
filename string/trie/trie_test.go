package trie

import "testing"


func TestPut(t *testing.T) {
    trie := Trie[int]{}
    trie.Put("she",10)
    trie.Put("sells",10)
    trie.Put("shells",10)
    trie.Put("by",10)
    trie.Put("the",20)
    trie.Put("sea",20)
    trie.Put("shore",20)
    t.Log(trie.Keys())
    trie.Delete("shore")
    t.Log(trie.Keys())
    trie.Delete("sea")
    t.Log(trie.Keys())
}
