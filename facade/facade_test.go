package facade

import "testing"

func TestFacade(t *testing.T)  {
	f := NewFacade(21, 2880.5, "diegolovemetal")
	f.OutOne()
	f.OutTwo()
}
