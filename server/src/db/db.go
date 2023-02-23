package db

import (
	"database/sql"
	"project/zj"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var d *sql.DB

var mainTable = `combine`

// Conn ...
func Conn(dsn string) (err error) {

	d, err = sql.Open(`mysql`, dsn)
	if err != nil {
		return
	}

	d.SetConnMaxLifetime(time.Minute * 3)
	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(10)

	return
}

// Ping ...
func Ping() (err error) {
	return d.Ping()
}

// Close ...
func Close() {
	d.Close()
}

// WaitConn 服务器刚启的时候可能 app 启动了但是 mysql 没启动，一直等到 mysql 正常运行
func WaitConn(dns string) {

	for {
		err := Conn(dns)
		if err != nil {
			zj.W(`db`, err)
			time.Sleep(time.Second)
			continue
		}

		err = d.Ping()
		if err != nil {
			zj.W(`db`, err)
			time.Sleep(time.Second)
			continue
		}

		break
	}
}

// Clean ...
func Clean() {

	query := `TRUNCATE TABLE item`
	d.Exec(query)

	query = `TRUNCATE TABLE revision`
	d.Exec(query)
}
