package routers

import (
	admin "project/admin/controllers"
	"github.com/astaxie/beego"

)

func init() {
    beego.AutoRouter(&admin.UserController{})
}
