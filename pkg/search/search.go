package search

import (
	"database/sql"
	"github.com/fumanne/IP2Country/pkg/update"
	"github.com/fumanne/IP2Country/pkg/utils"
)

func Do(ipaddress string) string {
	target := utils.Ip2long(ipaddress)
	if ! utils.IsExist(utils.DBFile) {
		update.Do(false)
	}
	db, _ := sql.Open("sqlite3", utils.DBFile)
	defer db.Close()
	stmt, err := db.Prepare("select country from ip2country where start <= ? and end >= ?")
	utils.Checkerr(err)
	var c string
	err = stmt.QueryRow(target, target).Scan(&c)
	utils.Checkerr(err)
	return c
}
