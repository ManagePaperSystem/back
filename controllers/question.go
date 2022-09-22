package controllers

import (
	"APIProject/models"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
)

// Operations about Users
type QuesController struct {
	beego.Controller
}

// @Title Generate
// @Description 生成题目
// @Param	Account		query 	string		true		"账号"
// @Param	Phase		query 	string		true		"年级"
// @Param	Number		query 	string		true		"数量"
// @Success 200 {[]object} ans
// @Failure 403 body is empty
// @router /gen [post]
func (q *QuesController) Generate() {
	a, p, i := q.GetString("Account"), q.GetString("Phase"), q.GetString("Number")
	n, _ := strconv.Atoi(i)
	ans := models.Generate(a, p, n)
	q.Data["json"] = ans
	q.ServeJSON()
}

// @Title Check
// @Description 判断题目
// @Param	Account		query 	string		true		"账号"
// @Param	Question	query 	string		true		"题目表达式"
// @Param	Choice		query 	string		true		"选项"
// @Success 200 {[]bool} ans
// @Failure 403 body is empty
// @router /check [post]
func (q *QuesController) Check() {
	a, q2, c := q.GetString("Account"), q.GetString("Question"), q.GetString("Choice")
	QuestionList, ChoiceList := strings.Split(q2, "_"), strings.Split(c, "_")
	ans := models.Check(a, QuestionList, ChoiceList)
	q.Data["json"] = ans
	q.ServeJSON()
}

// @Title CheckList
// @Description 题目列表
// @Param	Account		query 	string		true		"账号"
// @Success 200 {[]string} ans
// @Failure 403 body is empty
// @router /check/list [post]
func (q *QuesController) CheckList() {
	a := q.GetString("Account")
	q.Data["json"] = models.CheckList(a)
	q.ServeJSON()
}

// @Title CheckPaper
// @Description 查看题目
// @Param	Account		query 	string		true		"账号"
// @Param	TimeStr		query 	string		true		"时间"
// @Success 200 {[]object} ans
// @Failure 403 :uid is empty
// @router /check/paper [post]
func (q *QuesController) CheckPaper() {
	a, t := q.GetString("Account"), q.GetString("TimeStr")
	q.Data["json"] = models.CheckPaper(a, t)
	q.ServeJSON()
}

// @Title DeletePaper
// @Description 删除题目
// @Param	Account		query 	string		true		"账号"
// @Param	TimeStr		query 	string		true		"时间"
// @Success 200 {bool} flag
// @Failure 403 :uid is empty
// @router /delete [post]
func (q *QuesController) DeletePaper() {
	a, t := q.GetString("Account"), q.GetString("TimeStr")
	q.Data["json"] = map[string]bool{"flag": models.DeletePaper(a, t)}
	q.ServeJSON()
}
