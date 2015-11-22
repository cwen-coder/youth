package main

import (
	// "encoding/json"
	// "fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"html/template"
	"youth/check"
	"youth/model"
)

type msg struct {
	ErrorStatue   bool
	SuccessStatue bool
	SuccessInfo   string
	ErrorInfo     string
}

type SiginIfo struct {
	UserName   string `form:"user_name"`
	UserNumber string `form:"user_number"`
}

type EditReturnInfo struct {
	Status  bool
	Message string
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Extensions: []string{".tmpl", ".html"},
		Charset:    "UTF-8",
		Funcs: []template.FuncMap{
			{
				"equal": func(args ...interface{}) bool {
					return args[0] == args[1]
				},
			},
		},
	}))
	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))
	m.Get("/", youth)
	m.Get("/firtConfirm", firtConfirm)
	m.Post("/firtConfirm", binding.Form(model.User{}), firtConfirmPost)
	m.Post("/userSiginCheck", binding.Bind(SiginIfo{}), userSiginCheck)
	m.Get("/userInforEdit", userInforEdit)
	m.Post("/userInforEdit", binding.Form(model.User{}), userInforEditPost)
	m.Get("/editReruenInfo/:status", editReruenInfo)
	m.Run()
}

func youth(r render.Render) {
	r.HTML(200, "home", nil)
}

func firtConfirm(r render.Render) {
	r.HTML(200, "firtConfirm", nil)
}

func firtConfirmPost(user model.User, r render.Render) {
	var result bool
	user.UserName, result = check.CheckUserName(user.UserName)
	if result != true {
		msgInfo := msg{
			ErrorStatue:   true,
			SuccessStatue: false,
			SuccessInfo:   "成功",
			ErrorInfo:     "对不起！你输入的姓名有误",
		}
		r.HTML(200, "firtConfirm", msgInfo)
		return
	}

	user.UserNumber, result = check.CheckNumber(user.UserNumber)
	if result != true {
		msgInfo := msg{
			ErrorStatue:   true,
			SuccessStatue: false,
			SuccessInfo:   "成功",
			ErrorInfo:     "对不起！你不在名单中，请仔细检查重新填写",
		}
		r.HTML(200, "firtConfirm", msgInfo)
		return
	}

	isSelected := model.CheckSelected(user.UserNumber)
	if isSelected == true {
		msgInfo := msg{
			ErrorStatue:   true,
			SuccessStatue: false,
			SuccessInfo:   "成功",
			ErrorInfo:     "对不起！你不是首次确认信息，如要修改请返回首页点击【修改确认】",
		}
		r.HTML(200, "firtConfirm", msgInfo)
		return
	}

	var result1 bool
	var result2 bool
	var result3 bool
	var result4 bool

	user.UserSex, result1 = check.CheckSex(user.UserSex)
	user.UserBrand, result2 = check.CheckBrand(user.UserBrand)
	user.UserColor, result3 = check.CheckColor(user.UserColor)
	user.UserSize, result4 = check.CheckSize(user.UserSize)
	// result1 != true || result2 != true || result3 != true || result4 != true
	if result1 != true || result2 != true || result3 != true || result4 != true {
		msgInfo := msg{
			ErrorStatue:   true,
			SuccessStatue: false,
			SuccessInfo:   "成功",
			ErrorInfo:     "对不起！你填写的信息有误",
		}
		r.HTML(200, "firtConfirm", msgInfo)
		return
	}
	user.IsSelected = true
	model.Update(user)
	msgInfo := msg{
		ErrorStatue:   false,
		SuccessStatue: true,
		SuccessInfo:   "恭喜你！确认成功",
		ErrorInfo:     "对不起！你填写的信息有误",
	}
	r.HTML(200, "firtConfirm", msgInfo)
}

func userSiginCheck(info SiginIfo, session sessions.Session) string {
	var result bool
	info.UserName, result = check.CheckUserName(info.UserName)
	if result != true {
		return "-1"
	}

	info.UserNumber, result = check.CheckNumber(info.UserNumber)
	if result != true {
		return "-2"
	}

	isSelected := model.CheckSelected(info.UserNumber)
	if isSelected == false {
		return "-3"
	}

	resultSigin := model.Sigin(info.UserName, info.UserNumber)
	if resultSigin != true {
		return "-4"
	} else {
		session.Set("user_name", info.UserName)
		session.Set("user_number", info.UserNumber)

		return "1"
	}
}

func userInforEdit(session sessions.Session, r render.Render) {
	user_number := session.Get("user_number")
	user_name := session.Get("user_name")
	if user_number == nil || user_name == nil {
		r.HTML(200, "home", nil)
	}
	var name string
	var number string
	if value, ok := user_name.(string); ok {
		name = value
	} else {
		r.HTML(200, "home", nil)
	}

	if value, ok := user_number.(string); ok {
		number = value
	} else {
		r.HTML(200, "home", nil)
	}

	var user model.User
	user, err := model.GetUserInfo(name, number)
	if err != true {
		r.HTML(200, "home", nil)
	} else {
		r.HTML(200, "userInforEdit", user)
	}
}

func userInforEditPost(user model.User, r render.Render) {
	var result bool
	user.UserName, result = check.CheckUserName(user.UserName)
	if result != true {
		r.Redirect("/editReruenInfo/error")
		return
	}

	user.UserNumber, result = check.CheckNumber(user.UserNumber)
	if result != true {
		r.Redirect("/editReruenInfo/error")
		return
	}

	isSelected := model.CheckSelected(user.UserNumber)
	if isSelected != true {
		r.Redirect("/editReruenInfo/error")
		return
	}

	var result1 bool
	var result2 bool
	var result3 bool
	var result4 bool

	user.UserSex, result1 = check.CheckSex(user.UserSex)
	user.UserBrand, result2 = check.CheckBrand(user.UserBrand)
	user.UserColor, result3 = check.CheckColor(user.UserColor)
	user.UserSize, result4 = check.CheckSize(user.UserSize)
	// result1 != true || result2 != true || result3 != true || result4 != true
	if result1 != true || result2 != true || result3 != true || result4 != true {
		r.Redirect("/editReruenInfo/error")
		return
	}
	user.IsSelected = true
	model.Update(user)
	r.Redirect("/editReruenInfo/success")
}

func editReruenInfo(params martini.Params, session sessions.Session, r render.Render) {
	user_number := session.Get("user_number")
	user_name := session.Get("user_name")
	if user_number == nil || user_name == nil {
		r.HTML(200, "home", nil)
	}
	if params["status"] == "success" {
		msg := EditReturnInfo{
			Status:  true,
			Message: "修改成功！",
		}
		session.Delete("user_name")
		session.Delete("user_number")
		r.HTML(200, "editReruenInfo", msg)
	} else {
		msg := EditReturnInfo{
			Status:  false,
			Message: "对不起！修改失败！",
		}
		r.HTML(200, "editReruenInfo", msg)
	}
}
