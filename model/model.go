package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type Activity struct {
	Id        int64     `form:"id"`
	Name      string    `xorm:"varchar(10)" form:"name"`
	StartTime time.Time `form:"start_time"`
	EndTime   time.Time `form:"end_time"`
}

type User struct {
	Id int64 `form:"id"`
	//Activity   Activity `xorm:"activity_id bigint" form:"activily_id"`
	UserName   string `xorm:"varchar(10)" form:"user_name"`
	UserNumber string `xorm:"varchar(10)" form:"user_number"`
	UserSex    string `xorm:"varchar(10)" form:"user_sex"`
	UserBrand  string `xorm:"varchar(20)" form:"user_brand"`
	UserColor  string `xorm:"varchar(10)" form:"user_color"`
	UserSize   string `xorm:"varchar(10)" form:"user_size"`
	IsSelected bool
}

var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "000000:00000@/youth?charset=utf8")
	err = engine.Sync(
		new(Activity),
		new(User),
	)
	if err != nil {
		panic(err)
	}
}

func CheckNumber(num string) bool {
	user := new(User)
	count, _ := engine.Where("user_number = ?", num).Count(user)
	if count == 0 {
		return false
	}
	return true
}

func Update(user User) bool {
	_, err := engine.Where("user_number = ?", user.UserNumber).Cols("user_name", "user_sex", "user_brand", "user_color", "user_size", "is_selected").Update(&user)
	if err != nil {
		return false
	}
	return true
}

func CheckSelected(num string) bool {
	user := new(User)
	engine.Where("user_number = ?", num).Cols("is_selected").Get(user)
	return user.IsSelected
}

func Sigin(user_name string, user_number string) bool {
	user := new(User)
	count, _ := engine.Where("user_name = ? and user_number = ?", user_name, user_number).Count(user)
	if count == 0 {
		return false
	}
	return true
}

func GetUserInfo(user_name string, user_number string) (User, bool) {
	var user User
	_, err := engine.Where("user_name = ? and user_number = ?", user_name, user_number).Get(&user)
	if err != nil {
		return user, false
	} else {
		return user, true
	}
}
