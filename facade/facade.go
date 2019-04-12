package facade

import "fmt"

// Facade 外观模式：
//        为子系统中的一组接口提供一个一致的界面，此模式定义了一个高层接口，
//		这个接口使得这一子系统更加容易使用（投资：基金，股票，房产）

type facade struct {
	one	funcone
	two functwo
	three functhree
}

func (f facade) OutOne()  {
	f.one.Out()
	f.three.Out()
}

func (f facade) OutTwo() {
	f.two.Out()
	f.three.Out()
}

type funcone struct {
	str string
}

func (f funcone) Out() {
	fmt.Println("funcone", f.str)
}

type functwo struct {
	i int
}

func (f functwo) Out() {
	fmt.Println("functwo", f.i)
}

type functhree struct {
	f float32
}

func (f functhree) Out() {
	fmt.Println("functhree", f.f)
}

func NewFacade(i int, f float32, str string) *facade {
	return &facade{funcone{str}, functwo{i}, functhree{f}}
}