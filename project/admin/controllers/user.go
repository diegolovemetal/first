package admin

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	this.Ctx.WriteString("自动匹配路由 user/index")
}

func (this *UserController) Test() {
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("自动匹配路由 user/test, values are " + str)
}
