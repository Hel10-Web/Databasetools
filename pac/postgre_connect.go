package pac

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbname = "postgres"
)

func postgre_connect(Rhost string, Rport string, Ruser string, PWD string) (conn *sql.DB) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", Rhost, Rport, Ruser, PWD, dbname)
	Info(psqlInfo)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	Success("连接成功")
	return conn
}
