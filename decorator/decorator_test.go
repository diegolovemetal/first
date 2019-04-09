package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T)  {
	person := &person{"diego"}
	person.show()
	fmt.Println("")
	ts := new(Tshirts)
	ts.SetDecorator(person)
	ts.show()
	fmt.Println()

	bt := new(BigTrouser)
	ts.SetDecorator(ts)
	bt.show()
	fmt.Println()

	sk := new(Sneaker)
	sk.SetDecorator(bt)
	sk.show()


}
