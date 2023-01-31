package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"
)

// 执行sql命令行
func MssqlCMD(sqlstr string, conn *sql.DB) []interface{} {

	stmt, err := conn.Prepare(sqlstr)
	if err != nil {
		Err(err)
		return nil
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		Err(err)
		return nil
	}

	cols, _ := rows.Columns()
	var colsdata = make([]interface{}, len(cols))
	for i := 0; i < len(cols); i++ {
		colsdata[i] = new(interface{})
	}

	for rows.Next() {
		rows.Scan(colsdata...) //将查到的数据写入到这行中
		PrintRow(colsdata)     //打印此行
	}
	defer rows.Close()
	return colsdata
}

func PrintRow(colsdata []interface{}) (err error, result interface{}) {
	for _, val := range colsdata {
		switch v := (*(val.(*interface{}))).(type) {
		case nil:
			//fmt.Print("NULL")
		case bool:
			if v {
				fmt.Print("True")
			} else {
				fmt.Print("False")
			}

		case []byte:
			fmt.Print(string(v))
		case time.Time:
			fmt.Print(v.Format("2022-10-31 19:10:00.999"))
		default:
			fmt.Print(v)
		}
		fmt.Println()
	}
	return err, result
}

// 循环执行sql语句
func loopMssqlCMD(conn *sql.DB) {
	Info("执行mssql命令")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s:%s> ", Rhost, Rport)
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		MssqlCMD(cmd, conn)
	}
}
