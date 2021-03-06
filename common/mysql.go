package common

import (
	"database/sql"
	"fmt"
	"time"
)

func ScanMysql(ip string, port string, username string, password string) (err error, result bool) {
	result = false
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/?timeout=%ds", username, password, ip+":"+port, time.Second*3)
	db, err := sql.Open("mysql", connStr)
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			defer db.Close()
			result = true
		}
	}
	return err, result
}

