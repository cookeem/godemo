package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gin-gonic/gin.v1"
	"io"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

/*
	CREATE TABLE users
	(
	nickname VARCHAR(20) DEFAULT '' NOT NULL,
	loginname VARCHAR(20) DEFAULT '' NOT NULL,
	passwd VARCHAR(20) DEFAULT '' NOT NULL,
	mobile VARCHAR(11) DEFAULT '' NOT NULL,
	code VARCHAR(20) DEFAULT '' NOT NULL
	);
	CREATE UNIQUE INDEX users_loginname_uindex ON users (loginname);
	CREATE UNIQUE INDEX users_mobile_uindex ON users (mobile);
*/

type Users struct {
	nickname, loginname, passwd, mobile, code string
}

func main() {
	r := gin.Default()
	r.StaticFile("/", "index.html")
	r.POST("/login", func(c *gin.Context) {
		loginname := c.DefaultPostForm("loginname", "")
		passwd := c.DefaultPostForm("passwd", "")

		success := 0
		msg := ""

		db, err := sql.Open("sqlite3", "./db.sqlite")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		stmt, err := db.Prepare("select nickname, loginname, passwd, mobile, code from users where loginname=? or mobile=?")
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		var u Users
		err = stmt.QueryRow(loginname, loginname).Scan(&u.nickname, &u.loginname, &u.passwd, &u.mobile, &u.code)
		if err != nil {
			success = 0
			msg = "登录失败，用户不存在"
		} else {
			if loginname == u.loginname && passwd == u.passwd {
				success = 1
				msg = fmt.Sprintf("账号登录成功: %v", u)
			} else if loginname == u.mobile && passwd == "" {
				seed := rand.Int()
				h := md5.New()
				io.WriteString(h, strconv.Itoa(seed))
				code := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil))[0:6])

				stmt, err = db.Prepare("update users set code=? where mobile=?")
				if err != nil {
					msg = err.Error()
				} else {
					_, err = stmt.Exec(code, loginname)
					if err != nil {
						msg = err.Error()
					} else {
						success = 1
						msg = fmt.Sprintf("请输入手机验证码登录：%v", code)
					}
					success = 1
					msg = fmt.Sprintf("请输入手机验证码登录：%v", code)
				}

			} else if loginname == u.mobile && passwd == u.code {
				success = 1
				msg = fmt.Sprintf("手机登录成功: %v", u)
			} else {
				msg = "登录失败，密码错误！"
			}
		}

		c.JSON(200, gin.H{
			"success": success,
			"msg":     msg,
		})
	})
	r.Run(":9090")
}
