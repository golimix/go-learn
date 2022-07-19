package main

import (
	"encoding/gob"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type user struct {
	Name string
	Age  int
}

// LIMIX_INFO 会话的意义-可用于保存多次请求之间的状态信息
// 1. 用户认证信息,例如用户标识
// 2. 用户缓存信息,例如用户配置信息
func main() {
	// LIMIX_INFO 注册会话可保存类型
	gob.Register(user{})

	r := gin.Default()

	// 1. 注册中间件
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 2. 样例程序-信息->会话
	r.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)

		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			//记着调用save方法，写入session
			error := session.Save()
			if error == nil {
				panic("session write error")
			}
		}
		//把结构体存入session中
		session.Set("user", user{"hanyun", 30})
		session.Save()

		c.JSON(200, gin.H{"hello": session.Get("hello")})
	})

	// 3. 样例程序-会话->信息
	r.GET("/user", func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
	r.Run(":8080")
}
