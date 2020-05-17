package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Ctx.WriteString("新闻列表")
}

func (c *ArticleController) AddArticle() {
	c.Ctx.WriteString("增加新闻")
}

// func (c *ArticleController) EditArticle() {
// 	c.Ctx.WriteString("修改新闻")
// }

func (c *ArticleController) EditArticle() {

	// 获取get传值
	// id := c.GetString("id")
	// beego.Info(id)
	// c.Ctx.WriteString("修改新闻--" + id)

	id, err := c.GetInt("id")
	if err != nil {
		beego.Info(err)
		c.Ctx.WriteString("入参错误")
		return
	}
	fmt.Printf("值：%v 类型：%T \n", id, id)

	c.Ctx.WriteString("修改新闻")
}
