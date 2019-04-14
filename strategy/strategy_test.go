package strategy

import (
	"fmt"
	"testing"
)

func TestStrategy(t *testing.T)  {
	context, err := NewCashContext("return100")
	if err != nil {
		t.Error(err)
	} else {
		result := context.acceptCash(350)
		fmt.Println(result)
	}

	context, err = NewCashContext("normal")
	if err != nil {
		t.Error(err)
	} else {
		result := context.acceptCash(1000)
		fmt.Println(result)
	}

	context, err = NewCashContext("rebate8")
	if err != nil {
		t.Error(err)
	} else {
		result := context.acceptCash(300)
		fmt.Println(result)
	}
}
