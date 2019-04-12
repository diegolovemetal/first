package builder

import "testing"

func TestBuilder(t *testing.T) {
	thin := Thin{}
	fat := Fat{}

	director := Director{&thin}
	director.CreateInfo()

	director = Director{&fat}
	director.CreateInfo()

}
