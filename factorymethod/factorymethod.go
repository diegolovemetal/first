//Factory Method 工厂方法模式：
//定义一个用于创建对象的接口，让子类决定实例化哪一个类。
//工厂方法使一个类的实例化延迟到其子类
package factorymethod


type OperationFunc interface {
	SetNumA(float32)
	SetNumB(float32)
	Result() (float32, bool)
}

type Operation struct {
	numberA float32
	numberB float32
}

func (o *Operation) SetNumA(num float32) {
	if o == nil {
		return
	}
	o.numberA = num
}
func (o *Operation) SetNumB(num float32) {
	if o == nil {
		return
	}
	o.numberB = num
}

type OperationAdd struct {
	Operation
}

func (a *OperationAdd) Result() (float32, bool) {
	if a == nil {
		return 0, false
	}
	return a.numberA + a.numberB, true
}

type OperationSub struct {
	Operation
}

func (a *OperationSub) Result() (float32, bool) {
	if a == nil {
		return 0, false
	}
	return a.numberA - a.numberB, true
}

type CreateOperation interface {
	createoperation(string) OperationFunc
}

type COperationAdd struct {
}

func (c *COperationAdd) createoperation(op string) OperationFunc {
	if c == nil {
		return nil
	}
	return &OperationAdd{}
}

type COperationSub struct {
}

func (c *COperationSub) createoperation(op string) OperationFunc {
	if c == nil {
		return nil
	}
	return &OperationSub{}
}
