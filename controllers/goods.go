package controllers

import (
	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {
	c.Data["title"] = "你好，商品"
	c.Data["num"] = 12
	c.TplName = "goods.tpl"
}

func (c *GoodsController) DoAdd() {
	c.Ctx.WriteString("执行增加操作")
}

func (c *GoodsController) DoEdit() {
	c.Ctx.WriteString("执行增加操作")
}
