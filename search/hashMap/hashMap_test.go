package hashmap

import (
	"math/rand"
	"testing"
)

type mystring string

func (s mystring) Hash() int{
    result := 0
    for _,c := range s {
        result += int(c)
    }
    return result
}

func (s mystring) Equal(another Hasher) bool {
    v,ok := another.(mystring)
    if !ok {
        return false
    }
    return v == s
}

func testSetGetMethod(t *testing.T)  {
    m := NewMap[mystring,int]()
    keys := []mystring{"Tehran","LA","London","kyiv","Madrid","Warsaw","Washington"}
    for _, k := range keys {
        m.Set(k,rand.Intn(1000))
    }

    if m.size != len(keys) {
        t.Error("size of the map not equal with len of keys")
    }

    for _,k := range keys {
        _,ok := m.Get(k) 
        if !ok {
            t.Error("Get method doesn't return value")
        }
    }
}

func TestMain(t *testing.T) {
    testCases := []struct {
        desc    string
        testFunc func(t *testing.T) 
    }{
        {
            desc: "Testing Get and Set methods",
            testFunc: testSetGetMethod,
        },
    }
    for _, tC := range testCases {
        t.Run(tC.desc, func(t *testing.T) {
            tC.testFunc(t)
        })
    }
}
