package routers

import (
	"example-golang-oauth2/controllers"
	"example-golang-oauth2/controllers/facebook"
	"example-golang-oauth2/controllers/google"
	"example-golang-oauth2/controllers/twitter"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/google/oauth2", &controllersGoogle.Oauth2Controller{})
	beego.Router("/google/callback", &controllersGoogle.CallbackController{})
	beego.Router("/twitter/oauth", &controllersTwitter.Oauth2Controller{})
	beego.Router("/twitter/callback", &controllersTwitter.CallbackController{})
	beego.Router("/facebook/oauth2", &controllersFacebook.Oauth2Controller{})
	beego.Router("/facebook/callback", &controllersFacebook.CallbackController{})

	beego.Router("/", &controllers.MainController{})
}
