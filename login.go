package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"os"
	"strings"
)

type LoginData struct {
	LoginUsername string `json:"username"`
	LoginPassword string `json:"password"`
}

func getEnv(key string, def string) string {
	if v := os.Getenv(strings.ToUpper(key)); v != "" {
		return v
	}
	return def
}

func login(ctx iris.Context)  {
	session := sessions.Get(ctx)
	if auth, _ := sessions.Get(ctx).GetBoolean("authenticated"); auth{
		ctx.JSON(iris.Map{
			"result":true,
		})
		return
	}
	//isNew := session.IsNew()

	var info LoginData
	err := ctx.ReadJSON(&info)
	// context.ReadJSON method
	// https://github.com/kataras/iris/blob/567c06702fa4359dc4835593a55c55854069954a/context/context.go#L2298
	if err != nil{
		fmt.Println("JSON Format Error")
		return
	}

	var user User
	if err := MYSQLDB.Where("username = ?", info.LoginUsername).First(&user).Error; err != nil {
		ctx.JSON(iris.Map{
			"result":false,
		})
		return
	}
	// here to connect to database, query Username/password
	CheckUserName := user.Username
	CheckPassword := user.Password

	if (info.LoginUsername == CheckUserName)  && (info.LoginPassword == CheckPassword) {
		session.Set("authenticated", true)
		ctx.JSON(iris.Map{
			"result":true,
		})
	} else {
		ctx.JSON(iris.Map{
			"result":false,
		})
	}
}

func logout(ctx iris.Context) {
	session := sessions.Get(ctx)
	// Revoke users authentication
	session.Set("authenticated", false)
	if auth, _ := sessions.Get(ctx).GetBoolean("authenticated"); !auth{
		ctx.JSON(iris.Map{
			"redirect": true,
		})
		return
	}
}
