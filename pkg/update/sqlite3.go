package update

import (
	"database/sql"
	"github.com/fumanne/IP2Country/pkg/utils"
	"os"

	_ "github.com/mattn/go-sqlite3"
)


func Prepare(db *sql.DB) {
	initSql := "CREATE TABLE IF NOT EXISTS ip2country (start BIGINT NOT NULL, end BIGINT NOT NULL, country CHARACTER(10) NOT NULL);"
	initSqlIndex := "CREATE INDEX  IF NOT EXISTS  start_end_index ON ip2country (start, end);"
	_, err := db.Exec(initSql)
	utils.Checkerr(err)
	_, err = db.Exec(initSqlIndex)
	utils.Checkerr(err)

}

func Clean(dbfile string) {
	if _, e := os.Stat(dbfile); e == nil {
		err := os.Remove(dbfile)
		utils.Checkerr(err)
	}
}

//func IsAlive(db *sql.DB) bool {
//	if err := db.Ping(); err != nil {
//		return false
//	}
//	return true
//}