package service

import (
	"errors"
	dbm "github.com/s-shin/gobbs/db"
	m "github.com/s-shin/gobbs/model"
	v "github.com/s-shin/gobbs/validation"
	"log"
	"time"
)

func FetchThreadPosts(threadId int64, offset, limit int) []*m.Post {
	rows := dbm.Query(dbm.Slave(), `
		SELECT id, thread_id, name, content, created_at
        FROM posts
        WHERE thread_id=?
        ORDER BY id
        LIMIT ?, ?
	`, threadId, offset, limit)
	defer rows.Close()
	posts := make([]*m.Post, 0, limit)
	for rows.Next() {
		p := new(m.Post)
		err := rows.Scan(&p.Id, &p.ThreadId, &p.Name, &p.Content, &p.CreatedAt)
		if err != nil {
			log.Panic(err)
		}
		posts = append(posts, p)
	}
	return posts
}

func FetchAllThreadPosts(threadId int64) []*m.Post {
	return FetchThreadPosts(threadId, 0, 1000)
}

func CreatePost(threadId int64, name string, content string) (int64, error) {
	p := &m.Post{ThreadId: threadId, Name: name, Content: content}

	// validation
	{
		var vErr v.Error
		if e := p.Validate(); e != nil {
			vErr = vErr.Combine(e)
		}
		if vErr != nil {
			return 0, vErr
		}
	}

	db := dbm.Master()
	tx, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	{
		var postNum int64
		err := tx.QueryRow(`
			SELECT post_num FROM threads WHERE id=? -- FOR UPDATE
		`, threadId).Scan(&postNum)
		if err != nil {
			tx.Rollback()
			log.Panic(err)
		}
		if postNum == 1000 {
			tx.Rollback()
			return 0, errors.New("Max posts in this thread.")
		}
	}
	now := time.Now().Unix()
	{
		_, err := tx.Exec(`
			UPDATE threads SET post_num=post_num+1, updated_at=? WHERE id=?
		`, now, threadId)
		if err != nil {
			tx.Rollback()
			log.Panic(err)
		}
	}
	var lastInsertId int64
	{
		result, err := tx.Exec(`
			INSERT INTO posts (thread_id, name, content, created_at)
			VALUES (?, ?, ?, ?)
		`, threadId, p.Name, p.Content, now)
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
	tx.Commit()
	return lastInsertId, nil
}
