package routers

import (
	"github.com/astaxie/beegae"
	"github.com/jackgris/angularjsAndTestWithGo/controllers"
)

func init() {
	beegae.Router("/", &controllers.MainController{})
}
