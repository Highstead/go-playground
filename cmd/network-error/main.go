package main

import (
	"database/sql"
	"fmt"
	"net"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func main() {
	db, err := openTopology("127.0.0.1", 3306, 5)
	if err != nil {
		fmt.Println("error", err)
		fmt.Println("recoverable?", recoverableError(err))
		return
	}
	var rows *sql.Rows
	rows, err = db.Query("select * from mytable")
	if err != nil {
		fmt.Println("error", err)
		fmt.Println("recoverable?", recoverableError(err))
		return
	}
	fmt.Println("rows", rows)

}

func openTopology(host string, port int, readTimeout int) (*sql.DB, error) {
	mysql_uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/?timeout=%ds&readTimeout=%ds&interpolateParams=true",
		"root",
		"",
		host, port,
		readTimeout,
		readTimeout,
	)

	db, err := sql.Open("mysql", mysql_uri)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, err
}

func recoverableError(err error) bool {
	rootErr := errors.Cause(err)

	fmt.Println("Got to type assertion")
	// There ar a variety of libraries that will refuse connections and we're covering what we know
	if netErr, ok := rootErr.(net.Error); ok {
		if !netErr.Temporary() {
			return false
		}
	}

	fmt.Println("Got to switch")

	switch t := err.(type) {
	case *net.DNSError:
		return false
	//Os level error for connection refused
	case syscall.Errno:
		if t == syscall.ECONNREFUSED {
			return false
		}
	//Network errors for connection refused that use opError instead of network error
	case *net.OpError:
		if t.Op == "dial" {
			//This is a unknown host
			return false
		} else if t.Op == "read" {
			//connection refused
			return false
		}
	}
	fmt.Printf("%v", rootErr)
	return true
}
