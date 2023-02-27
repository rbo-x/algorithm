package trie


type tstNode[T any] struct {
    left,mid,right *tstNode[T] 
    c byte
    value T
    hasValue bool
}

type Tst[T any] struct {
    root *tstNode[T]
}


func (tst *Tst[T]) put(n *tstNode[T],key string,value T,d int) *tstNode[T]{
    x := key[d] 
    if n == nil {
        n = &tstNode[T]{c:x}
    }
    if x < n.c {
        n.left = tst.put(n.left,key,value,d) 
    }else if x > n.c {
        n.right = tst.put(n.right,key,value,d) 
    }else if d < len(key) - 1 {
        n.mid = tst.put(n.mid,key,value,d+1) 
    }else {
        n.value = value 
        n.hasValue = true
    }
    return n
}

func (tst *Tst[T]) Put(key string,value T){ 
    tst.root = tst.put(tst.root,key,value,0)
}

func (tst *Tst[T]) get(n *tstNode[T],key string,d int) *tstNode[T]{
    if n == nil {
        return nil
    }
    x := key[d] 
    if x < n.c {
        return tst.get(n.left,key,d)
    }else if x > n.c {
        return tst.get(n.right,key,d)
    }else if d < len(key) - 1 {
        return tst.get(n.mid,key,d+1)
    }
    return n
}

func (tst *Tst[T]) Get(key string) (T,bool) { 
    var d T
    result := tst.get(tst.root,key,0)
    if result == nil {
        return d,false
    }
    if result.hasValue {
        return result.value,true
    }
    return d,false
}
