package pac

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	"strings"
)

func OracleConnect(ruser string, pwd string, rhost string, rport string, sid string) (conn *sql.DB, err error) {
	conn, err = sql.Open("godror", fmt.Sprintf(`user=%s password="%s" connectString="%s:%s/%s"`, ruser, pwd, rhost, rport, sid))
	if err != nil {
		Info("连接出错")
	}
	err = conn.Ping()
	if err != nil {
		Err(err)
	}
	Info("Oracle数据库连接成功")
	resultSet, err := OracleCMD(fmt.Sprintf("select version from v$instance"), conn)
	for _, m := range resultSet {
		for _, value := range m {
			//fmt.Println(fmt.Sprintf("%s", value))
			Info(fmt.Sprintf("当前数据库版本为：%s", value))
		}
	}
	if err != nil {
		Err(err)
	}

	// isdba, err := OracleCMD(fmt.Sprintf("select userenv('ISDBA') from dual"), conn)
	isdba, err := OracleCMD("select userenv('ISDBA') from dual", conn)
	for _, m := range isdba {
		for _, value := range m {
			fmt.Println(fmt.Sprintf("%s", value))
			if strings.ToLower(fmt.Sprintf("%s", value)) == "true" {
				Success("当前账号为DBA权限")
			} else {
				Info("当前账号非DBA权限")
			}
		}
	}

	return conn, nil
}
