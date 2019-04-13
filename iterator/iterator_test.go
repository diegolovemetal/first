package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T)  {
	bg := BookGroup{}
	nb1 := Book{"python"}
	bg.add(nb1)
	nb2 := Book{"golang"}
	bg.add(nb2)

	it := bg.createIterator()

	for b := it.first(); b != nil; b = it.next() {
		fmt.Println(b)
	}
}
