package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

func postgrecmd(sqlstr string, conn *sql.DB) ([]map[string]interface{}, error) {
	rows, err := conn.Query(sqlstr)
	if err != nil {
		Err(err)
		return nil, nil
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			log.Fatal(err)
		}
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			log.Fatal(err)
		}

		row := make(map[string]interface{})
		for i, column := range columns {
			row[column] = values[i]
		}

		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	for i, _ := range result {
		for _, w := range result[i] {
			fmt.Println(fmt.Sprintf("%s", w))
		}
	}
	return result, err
}

// 循环执行sql语句
func loopPostgreCMD(conn *sql.DB) {
	Info("执行PostgreSQL命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		_, err := postgrecmd(cmd, conn)

		if err != nil {
			Info("循环执行sql语句报错")
		}

	}
}

func postgreisdba(conn *sql.DB) {
	result, err := postgrecmd("SELECT current_setting('is_superuser');", conn)
	if err != nil {
		Err(err)
	}
	for i, _ := range result {
		for _, w := range result[i] {
			OnOrOFF := fmt.Sprintf("%s", w)
			if OnOrOFF == "on" {
				Info("当前用户为管理员权限")
			} else {
				Info("非管理员权限")
			}
		}
	}
}
