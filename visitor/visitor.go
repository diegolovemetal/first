package visitor

import "fmt"

// Visitor 访问者模式：
//        表示一个作用于某对象结构中的各元素的操作，
//		它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作

// 访问接口
type IVisitor interface {
	VisitorConcreteElementA(ConcreteElementA)
	VisitorConcreteElementB(ConcreteElementB)
}
// 具体访问者A
type ConcreteVisitorA struct {
	name string
}

type ConcreteElementA struct {
	name string
}
// 具体访问者B

type ConcreteVisitorB struct {
	name string
}
type ConcreteElementB struct {
	name string
}

func (c *ConcreteVisitorA) VisitorConcreteElementA(ce ConcreteElementA)  {
	if c == nil {
		return
	}

	fmt.Println(ce.name, c.name)
	ce.OperatorA()
}

func (c *ConcreteVisitorA) VisitorConcreteElementB(ce ConcreteElementB)  {
	if c == nil {
		return
	}

	fmt.Println(ce.name, c.name)
	ce.OperatorB()
}

func (c *ConcreteVisitorB) VisitorConcreteElementA(ce ConcreteElementA)  {
	if c == nil {
		return
	}

	fmt.Println(ce.name, c.name)
	ce.OperatorA()
}

func (c *ConcreteVisitorB) VisitorConcreteElementB(ce ConcreteElementB)  {
	if c == nil {
		return
	}

	fmt.Println(ce.name, c.name)
	ce.OperatorB()
}
//元素接口
type IElement interface {
	Accept(IVisitor)
}

func (c *ConcreteElementA) Accept(visitor IVisitor) {
	if c == nil {
		return
	}
	visitor.VisitorConcreteElementA(*c)
}

func (c *ConcreteElementA) OperatorA(){
	if c == nil {
		return
	}
	fmt.Println("OperatorA")
}

func (c *ConcreteElementB) OperatorB() {
	if c == nil {
		return
	}
	fmt.Println("OperatorB")
}

func (c *ConcreteElementB) Accept(visitor IVisitor) {
	if c == nil {
		return
	}
	visitor.VisitorConcreteElementB(*c)
}

//维护元素集合
type ObjectStructure struct {
	list []IElement
}

func (o *ObjectStructure) Attach(e IElement) {
	if o == nil && e == nil {
		return
	}
	o.list = append(o.list, e)
}

func (o *ObjectStructure) Delete(e IElement)  {
	if o == nil && e == nil {
		return
	}
	for i, val := range o.list {
		if val == e {
			o.list = append(o.list[:i], o.list[i+1:]...)
			break
		}
	}

}

func (o *ObjectStructure) Accept(v IVisitor) {
	if o == nil {
		return
	}
	for _, val := range o.list {
		val.Accept(v)
	}
}
