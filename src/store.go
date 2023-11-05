package chat

import (
	m "main/domain/model"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	AddMessage(in *m.CreateMessage) error
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) AddMessage(in *m.CreateMessage) error {
	_, err := s.db.Exec(`INSERT INTO messages (chatID, text, isAuthorTeacher, time, isRead) VALUES ($1, $2, $3, $4, $5);`, in.ChatID, in.Text, in.IsAuthorTeacher, time.Now().Format("2006.01.02 15:04:05"), in.IsRead)
	if err != nil {
		return err
	}
	return nil
}
