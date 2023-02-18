package bst

import (
	"math/rand"
	"strings"
	"testing"
)



func TestBstPut(t *testing.T) {
    testCases := []struct {
        desc    string
        bst Bst[string,int]
        keys []string
        h int
    }{
        {
            desc: "put method",
            bst: NewBst[string,int](strings.Compare), 
            keys: []string{"A","B","C","D","E","F","G","H"},
            h : 8,
        },
    }
    for _, tC := range testCases {
        t.Run(tC.desc, func(t *testing.T) {
            for _, k := range tC.keys {
                tC.bst.Put(k,rand.Intn(100))
            }        
            if tC.bst.Size() != len(tC.keys) {
                t.Errorf("size: expected %d got %d\n",len(tC.keys),tC.bst.Size())
            }
            if tC.bst.Heigth() != tC.h {
                t.Errorf("heigth : expected %d got %d\n",tC.h,tC.bst.Heigth())
            }
        })
    }
}

func TestBstMinAndMAx(t *testing.T) {
    bst := NewBst[int,string](func(i1, i2 int) int {return i1-i2})
    bst.Put(1,"tokyo")
    bst.Put(0,"london")
    bst.Put(102,"madrid")
    bst.Put(77,"kyiv")
    bst.Put(45,"berlin")
    bst.Put(3,"paris")
    bst.Put(-1,"washington dc")

    max := bst.Max()
    min := bst.Min()
    if max != 102 {
        t.Errorf("max: expected %d got %d\n",102,max)
    }
    if min != -1 {
        t.Errorf("min: expected %d got %d\n",-1,min)
    }

    key := 46
    ceil := bst.Ceil(key)
    floor := bst.Floor(key)
    if ceil != 77 {
        t.Errorf("ceil : expected %d got %d\n",77,ceil)
    }
    if min != -1 {
        t.Errorf("floor: expected %d got %d\n",45,floor)
    }

    rank := bst.Rank(77)
    if rank != 5 {
        t.Errorf("rank : expected %d got %d\n",5,rank)
    }

    sel := bst.Select(5)
    if sel !=  77 {
        t.Errorf("rank : expected %d got %d\n",77,sel)
    }

    bst.Delete(77)
    _,err := bst.Get(77)
    if err == nil {
        t.Errorf("delete : expected %v got nil\n",err)
    }
}
