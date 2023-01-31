package pac

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

func OracleCMD(sqlstr string, conn *sql.DB) ([]map[string]interface{}, error) {
	rows, err := conn.QueryContext(context.Background(), fmt.Sprintf("%s", sqlstr))
	if err != nil {
		Err(err)
		return nil, nil
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		Err(err)
	}
	resultSet := make([]map[string]interface{}, 0)
	for rows.Next() {
		row := make([]interface{}, len(cols))
		rowPtrs := make([]interface{}, len(cols))
		for i := range row {
			rowPtrs[i] = &row[i]
		}
		if err := rows.Scan(rowPtrs...); err != nil {
			log.Fatal(err)
		}
		entry := make(map[string]interface{})
		for i, colName := range cols {
			val := row[i]
			b, ok := val.([]byte)
			if ok {
				entry[colName] = string(b)
			} else {
				entry[colName] = val
			}
		}
		resultSet = append(resultSet, entry)
	}
	return resultSet, err
}

// 循环执行sql语句
func loopOracleCMD(conn *sql.DB) {
	Info("执行Oracle SQL命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		resultSet, err := OracleCMD(cmd, conn)
		for _, m := range resultSet {
			for _, value := range m {
				fmt.Println(value)
			}
		}

		if err != nil {
			Info("循环执行sql语句报错")
		}
	}
}
