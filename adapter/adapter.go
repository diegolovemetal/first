//package adapter
//
//import "fmt"
//
//type Player interface {
//
//	attack()
//	defense()
//}
//
//type Forwards struct {
//	name string
//}
//
//func (f *Forwards) attack()  {
//	if f == nil {
//		return
//	}
//	fmt.Println(f.name,"attacking")
//}
//
//func (f *Forwards) defense() {
//	if f == nil {
//		return
//	}
//	fmt.Println(f.name, "defending")
//
//}
//
//func NewForwards(name string) Player  {
//	return &Forwards{name}
//}
//
//type Centers struct {
//	name string
//}
//
//func (f *Centers) attack() {
//	if f == nil {
//		return
//	}
//	fmt.Println(f.name, "在进攻")
//}
//func (f *Centers) defense() {
//	if f == nil {
//		return
//	}
//	fmt.Println(f.name, "在防守")
//}
//
//func NewCenter(name string) Player {
//	return &Centers{name}
//}
//
//type ForeignCenter struct {
//	name string
//}
//
//func (f *ForeignCenter) attack(what string) {
//	if f == nil {
//		return
//	}
//	fmt.Println(f.name, "在进攻")
//}
//func (f *ForeignCenter) defense() {
//	if f == nil {
//		return
//	}
//	fmt.Println(f.name, "在防守")
//}
//
////用户想要的接口
//
//type Translator struct {
//	f ForeignCenter
//}
//
//func (t *Translator) attack()  {
//	if t == nil {
//		return
//	}
//
//	t.f.attack("进攻")
//}
//func (t *Translator) defense()  {
//	if t == nil {
//		return
//	}
//
//	t.f.defense()
//}
//
//func NewTranslator(name string) Player {
//	return &Translator{ForeignCenter{name}}
//
package adapter

import (
	"fmt"
)

type Player interface {
	attack()
	defense()
}

type Forwards struct {
	name string
}

func (f *Forwards) attack() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}
func (f *Forwards) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

func NewForwards(name string) Player {
	return &Forwards{name}
}

type Centers struct {
	name string
}

func (f *Centers) attack() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}
func (f *Centers) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

func NewCenter(name string) Player {
	return &Centers{name}
}

type ForeignCenter struct {
	name string
}

func (f *ForeignCenter) attack(what string) {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在进攻")
}
func (f *ForeignCenter) defense() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "在防守")
}

type Translator struct {
	f ForeignCenter
}

// 这是用户想要的接口
func (t *Translator) attack() {
	if t == nil {
		return
	}
	t.f.attack("进攻")
}
func (t *Translator) defense() {
	if t == nil {
		return
	}
	t.f.defense()
}

func NewTranslator(name string) Player {
	return &Translator{ForeignCenter{name}}
}