package visitor

import "testing"

func TestVisitor(t *testing.T) {
	object := ObjectStructure{}
	object.Attach(&ConcreteElementA{"A"})
	object.Attach(&ConcreteElementB{"B"})

	cva := ConcreteVisitorB{"vA"}
	cvb := ConcreteVisitorB{"vB"}

	object.Accept(&cva)
	object.Accept(&cvb)
}
