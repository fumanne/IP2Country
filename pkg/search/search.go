package search

import (
	"database/sql"
	"github.com/fumanne/IP2Country/pkg/update"
	"github.com/fumanne/IP2Country/pkg/utils"
	"os"
	"path/filepath"
)

func Do(ipaddress string) string {
	target := utils.Ip2long(ipaddress)
	dbFile := filepath.Join(utils.Locate(update.DOWNLOAD), "ip.db")
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		update.Do()
	}

	DB, _ := sql.Open("sqlite3", dbFile)
	stmt, err := DB.Prepare("select country from ip2country where start <= ? and end >= ?")
	utils.Checkerr(err)
	var c string
	err = stmt.QueryRow(target, target).Scan(&c)
	utils.Checkerr(err)
	return c
}
