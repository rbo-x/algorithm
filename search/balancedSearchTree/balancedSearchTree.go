package balancedsearchtree

import "fmt"

type node[K,V any] struct {
    color bool
    key K 
    value V 
    left *node[K,V]
    right *node[K,V]
    size int
}

type Bst[K,V any] struct {
    root *node[K,V]
    compare func(K,K) int
}

func newNode[K,V any](key K,value V, size int) *node[K,V] {
    return &node[K,V]{
        key : key,
        value: value,
        size : size,
        color : true,
    }
}

func NewBst[K,V any](compare func(K,K)int) Bst[K,V] {
    return Bst[K,V]{
        nil,
        compare,
    }
}


func size[K,V any](n *node[K,V]) int{
    if n == nil {
        return 0
    }
    return n.size
}

func heigth[K,V any](n *node[K,V]) int{
    if n == nil {
        return 0
    }
    righH := heigth(n.right)
    leftH:= heigth(n.left)
    if righH > leftH {
        return righH + 1
    }
    return leftH + 1
}

func (bst Bst[K,V]) Heigth() int{
    return heigth(bst.root)
}

func (bst Bst[K,V]) Size() int{
    return size(bst.root)
}

func (bst Bst[K,V]) get(n *node[K,V],key K) *node[K,V]  {
    if n == nil {
        return nil
    }
    cmp := bst.compare(n.key,key)
    if cmp > 0 {
        return bst.get(n.left,key)     
    }else if cmp < 0 {
        return bst.get(n.right,key) 
    }
    return n
}

func (bst Bst[K,V]) Get(key K) (V,error)  {
    var defualt V
    v := bst.get(bst.root,key)
    if v == nil {
        return defualt,fmt.Errorf("not found")
    }
    return v.value,nil
}

func (bst Bst[K,V]) isRed(n *node[K,V]) bool {
    if n == nil {
        return false
    }
    return n.color
}

func (bst Bst[K,V]) flipColors(n *node[K,V]){
    n.left.color = false
    n.right.color = false
    n.color = true
}

func (bst Bst[K,V]) rotateLeft(n *node[K,V]) *node[K,V]{
    x := n.right
    n.right = x.left
    x.left = n
    x.color = n.color
    n.color = true
    x.size = n.size 
    n.size = 1 + size(n.left) + size(n.right)
    return x
}

func (bst Bst[K,V]) rotateRight(n *node[K,V])*node[K,V]{
    x := n.left
    n.left= x.right
    x.right= n
    x.color = n.color
    n.color = true
    x.size = n.size 
    n.size = 1 + size(n.left) + size(n.right)
    return x
}

func (bst Bst[K,V]) put(n *node[K,V],key K, value V) *node[K,V] {
    if n == nil {
        return newNode(key,value,1)
    } 
    cmp := bst.compare(n.key,key)
    if cmp > 0 {
        n.left = bst.put(n.left,key,value)     
    }else if cmp < 0 {
        n.right = bst.put(n.right,key,value) 
    }else {
        n.value = value
    }

    if bst.isRed(n.right) && !bst.isRed(n.left) {
        n = bst.rotateLeft(n)
    }
    if bst.isRed(n.right) && bst.isRed(n.left) {
        bst.flipColors(n)
    }
    if bst.isRed(n.left) && bst.isRed(n.left.left) {
        n = bst.rotateRight(n)
    }
    n.size = size(n.right)+size(n.left) + 1
    return n
}

func (bst *Bst[K,V]) Put(key K,value V)   {
    bst.root = bst.put(bst.root,key,value)
}

func (bst Bst[K,V]) min(n *node[K,V]) *node[K,V]{
    if n.left == nil {
        return n
    }
    return bst.min(n.left)
}

func (bst Bst[K,V]) Min() K  {
    return bst.min(bst.root).key
}

func (bst Bst[K,V]) max(n *node[K,V]) *node[K,V]{
    if n.right == nil {
        return n
    }
    return bst.max(n.right)
}

func (bst Bst[K,V]) Max() K {
    return bst.max(bst.root).key
}

func (bst Bst[K,V]) floor(n *node[K,V],key K) *node[K,V]  {
    if n == nil {
        return nil
    }
    cmp := bst.compare(n.key,key)
    if cmp > 0 {
        return bst.floor(n.left,key) 
    }
    if cmp == 0 {
        return n
    }
    t := bst.floor(n.right,key)
    if t != nil {
        return t
    }
    return n
}

func (bst Bst[K,V]) Floor(key K) K  {
    return bst.floor(bst.root,key).key
}

func (bst Bst[K,V]) ceil(n *node[K,V],key K) *node[K,V]  {
    if n == nil {
        return nil
    }
    cmp := bst.compare(n.key,key)
    if cmp < 0 {
        return bst.ceil(n.right,key) 
    }
    if cmp == 0 {
        return n
    }
    t := bst.ceil(n.left,key)
    if t != nil {
        return t
    }
    return n
}

func (bst Bst[K,V]) Ceil(key K) K  {
    return bst.ceil(bst.root,key).key
}

func (bst Bst[K,V]) sel(n *node[K,V],rank int) *node[K,V]{
    if n == nil {
        return nil 
    }
    s := size(n.left)
    if s == rank {
        return n
    }else if s > rank {
        return bst.sel(n.left,rank)
    }else {
        return bst.sel(n.right,rank-s-1)
    }
}

func (bst Bst[K,V]) Select(rank int) K  {
    return bst.sel(bst.root,rank).key
}

func (bst Bst[K,V]) rank(n *node[K,V],key K) int {
    if n == nil {
        return 0
    } 
    cmp := bst.compare(n.key,key)
    if cmp > 0 {
        return bst.rank(n.left,key) 
    }else if cmp < 0 {
        return 1+size(n.left) + bst.rank(n.right,key)
    }
    return  size(n.left)
}

func (bst Bst[K,V]) Rank(key K) int  {
    return bst.rank(bst.root,key) 
}


func (bst Bst[K,V]) deleteMin(n *node[K,V])  *node[K,V]{
    if n.left == nil {
        return n.right
    } 
    n.left = bst.deleteMin(n.left)
    return n
}

func (bst Bst[K,V]) DeleteMin()  {
    bst.root = bst.deleteMin(bst.root) 
}

func (bst Bst[K,V]) deleteMax(n *node[K,V])  *node[K,V]{
    if n.right == nil {
        return n.left
    } 
    n.right= bst.deleteMax(n.right)
    return n
}

func (bst Bst[K,V]) DeleteMax()  {
    bst.root = bst.deleteMax(bst.root) 
}

func (bst Bst[K,V]) del(n *node[K,V],key K) *node[K,V]{
    if n == nil {
        return nil
    }
    cmp := bst.compare(key,n.key)
    if cmp > 0 {
        n.right = bst.del(n.right,key)
    }else if cmp < 0 {
        n.left = bst.del(n.left ,key)
    }else {
        if n.right == nil {
            return n.left
        }
        if n.left == nil {
            return n.right
        }
        x := n
        n = bst.min(x.right)
        n.right = bst.deleteMin(x.right)
        n.left = x.left
    }
    n.size = size(n.left) + size(n.right) + 1 
    return n
}

func (bst Bst[K,V]) Delete(key K) {
    bst.root = bst.del(bst.root,key)
}

