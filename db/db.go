package db

import (
	"database/sql"
)

var manager = NewManager()

func Connect(name string) *sql.DB {
	return manager.Connect(name)
}

func DisconnectAll() {
	manager.DisconnectAll()
}
