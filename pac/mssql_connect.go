package pac

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"time"
)

var (
	isdebug  = true
	user     = "sa"
	database = "master"
)

func MssqlConnect(Rhost string, Rport string, pwd string) (err error, db *sql.DB) {

	_, err = WrapperTcpWithTimeout("tcp", Socks5Proxy, time.Second*10)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		Err(err)
	}

	var password = pwd
	var server = Rhost
	var port = Rport

	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s", server, port, database, user, password)
	if isdebug {
		Info(connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		Err(err)
	} else {
		MssqlCMD("select @@version;", conn)
		Success("连接成功!")
	}

	return err, conn
}
