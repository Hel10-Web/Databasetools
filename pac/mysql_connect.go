package pac

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

func MysqlConnect(Ruser string, Rhost string, PWD string, Rport string) (err error, conn *sql.DB, sign bool) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/information_schema?charset=gbk&parseTime=True", Ruser, PWD, Rhost, Rport)
	conn, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, false
	}
	err = conn.Ping()
	if err != nil {
		return nil, nil, false
	}
	fmt.Println("连接数据库成功！")
	sign = true
	return nil, conn, sign
}

func MysqlCrack(Rhost string, Rport string) {
	Info("开始爆破,请稍等.....")
	sign = false
	for _, user := range Userdict["mysql"] {
		for _, pass := range Passwords {
			pass = strings.Replace(pass, "{user}", user, -1)

			_, _, sign := MysqlConnect(user, Rhost, pass, Rport)
			if sign == true {
				Success(fmt.Sprintf("账号密码为：%s:%s", user, pass))
				os.Exit(0)
			} else {
				fmt.Println(fmt.Sprintf("%s:%s 未成功爆破出账号密码", user, pass))
			}
		}
	}
}
