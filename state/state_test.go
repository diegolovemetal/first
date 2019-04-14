package state

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T)  {
	w := NewWork()

	fmt.Println("===============")
	w.writeProgram()
	fmt.Println("===============")
	w.SetHour(11)
	w.writeProgram()
	fmt.Println("===============")
	w.SetHour(12)
	w.writeProgram()
	fmt.Println("===============")
	w.SetHour(13)
	w.writeProgram()
	fmt.Println("===============")
	w.SetHour(18)
	w.writeProgram()
	fmt.Println("===============")
	w.SetHour(22)
	w.writeProgram()



}
