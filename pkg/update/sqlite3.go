package update

import (
	"database/sql"
	"github.com/fumanne/IP2Country/pkg/utils"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

// todo: how to single ?
var DB, _ = sql.Open("sqlite3", filepath.Join(utils.Locate(DOWNLOAD), "ip.db"))

func Prepare() {
	initSql := "CREATE TABLE IF NOT EXISTS ip2country (start BIGINT NOT NULL, end BIGINT NOT NULL, country CHARACTER(10) NOT NULL);"
	initSqlIndex := "CREATE INDEX  IF NOT EXISTS  start_end_index ON ip2country (start, end);"
	_, err := DB.Exec(initSql)
	utils.Checkerr(err)
	_, err = DB.Exec(initSqlIndex)
	utils.Checkerr(err)

}

func Clean() {
	cleanSql := "Delete From ip2country"
	DB.Exec(cleanSql)
}
