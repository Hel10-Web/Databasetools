package pac

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"
)

var Userdict = map[string][]string{
	"mysql":      {"root", "mysql"},
	"mssql":      {"sa", "sql"},
	"postgresql": {"postgres", "admin"},
	"oracle":     {"system", "sys", "admin", "test", "web", "orcl"},
	"redis":      {"redis"},
}

var Passwords = []string{"123456", "admin", "admin123", "root", "", "pass123", "pass@123", "password", "123123", "654321", "111111", "123", "1", "admin@123", "Admin@123", "admin123!@#", "{user}", "{user}1", "{user}111", "{user}123", "{user}@123", "{user}_123", "{user}#123", "{user}@111", "{user}@2019", "{user}@123#4", "P@ssw0rd!", "P@ssw0rd", "Passw0rd", "qwe123", "12345678", "test", "test123", "123qwe", "123qwe!@#", "123456789", "123321", "666666", "a123456.", "123456~a", "123456!a", "000000", "1234567890", "8888888", "!QAZ2wsx", "1qaz2wsx", "abc123", "abc123456", "1qaz@WSX", "a11111", "a12345", "Aa1234", "Aa1234.", "Aa12345", "a123456", "a123123", "Aa123123", "Aa123456", "Aa12345.", "sysadmin", "system", "1qaz!QAZ", "2wsx@WSX", "qwe123!@#", "Aa123456!", "A123456s!", "sa123456", "1q2w3e", "Charge123", "Aa123456789", "postgres"}

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
