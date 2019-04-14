package observer

import "testing"

func TestObserver(t *testing.T)  {
	var s = NewSubjectA(12)
	var oa = NewObserverA(s, 1)
	var ob = NewObserverB(s, 1)
	var oafunc update = oa.Update
	s.AddCallFunc(&oafunc)

	var obfunc update = ob.Update
	s.AddCallFunc(&obfunc)
	s.Notify()

	s.SetState(3)
	s.Notify()
}
