package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

type Node struct {
	Idx int32
	Val int32
}

func (n1 *Node) Less(i interface{}) bool {
	n2 := i.(*Node)
	return n1.Val < n2.Val
}

func Test_sort(t *testing.T) {
	vec := SortIf{}
	rand.Seed(12345)
	for i := int32(0); i < 100; i++ {
		vec = append(vec, &Node{
			Idx: i,
			Val: rand.Int31() % 1000,
		})
	}
	for _, np := range vec {
		n := np.(*Node)
		fmt.Printf("before : idx =%v, val = %v\n", n.Idx, n.Val)
	}
	sort.Sort(vec)
	for _, np := range vec {
		n := np.(*Node)
		fmt.Printf("after : idx =%v, val = %v\n", n.Idx, n.Val)
	}
}
