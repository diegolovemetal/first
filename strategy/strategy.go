package strategy

import "errors"

//Strategy 策略模式：
//        它定义了算法家族，分别封装起来，让它们可以相互替换，
//		此模式让算法的变化，不会影响到使用算法的客户。

type CashSuper interface {
	acceptCash(memory float32) float32
}

type CashNormal struct {

}
//正常收费
func (this *CashNormal) acceptCash(money float32) float32  {
	return money
}
//打折
type CashRebate struct {
	moneyRebate float32
}

func (this *CashRebate) acceptCash(money float32) float32  {
	return this.moneyRebate * money
}

//返现
type CashReturn struct {
	moneyCondition float32 //最低满减金额总值
	moneyReturn float32		//满减值
}

func (this *CashReturn) acceptCash(money float32)float32  {
	if money >= this.moneyCondition {
		return money - float32(this.moneyReturn)
	} else {
		return money
	}
}

type Context struct {
	CashSuper
}

func NewCashContext(str string)	(cash *Context, err error)  {
	cash = new(Context)
	switch str {
	case "normal":
		cash.CashSuper = &CashNormal{}
	case "rebate8":
		cash.CashSuper = &CashRebate{0.8}
	case "return100":
		cash.CashSuper = &CashReturn{300, 100}
	default:
		err = errors.New("no match")
	}
	return cash, err
}
