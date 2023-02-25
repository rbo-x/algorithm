package hashmap


type Hasher interface {
    Hash() int
    Equal(another Hasher) bool
}

type entry[K Hasher, V any] struct {
    key K
    value V
    hasValue bool
}

type HashMap[K Hasher ,V any] struct {
    entries []entry[K,V]
    size int
}

func NewMap[K Hasher,V any]() HashMap[K,V] {
    return HashMap[K,V]{
        entries: make([]entry[K,V], 8),
    } 
}

func findEntry[K Hasher,V any](entries []entry[K,V],key K) int {
    idx := key.Hash() % len(entries) 
    for {
        e := entries[idx] 
        if  !e.hasValue || key.Equal(e.key) {
            return idx
        }
        idx = (idx+1)%len(entries)
    }
}

func (hm *HashMap[K, V]) Set(key K,value V) {
    if ( hm.size >= len(hm.entries)) {
        hm.adjust(len(hm.entries)*2)    
    }
    e := entry[K,V]{key,value,true}
    idx := findEntry(hm.entries,e.key)
    if !hm.entries[idx].hasValue {
        hm.size++
    }
    hm.entries[idx] = e
}

func (hm *HashMap[K, V]) Get(key K) (V,bool) {
    var result V
    idx := findEntry(hm.entries,key)
    if hm.entries[idx].hasValue {
        return hm.entries[idx].value,true
    }
    return result,false
}

func (hm *HashMap[K, V]) Delete(key K) bool {
    idx := findEntry(hm.entries,key)
    if hm.entries[idx].hasValue {
        hm.entries[idx] = entry[K, V]{}
        hm.size--
        return true
    }
    return false
}


func (hm *HashMap[K, V]) adjust(newSize int) {
    newEntries := make([]entry[K,V],newSize)
    for _,v := range hm.entries {
        if  v.hasValue {
            idx := findEntry(newEntries,v.key)
            newEntries[idx] = v
        }
    }
    hm.entries = newEntries
}
