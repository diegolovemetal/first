/*
 Proxy 代理模式：
        为其他对象提供一种代理，以控制对这个对象的访问。
*/
package proxy

import "fmt"

type GiveGifts interface {
	giveDoll()
	giveFlower()
	giveChocolate()
}

type Girl struct {
	name string
}

func (g *Girl) Name() string {
	if g == nil {
		return ""
	}
	return g.name
}

func (g *Girl) SetName(name string) {
	if g == nil {
		return
	}
	g.name = name
}

type Pursuit struct {
	girl Girl
}

func (p *Pursuit) giveDoll() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你一个洋娃娃")
}

func (p *Pursuit) giveFlower() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你一束鲜花")
}

func (p *Pursuit) giveChocolate() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你一块巧克力")
}

type Proxy struct {
	p Pursuit
}

func (p *Proxy) giveDoll() {
	if p == nil{
		return
	}
	p.p.giveDoll()
}

func (p *Proxy) giveFlower() {
	if p == nil {
		return
	}
	p.p.giveFlower()
}

func (p *Proxy) giveChocolate() {
	if p == nil {
		return
	}
	p.p.giveChocolate()
}

func NewProxy(mm Girl) *Proxy {
	gg := Pursuit{mm}
	return &Proxy{gg}
}