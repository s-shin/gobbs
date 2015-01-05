package service

import (
	dbm "github.com/s-shin/gobbs/db"
	m "github.com/s-shin/gobbs/model"
	v "github.com/s-shin/gobbs/validation"
	"log"
	"time"
)

func FetchThreadById(id int64) *m.Thread {
	t := new(m.Thread)
	err := dbm.Slave().QueryRow(`
		SELECT id, title, default_name, post_num, created_at, updated_at
		FROM threads
		WHERE id=?
	`, id).Scan(&t.Id, &t.Title, &t.DefaultName, &t.PostNum, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		log.Panic(err)
	}
	return t
}

func FetchRecentThreads(offset, limit int) []*m.Thread {
	rows := dbm.Query(dbm.Slave(), `
		SELECT id, title, default_name, post_num, created_at, updated_at
		FROM threads
		ORDER BY updated_at DESC, id DESC
		LIMIT ?, ?
	`, offset, limit)
	defer rows.Close()
	threads := make([]*m.Thread, 0, limit)
	for rows.Next() {
		t := new(m.Thread)
		err := rows.Scan(&t.Id, &t.Title, &t.DefaultName, &t.PostNum, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			log.Panic(err)
		}
		threads = append(threads, t)
	}
	return threads
}

func CreateThread(threadTitle, threadDefaultName, postName, postContent string) (int64, error) {
	var vErr v.Error

	t := &m.Thread{Title: threadTitle, DefaultName: threadDefaultName}
	if e := t.Validate(); e != nil {
		vErr = vErr.Combine(e)
	}

	p := &m.Post{Name: postName, Content: postContent}
	if e := p.Validate(); e != nil {
		vErr = vErr.Combine(e)
	}

	if vErr != nil {
		return 0, vErr
	}

	db := dbm.Master()
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	now := time.Now().Unix()
	var lastInsertId int64
	{
		result, err := tx.Exec(`
			INSERT INTO threads (title, default_name, created_at, updated_at)
			VALUES (?, ?, ?, ?)
		`, t.Title, t.DefaultName, now, now)
		if err != nil {
			tx.Rollback()
			log.Panic(err)
		}
		lastInsertId, err = result.LastInsertId()
		if err != nil {
			tx.Rollback()
			log.Panic(err)
		}
	}
	{
		_, err = tx.Exec(`
			INSERT INTO posts (thread_id, name, content, created_at)
			VALUES (?, ?, ?, ?)
		`, lastInsertId, p.Name, p.Content, now)
		if err != nil {
			tx.Rollback()
			log.Panic(err)
		}
	}
	tx.Commit()
	return lastInsertId, nil
}
