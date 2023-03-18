package controllers

import (
	"fmt"
	"pair_server/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) GetHome() {
	c.Data["isHome"] = true
	c.TplName = "index.tpl"
}

func (c *MainController) ShowAll() {
	o := orm.NewOrm()
	var results []orm.Params
	num, err := o.Raw("SELECT * FROM pairs").Values(&results)
	if err != nil {
		logs.Info("查询失败：", err)
		return
	}
	logs.Info("查询成功，行数：", num)

	c.Data["number"] = num
	c.Data["results"] = results
	c.Data["isAll"] = true
	c.TplName = "show.tpl"
}

func (c *MainController) Add() {
	c.Data["isAdd"] = true
	c.TplName = "add.tpl"
}

func (c *MainController) DoAdd() {
	o := orm.NewOrm()
	p := models.Pairs{Chs: c.GetString("chs"), En: c.GetString("eng")}
	num, err := o.Insert(&p)
	if err != nil {
		logs.Info("插入失败：", err)
		return
	}
	logs.Info("插入成功，行数：", num)

	c.Redirect("/add", 302)
}

func (c *MainController) Search() {
	keyword := c.GetString("word")
	o := orm.NewOrm()
	var results []orm.Params
	lookupString := fmt.Sprintf("SELECT chs, en FROM pairs WHERE chs like '%%%s%%' OR en like '%%%s%%'", keyword, keyword)
	num, err := o.Raw(lookupString).Values(&results)
	if err != nil {
		logs.Info("查询失败：", err)
		return
	}
	logs.Info("查询成功，行数：", num)

	c.Data["number"] = num
	c.Data["results"] = results
	c.Data["isAll"] = true
	c.TplName = "show.tpl"
}

func (c *MainController) Delete() {
	id, err := strconv.Atoi(c.Ctx.Input.Param(":id"))
	if err != nil {
		logs.Info("转换失败：", err)
		return
	}
	o := orm.NewOrm()
	p := models.Pairs{Id: id}
	num, err := o.Delete(&p)
	if err != nil {
		logs.Info("删除失败：", err)
		return
	}
	logs.Info("删除成功，行数：", num)
	c.Redirect("/all", 302)
}
