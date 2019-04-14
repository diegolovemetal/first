package templatemethod

import "testing"

func TestTemplate(t *testing.T)  {
	b := NewRefrigerator()
	b.g = b
	b.getsomefood()

	p := NewPot()
	p.g = p
	p.getsomefood()
}
