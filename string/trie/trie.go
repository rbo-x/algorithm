package trie

import "fmt"

type node[T any] struct {
    value T 
    hasValue bool
    next [255]*node[T]
}

type Trie[T any] struct {
    root *node[T] 
}

func (t *Trie[T]) get(n *node[T],key string,d int) *node[T]{
    if n == nil {
        return nil
    }
    if d == len(key) {
        return n
    }
    x := key[d]
    return t.get(n.next[x],key,d+1);
}

func (t *Trie[T]) Get(key string) (T,bool){
    var result T
    n := t.get(t.root,key,0)
    if n == nil {
        return result,false 
    }
    return n.value,true
}

func (t *Trie[T]) put(n *node[T],key string,value T,d int) *node[T]{
    if n == nil {
        n = &node[T]{}
    }
    if d == len(key) {
        n.value = value
        n.hasValue = true
        return n
    }
    x := key[d]
    n.next[x] = t.put(n.next[x],key,value,d+1)
    return n
}

func (t *Trie[T]) Put(key string,value T) {
    t.root = t.put(t.root,key,value,0)
}

func (t *Trie[T]) Keys() []string {
    queue := []string{}
    t.collect(t.root,"",&queue)
    return queue
}

func (t *Trie[T]) KeysWithPrefix(prefix string) []string {
    queue := []string{}
    t.collect(t.get(t.root,prefix,0),prefix,&queue)
    return queue
}

func (t *Trie[T]) collect(n *node[T],pre string,queue *[]string) {
    if n == nil {
        return
    }
    if n.hasValue {
        *queue = append(*queue, pre) 
    }
    for i,c := range n.next {
        t.collect(c,pre+fmt.Sprintf("%c",i),queue)
    }
}

func (t *Trie[T]) KeysThatMatch(pattern string) []string {
    queue := []string{}
    t.collectWithPattern(t.root,"",pattern,&queue)
    return queue
}

func (t *Trie[T]) collectWithPattern(n *node[T],pre string,pattern string,queue *[]string) {
    if n == nil {
        return
    }
    if n.hasValue && len(pattern) == len(pre) {
        *queue = append(*queue, pre) 
    }
    if len(pattern) == len(pre) {
        return
    }
    next := pattern[len(pre)]
    for i,c := range n.next {
        if i == int(next) || next == '.' {
            t.collectWithPattern(c,pre+fmt.Sprintf("%c",i),pattern,queue)
        }
    }
}

func (t *Trie[T]) search(n *node[T],s string,length int,d int) int{
    if n == nil {
        return length
    } 
    if n.hasValue {
        length = d
    }
    if d == len(s) {
        return length
    }
    x := s[d]
    return t.search(n.next[x],s,length,d+1)
}

func (t *Trie[T]) LongestPrefixOf(s string) string {
    idx := t.search(t.root,s,0,0)
    return s[0:idx]
}

func (t *Trie[T]) delete(n *node[T],key string, d int) *node[T]{
    if n == nil {
        return nil
    }
    if len(key) == d {
        n.hasValue = false
    }else{
        x:=key[d]
        n.next[x] = t.delete(n.next[x],key,d+1)
    }

    if n.hasValue {
        return n
    }

    for _,c := range n.next {
        if c != nil {
            return n
        }
    }
    return nil
}

func (t *Trie[T]) Delete(key string) {
    t.root = t.delete(t.root,key,0)
}


func (t *Trie[T]) print(n *node[T],char string) {
    fmt.Printf("%s : [ ",string(char))  
    for i,c := range n.next {
        if c != nil {
            fmt.Printf("%s ",string(byte(i)))
        }
    }
    fmt.Printf(" ]\n")  
    for i,c := range n.next {
        if c != nil {
            t.print(c,string(byte(i)))
        }
    }
}

func (t *Trie[T]) Print() {
    t.print(t.root,"root")
}
