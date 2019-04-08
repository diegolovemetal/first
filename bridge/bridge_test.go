package bridge

import (
	"fmt"
	"testing"
)

func TestBridge(t *testing.T) {
	sa := SoftwareA{Software{"softwareA"}}
	sb := SoftwareB{Software{"softwareB"}}

	pa := NewPhoneA("phoneA")
	pb := NewPhoneB("phoneB")

	pa.setSoft(&sa)
	pa.Run()

	pb.setSoft(&sb)
	pb.Run()

	fmt.Println()
	p := TSoftware{&sb}
	p.Run()
}