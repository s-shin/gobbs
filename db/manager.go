package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Manager struct {
	DBs map[string]*sql.DB
}

func NewManager() *Manager {
	return &Manager{
		make(map[string]*sql.DB),
	}
}

func (m *Manager) Connect(name string) *sql.DB {
	if db, ok := m.DBs[name]; ok {
		return db
	}
	db, err := sql.Open("sqlite3", GetDSN(name))
	if err != nil {
		log.Panic(err)
	}
	m.DBs[name] = db
	return db
}

func (m *Manager) DisconnectAll() {
	for name, db := range m.DBs {
		db.Close()
		delete(m.DBs, name)
	}
}
