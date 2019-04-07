package adapter

import (
	"testing"
)
//func TestAdapter(t *testing.T)  {
//	pF := NewForwards("F")
//	pC := NewCenter("C")
//	pFC := NewTranslator("FC")
//
//	pF.attack()
//	pC.defense()
//
//	pFC.attack()
//	pFC.defense()
//}

func TestAdapter(t *testing.T) {
	pF := NewForwards("F")
	pC := NewCenter("C")
	pFC := NewTranslator("FC")

	pF.attack()
	pC.defense()
	pFC.attack()
	pFC.defense()
}