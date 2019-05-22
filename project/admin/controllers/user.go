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

func (this *UserController) Insert() {
	user_name := this.Ctx.Input.Param(":user_name")
	values := this.Ctx.Input.Params()
	str := fmt.Sprintln(values)

	this.Ctx.WriteString("这是自动匹配路由 user/add, values are " + str + "user name is " + user_name)
}