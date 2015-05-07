package controllers

//AppConfigPath

import (
	hp "causal/helpers"
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	fmt.Println(hp.GetBuildNames())
	c.TplNames = "index.tpl"
}
