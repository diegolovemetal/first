package routers

import (
	"github.com/astaxie/beego"
	"project/admin/controllers"
)

func init() {
   	//固定路由
   	beego.Router("/admin", &admin.UserController{}, "*:Index")
   	beego.Router("/admin/add/:user_name/:user_pwd/:mobile", &admin.UserController{}, "get:Insert")
   	beego.Router("/admin/index", &admin.UserController{}, "get:Index")
}