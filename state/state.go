package state

import "fmt"

//当一个对象的内在状态改变时，允许改变其行为，这个对象看起来像是改变了其类
type Work struct {
	hour int
	state State
}

type State interface {
	writeProgram(w *Work)
}

func (w *Work) State() State {
	if w == nil {
		return nil
	}
	return  w.state
}

func (w *Work) Hour() int {
	if w == nil {
		return -1
	}
	return w.hour
}

func (w *Work) SetState(s State) {
	if w == nil {
		return
	}
	w.state = s
}

func (w *Work) SetHour(h int)  {
	if w == nil {
		return
	}
	w.hour = h
}

func (w *Work) writeProgram()  {
	if w == nil {
		return
	}
	w.state.writeProgram(w)
}
//上午时分
type morningState struct {
	State
}
func NewWork() *Work{
	state := new(morningState)
	return &Work{state:state}
}

func (m *morningState) writeProgram(w *Work) {
	if w.Hour() < 12 {
		fmt.Println("现在是上午：", w.Hour())
	} else {
		w.SetState(new(NoonState))
		w.writeProgram()
	}
}
//中午时分
type NoonState struct {
	State
}

func (n *NoonState) writeProgram(w *Work)  {
	if w.Hour() < 13 {
		fmt.Println("现在是中午", w.Hour())
	} else {
		w.SetState(new(AfternoonState))
		w.writeProgram()
	}
}

type AfternoonState struct {
	State
}

func (a *AfternoonState) writeProgram(w *Work) {
	if w.Hour() <17 {
		fmt.Println("现在是下午", w.Hour())
	} else {
		w.SetState(new(EveningState))
		w.writeProgram()
	}
}

type EveningState struct {
	State
}

func (e *EveningState) writeProgram(w *Work) {
	if w.Hour() < 21 {
		fmt.Println("现在是晚上", w.Hour())
	} else {
		fmt.Println("现在开始睡觉", w.Hour())
	}
}