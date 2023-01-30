package pac

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

// 产生随机字符
func RandStr(length int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func RandStrnum(length int) string {
	str := "1234567890"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func Creatable(conn *sql.DB) (table string) {
	table = RandStr(6)
	Info(table)
	sqlstr1 := fmt.Sprintf("CREATE TABLE %s (data varchar(2000));", table)
	Info(sqlstr1)
	MssqlCMD(sqlstr1, conn)

	return table
}

func Insertresult(table string, conn *sql.DB) {
	time.Sleep(time.Duration(1) * time.Second)

	Info("将结果写到表里面")
	sqlstr2 := fmt.Sprintf("BULK INSERT %s FROM 'c:\\test11.txt' WITH (ROWTERMINATOR ='\\n')", table)
	Info(sqlstr2)
	MssqlCMD(sqlstr2, conn)
	Info("查表取结果")
	sqlstr3 := fmt.Sprintf("select * from %s", table)
	Info(sqlstr3)
	Info("命令执行结果如下")
	MssqlCMD(sqlstr3, conn)

}
