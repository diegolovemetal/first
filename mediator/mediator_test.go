package mediator

import "testing"

func TestMediator(t *testing.T)  {
	m := NewConcreteMediator()
	ca := NewConcreteColleagueA(m)
	cb := NewConcreteColleagueB(m)

	m.AddColleague(ca)
	m.AddColleague(cb)

	ca.Send("你好,我是A同事")
	cb.Send("你好,我是B同事")
}