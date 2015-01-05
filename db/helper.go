package db

import (
	"database/sql"
	"github.com/s-shin/gobbs/util"
	"log"
	"os"
	"path"
)

func Query(db *sql.DB, q string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(q, args...)
	if err != nil {
		log.Panic(err)
	}
	return rows
}

func QueryFile(db *sql.DB, filePath string) error {
	_, err := db.Exec(util.ReadFile(filePath))
	return err
}

func InitDB() error {
	os.Remove(GetDSN("master"))
	DisconnectAll()
	sqlFilePath := path.Join(util.ProjectDir(), "script", "sql", "init.sql")
	return QueryFile(Master(), sqlFilePath)
}

func Master() *sql.DB {
	return Connect("master")
}

func Slave() *sql.DB {
	return Connect("slave")
}
