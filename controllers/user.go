package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

	c.Ctx.WriteString("用户中心")
}

func (c *UserController) AddUser() {

	c.TplName = "userAdd.html"
}

// 处理post请求 获取post提交的数据
func (c *UserController) DoAddUser() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Ctx.WriteString("id必须是int类型")
		return
	}
	// fmt.Printf("%v---%T", id, id)
	username := c.GetString("username")
	password := c.GetString("password")
	hobby := c.GetString("hobby")

	fmt.Printf("值：%v---类型：%T", hobby, hobby)

	c.Ctx.WriteString("用户中心--" + strconv.Itoa(id) + username + password)
}

func (c *UserController) EditUser() {

	c.TplName = "userEdit.html"
}

// 定义一个user的结构体
type User struct {
	Username string   `form:"username" json:"username"`
	Password string   `form:"password" json:"password"`
	Hobby    []string `form:"hobby" json:"hobby"`
}

func (c *UserController) DoEditUser() {

	u := User{}
	// err := c.ParseForm(&u)
	if err := c.ParseForm(&u); err != nil {
		c.Ctx.WriteString("Post提交失败")
	}
	fmt.Printf("%#v", u)
	c.Ctx.WriteString("解析json数据成功")
}

func (c *UserController) GetUser() {
	u := User{
		Username: "张三",
		Password: "123456",
		Hobby:    []string{"1", "2"},
	}
	// 返回一个json数据
	c.Data["json"] = u
	c.ServeJSON()
}
