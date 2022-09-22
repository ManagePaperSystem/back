package controllers

import (
	"APIProject/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description 登录账户
// @Param	Account		query 	string		true	"账号"
// @Param	Password	query 	string		true	"密码"
// @Success 200 {bool} flag
// @Failure 403 body is empty
// @router /login [post]
func (u *UserController) Login() {
	var U models.User
	U.Account, U.Password = u.GetString("Account"), u.GetString("Password")
	flag := models.Login(U)
	// flag == false 代表账号密码错误
	u.Data["json"] = map[string]bool{"flag": flag}
	u.ServeJSON()
}

// @Title GetCode
// @Description 请求验证码
// @Param	Account		query 	string		true	"邮箱"
// @Success 200 {bool} flag
// @Failure 403 user not exist
// @router /getcode [post]
func (u *UserController) GetCode() {
	Account := u.GetString("Account")
	flag := models.GetCode(Account)
	// flag == false 代表该邮箱已被注册
	u.Data["json"] = map[string]bool{"flag": flag}
	u.ServeJSON()
}

// @Title Register
// @Description 检查验证码
// @Param	Account		query 	string		true		"邮箱地址"
// @Param	Code		query 	string		true		"验证码"
// @Success 200 {bool} flag
// @Failure 403 user not exist
// @router /reg [post]
func (u *UserController) Register() {
	Account, Passcode := u.GetString("Account"), u.GetString("Passcode")
	flag := models.Register(Account, Passcode)
	// flag == false 代表验证码错误
	u.Data["json"] = map[string]bool{"flag": flag}
	u.ServeJSON()
}

// @Title SetAccount
// @Description 设置密码和年级
// @Param	Account		query 	string		true		"账号"
// @Param	Password	query 	string		true		"密码"
// @Success 200 {nil}
// @Failure 403 :uid is not int
// @router /set [post]
func (u *UserController) SetAccount() {
	a, p := u.GetString("Account"), u.GetString("Password")
	models.SetAccount(a, p)
	u.ServeJSON()
}

// @Title ModifyPassword
// @Description 修改密码
// @Param	Account		query 	string		true		"账号"
// @Param	Password	query 	string		true		"密码"
// @Success 200 {bool} flag
// @Failure 403 :uid is not int
// @router /modify [post]
func (u *UserController) ModifyPassword() {
	a, p := u.GetString("Account"), u.GetString("Password")
	fmt.Println(a, "  ", p)
	models.ModifyPassword(a, p)
	u.Data["json"] = map[string]bool{"flag": true}
	u.ServeJSON()
}
