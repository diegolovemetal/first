package command

import "testing"

func TestCommand(t *testing.T)  {
	invoker := NewInvoker()
	conComA := NewConcreteCommandA()
	receA := NewReceiverA()

	conComA.SetReceiver(*receA)
	invoker.AddCommand(conComA)

	conComB := NewConcreteCommandB()
	receB := NewReceiverB()

	conComB.SetReceiver(*receB)
	invoker.AddCommand(conComB)

	invoker.ExecuteCommand()

}
