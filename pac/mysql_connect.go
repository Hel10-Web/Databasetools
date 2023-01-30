package pac

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MysqlConnect(Ruser string, Rhost string, PWD string, Rport string) (err error, conn *sql.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/information_schema?charset=gbk&parseTime=True", Ruser, PWD, Rhost, Rport)
	conn, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("打开数据库失败,err:%v\n", err)
		os.Exit(1)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = conn.Ping()
	if err != nil {
		fmt.Printf("连接数据库失败,err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("连接数据库成功！")
	m, err := MysqlCMD("select @@version;", conn)
	fmt.Printf("数据库版本：Mysql %v\n", m[0]["@@version"])

	return nil, conn
}
