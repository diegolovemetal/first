package main

import "fmt"

const(
	White = iota
	Black
	Blue
	Red
	Yellow
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}
//a list of Box
type BoxList []Box

func (b Box) Volume() float64 {
	return b.width * b.depth *b.height
}

func (b *Box) SetColor(c Color) {
	b.color = c
}

func (bl BoxList) BiggestColor() Color {
	v := 0.00
	k := Color(White)

	for _, b := range bl {
		if bv := b.Volume(); bv > v {
			v = bv
			k = b.color
		}
	}
	return k
}

func (bl BoxList) PainItBlack() {
	for i := range bl {
		bl[i].SetColor(Black)
	}
}

func (c Color) String() string {
	strings := []string {"White", "Black", "Blue", "Red", "Yellow"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4,4,4,	Red},
		Box{1,2,3, Yellow},
		Box{2,3,4, Blue},
	}

	fmt.Printf("Boxes len:%d\n", len(boxes))
	fmt.Println("the volume of the first one is ", boxes[0].Volume(), "cm³")
	fmt.Println("the color of the last one is ", boxes[len(boxes)-1].color.String())
	fmt.Println("the biggest one is ", boxes.BiggestColor().String())

	fmt.Println("lets paint it black")
	boxes.PainItBlack()
	fmt.Println("the color of the second one is ",boxes[1].color.String())
	fmt.Println("Obviously, Now the biggest one is ", boxes.BiggestColor().String())

}
