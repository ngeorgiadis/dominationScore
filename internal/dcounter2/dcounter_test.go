package dcounter2

import (
	"fmt"
	"testing"
)

func TestDCounter(t *testing.T) {

	dc2 := New("../../data/nodes_all.csv")

	i := 0
	for _, n := range dc2.NameIndex {
		fmt.Println(n)
		i++
	}

	fmt.Println(i)
}
