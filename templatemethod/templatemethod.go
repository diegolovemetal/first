package templatemethod

import "fmt"

// Template Methed模板方法：
//        定义一个操作中的算法的骨架，而将一些具体步骤延迟到子类中。
//		模板方法使得子类可以不改变一个算法的结构即可重定义该算法的某些特定步骤。

type getFood interface {
	first()
	second()
	third()
}

type template struct {
	g getFood
}

func (b *template) getsomefood() { //b 只能为一个类而不能为皆可
	if b == nil{
		return
	}
	b.g.first()
	b.g.second()
	b.g.third()
}

type Refrigerator struct {
	template
}

func NewRefrigerator() *Refrigerator {
	return &Refrigerator{}
}

func (r *Refrigerator) first() {
	if r == nil {
		return
	}
	fmt.Println("open the refrigerator")
}

func (r *Refrigerator) second() {
	if r == nil {
		return
	}
	fmt.Println("get pie")
}

func (r *Refrigerator) third()  {
	if r == nil {
		return
	}
	fmt.Println("close the refrigerator")
}

type Pot struct {
	template
}

func NewPot() *Pot {
	return &Pot{}
}

func (p *Pot) first() {
	if p == nil {
		return
	}
	fmt.Println("open it")
}

func (p *Pot) second() {
	if p == nil {
		return
	}
	fmt.Println("cooking")
}

func (p *Pot) third() {
	if p == nil {
		return
	}
	fmt.Println("cleaning")
}