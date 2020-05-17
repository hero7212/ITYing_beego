package controllers

import (
	"encoding/xml"
	"fmt"
	"strconv"

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

func (c *GoodsController) DoAdd() { // post
	c.Ctx.WriteString("执行增加操作")
}

type Product struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

func (c *GoodsController) DoEdit() { // put
	// title := c.GetString("title")

	p := Product{}

	if err := c.ParseForm(&p); err != nil {
		c.Ctx.WriteString("获取数据失败")
		return
	}
	fmt.Printf("%#v", p)

	c.Ctx.WriteString("执行修改操作---")
}

func (c *GoodsController) DoDelete() { // delete
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("参数错误")
	}
	c.Ctx.WriteString("执行删除操作--" + strconv.Itoa(id))
}

// 接收xml
// 配置文件中加 copyrequestbody = true
func (c *GoodsController) Xml() {

	p := Product{}

	// str := string(c.Ctx.Input.RequestBody)
	// beego.Info(str)

	var err error
	if e := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); e != nil {
		c.Data["json"] = err.Error()
	} else {
		fmt.Printf("%#v \n", p)
		c.Data["json"] = p
		c.ServeJSON()
	}
}
