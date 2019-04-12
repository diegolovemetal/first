//Builder 生成器模式：
//（建造者模式）将一个复杂对象的构建与它表示分离，使得同样的构建过程可以创建不同的表示
package builder

import "fmt"

type IBuilder interface {
	head()
	body()
	foot()
	hand()
}

type Thin struct {

}

func (t *Thin) head()  {
	fmt.Println("我的头很瘦")
}

func (t *Thin) body() {
	fmt.Println("我的身体很瘦")
}

func (t *Thin) foot() {
	fmt.Println("我的脚很瘦")
}

func (t *Thin) hand() {
	fmt.Println("我的手很瘦\n---------------------")
}

type Fat struct {

}

func (t *Fat) head() {
	fmt.Println("我的头很胖")
}

func (t *Fat) body() {
	fmt.Println("我的身体很胖")
}
func (t *Fat) foot() {
	fmt.Println("我的脚很胖")
}
func (t *Fat) hand() {
	fmt.Println("我的手很胖")
}

type Director struct {
	person IBuilder
}

func (d *Director) CreateInfo()  {
	if d == nil {
		return
	}
	d.person.head()
	d.person.body()
	d.person.foot()
	d.person.hand()
}