package controllers

import (
	"BeegoPackage0922/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1、获取请求数据
	c.GetString("user")//返回字符串
	c.GetInt("age")//返回整数
	userName := c.Ctx.Input.Query("user")
	password := c.Ctx.Input.Query("psd")
	//2、使用固定数据进行数据校验
	//admin  123456
	if userName != "admin" || password != "123456"{
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误。"))
		return
	}

	//校验正确的情况
	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据校验成功。"))
}

func (c *MainController) Post(){

	//1、解析前端提交的json格式的数据
	var person models.Person
	dataBytes, err :=ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err = json.Unmarshal(dataBytes,&person)
	if err != nil {
		fmt.Println(err.Error())
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("姓名:",person.Name)
	fmt.Println("年龄：",person.Age)
	fmt.Println("性别:",person.Sex)
	//fmt.Println()
	c.Ctx.WriteString("数据解析成功")
}

