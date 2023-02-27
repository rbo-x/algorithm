package trie

import "testing"

func TestTst(t *testing.T) {
    tst := Tst[int]{}
    tst.Put("she",1)
    tst.Put("sells",2)
    tst.Put("shells",3)
    tst.Put("by",4)
    tst.Put("the",5)
    tst.Put("sea",6)
    tst.Put("shore",7)
    t.Log(tst.Get("she"))
    t.Log(tst.Get("shell"))
    t.Log(tst.Get("shells"))
    t.Log(tst.Get("shea"))
    t.Log(tst.Get("shore"))
}
