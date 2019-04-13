package memento

import "fmt"

type GameRole struct {
	vit int
	atk int
	def int
}

func (g *GameRole) GetInitState() {
	if g == nil {
		return
	}
	g.vit = 100
	g.atk = 100
	g.def = 100
}

func (g *GameRole) StateDisplay() {
	if g == nil {
		return
	}
	fmt.Println("当前角色状态：")
	fmt.Println("体力：", g.vit)
	fmt.Println("攻击：", g.atk)
	fmt.Println("防御：", g.def)
	fmt.Println("==============")
}

func (g *GameRole) SaveState() (r RoleStateMemento) {
	if g == nil {
		return RoleStateMemento{}
	}
	return RoleStateMemento{*g}
}

func (g *GameRole) RecoveryState(r RoleStateMemento)  {
	if g == nil {
		return
	}
	g.vit = r.vit
	g.atk = r.atk
	g.def = r.def
}
type RoleStateMemento struct {
	GameRole
}

//临时存放状态

type RoleStateCaretaker struct {
	memento RoleStateMemento
}

func (g *GameRole) Fight() {
	if g == nil {
		return
	}
	g.vit = 0
	g.atk = 0
	g.def = 0
}