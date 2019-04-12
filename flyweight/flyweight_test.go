package flyweight

import "testing"

func TestFlyweight(t *testing.T)  {
	ff := NewflyweightFactory()

	fya := ff.Flyweight("A")
	fya.Operation(1)
	fyb := ff.Flyweight("B")
	fyb.Operation(2)
	fyc := ff.Flyweight("C")
	fyc.Operation(3)
	fyd := ff.Flyweight("D")
	fyd.Operation(4)
	fyu := ff.Flyweight("u")
	fyu.Operation(5)

}
