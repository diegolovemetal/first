//  Observer 观察者模式：
//        定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。
//		这个主题对象在状态发生改变时，会通知所有观察者对象，使它们能够自动更新自己。
package observer

import "fmt"

type Subject interface {
	Notify()
	State() int
	SetState(int)
	AddCallFunc(*update)
	RemoveCallFunc(*update)
}

type update func(int)

type SubjectA struct {
	state int
	call []*update
}

func NewSubjectA(s int) *SubjectA {
	return &SubjectA{s, []*update{}}
}

func (s *SubjectA) Notify() {
	if s == nil {
		return
	}
	for _, c := range s.call {
		(*c)(s.state)
	}
}

func (s *SubjectA) State() int {
	if s == nil {
		return 0
	}
	return  s.state
}

func (s *SubjectA) SetState(i int) {
	if s == nil {
		return
	}
	s.state = i
}

func (s *SubjectA) AddCallFunc(f *update) {
	if s == nil {
		return
	}
	for _, c := range s.call {
		if c == f {
			return
		}
	}
	s.call = append(s.call, f)
}

func (s *SubjectA) RemoveCallFunc(f *update) {
	if s == nil {
		return
	}
	for key, val := range s.call {
		if val == f {
			s.call = append(s.call[:key], s.call[key+1:]...)
		}
	}
}

type Observer interface {
	Update(int)
}

type ObserverA struct {
	s Subject
	state int
}

func (o *ObserverA) Update(i int)  {
	if o == nil {
		return
	}
	fmt.Println("ObserverA")
	fmt.Println(i)
	fmt.Println(o)
}

func NewObserverA(sa Subject, i int) *ObserverA {
	return &ObserverA{sa, i}
}

type ObserverB struct {
	s Subject
	state int
}

func (o *ObserverB) Update(i int)  {
	if o == nil {
		return
	}
	fmt.Println("ObserverB")
	fmt.Println(i)
	fmt.Println(o)
}

func NewObserverB(sb Subject, i int) *ObserverB  {
	return &ObserverB{sb, i}
}
