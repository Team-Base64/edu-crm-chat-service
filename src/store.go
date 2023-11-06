package chat

import (
	"context"
	m "main/domain/model"
	"time"

	"github.com/jackc/pgx/v5"
)

type StoreInterface interface {
	AddMessage(in *m.CreateMessage) error
	GetChatsByClassID(chatID int) (*[]int, error)
}

type Store struct {
	db *pgx.Conn
}

func NewStore(db *pgx.Conn) StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) AddMessage(in *m.CreateMessage) error {
	_, err := s.db.Exec(
		context.Background(),
		`INSERT INTO messages (chatID, text, isAuthorTeacher, time, isRead) VALUES ($1, $2, $3, $4, $5);`,
		in.ChatID, in.Text, in.IsAuthorTeacher, time.Now().Format("2006.01.02 15:04:05"), in.IsRead,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetChatsByClassID(chatID int) (*[]int, error) {
	rows, err := s.db.Query(
		context.Background(),
		`SELECT id FROM chats WHERE classID =  $1;`,
		chatID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ans := []int{}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		ans = append(ans, id)
	}
	return &ans, nil
}
