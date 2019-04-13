// Mediator 中介者模式：
//        用一个中介对象来封装一系列的对象交互。
//        中介这使各对象不需要显式地相互引用，从而使其耦合松散，
//        而且可以独立地改变它们之间的交互。
package mediator

import "fmt"

// 中介者接口

type IMediator interface {
	Send(string, IColleague)
}

// 合作者接口

type IColleague interface{
	Send(string)
	Notify(string)
}

// 实现中介者接口的基本类型
type Mediator struct {

}
//具体的中介者

type ConcreteMediator struct {
	Mediator
	colleagues []IColleague
}

func (m *ConcreteMediator) AddColleague(c IColleague) {
	if m == nil {
		return
	}
	m.colleagues = append(m.colleagues, c)
}

func (m *ConcreteMediator) Send(message string, c IColleague) {
	if m == nil {
		return
	}
	for _, val := range m.colleagues {
		if c == val {
			continue
		}
		val.Notify(message)
	}
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{}
}

// 实现合作者接口的基本类型

type Colleague struct {
	mediator IMediator
}

//具体合作者对象A
type ConcreteColleagueA struct {
	Colleague
}

func (c *ConcreteColleagueA) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("ConcreteColleagueA get message:", message)
}

func (c *ConcreteColleagueA) Send(message string)  {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

func NewConcreteColleagueA(mediator IMediator) *ConcreteColleagueA {
	return &ConcreteColleagueA{Colleague{mediator}}
}

//具体合作对象B

type ConcreteColleagueB struct {
	Colleague
}

func (c *ConcreteColleagueB) Notify(message string) {
	if c == nil {
		return
	}
	fmt.Println("ConcreteColleagueB get message", message)
}

func (c *ConcreteColleagueB) Send(message string) {
	if c == nil {
		return
	}
	c.mediator.Send(message, c)
}

func NewConcreteColleagueB(mediator IMediator) *ConcreteColleagueB {
	return &ConcreteColleagueB{Colleague{mediator}}
}