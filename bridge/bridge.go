package bridge

import "fmt"

//  Bridge 桥接模式：
//        将抽象部分与它的实现部分分离，使它们都可以独立地变化

type Phone struct {
	soft ISoftware
	name string
}

type ISoftware interface {

	Run()
}


func (p *Phone) setSoft(soft ISoftware)  {
	if p == nil {
		return
	}
	p.soft = soft
}

func (p *Phone) Run()  {
	if p == nil {
		return
	}
	fmt.Println(p.name)
	p.soft.Run()
}

type PhoneA struct {
	Phone
}

func NewPhoneA(name string) *PhoneA {
	return &PhoneA{Phone{name: name}}
}
type PhoneB struct {
	Phone
}

func NewPhoneB(name string) *PhoneB {
	return &PhoneB{Phone{name: name}}
}

type TSoftware struct {
	ISoftware
}

type Software struct {
	name string
}

type SoftwareA struct {
	Software
}

func (s *Software) Run() {
	if s == nil {
		return
	}
	fmt.Println(s.name)
}

type SoftwareB struct {
	Software
}