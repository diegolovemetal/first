package proxy

import "testing"

func TestProxy(t *testing.T	)  {
	girl := Girl{}
	girl.SetName("Sasa")

	p := NewProxy(girl)
	p.giveChocolate()
	p.giveFlower()
	p.giveDoll()
}
