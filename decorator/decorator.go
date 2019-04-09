//Decorator 装饰模式：
//        动态地给一个对象添加一些额外的职责，就增加功能来说，装饰模式比生成子类更为灵活。
package decorator

import "fmt"

type person struct {
	Name string
}

func (p *person) show() {
	if p == nil {
		return
	}
	fmt.Println("姓名是：", p.Name)
}

type AbstractPerson interface {
	show()
}

type Decorator struct {
	AbstractPerson
}
func (d *Decorator) SetDecorator(component AbstractPerson)   {
	if  d == nil {
		return
	}
	d.AbstractPerson = component
}

func (d *Decorator) show() {
	if d == nil {
		return
	}
	if d.AbstractPerson != nil {
		d.AbstractPerson.show()
	}
}

type Tshirts struct {
	Decorator
}

func (t *Tshirts) show() {
	if t == nil {
		return
	}
	t.Decorator.show()
	fmt.Println("T-shirt!")
}

type BigTrouser struct {
	Decorator
}

func (b *BigTrouser) show() {
	if b == nil {
		return
	}
	b.Decorator.show()
	fmt.Println("trouser")
}

type Sneaker struct {
	Decorator
}

func (s *Sneaker) show() {
	if s == nil {
		return
	}
	s.Decorator.show()
	fmt.Println("Sneaker")
}
