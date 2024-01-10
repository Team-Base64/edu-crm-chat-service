package pg

import (
	e "main/domain/errors"
	m "main/domain/model"

	"github.com/lib/pq"
)

func (s *PostgreSqlStore) CreateChat(chat *m.Chat) (int, int, error) {
	var chatID int = -1
	var teacherID int = -1
	row := s.db.QueryRow(
		`SELECT teacherID FROM classes WHERE id = $1;`,
		chat.ClassID,
	)
	if err := row.Scan(&teacherID); err != nil {
		return -1, -1, e.StacktraceError(err)
	}

	err := s.db.QueryRow(
		`INSERT INTO chats (teacherID, studentID, classID) VALUES ($1, $2, $3) RETURNING id;`,
		teacherID, chat.StudentID, chat.ClassID,
	).Scan(&chatID)
	if err != nil {
		return -1, -1, e.StacktraceError(err)
	}
	return teacherID, chatID, nil
}

func (s *PostgreSqlStore) AddMessage(msg *m.CreateMessage) error {
	if msg.AttachmentURLs == nil {
		msg.AttachmentURLs = []string{}
	}

	_, err := s.db.Exec(`INSERT INTO messages (chatID, text, isAuthorTeacher, createtime, isRead, attaches) VALUES ($1, $2, $3, $4, $5, $6);`,
		msg.ChatID, msg.Text, msg.IsAuthorTeacher, msg.CreateTime, msg.IsRead, (*pq.StringArray)(&msg.AttachmentURLs))
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *PostgreSqlStore) GetChatsByClassID(classID int) ([]int, error) {
	rows, err := s.db.Query(
		`SELECT id FROM chats WHERE classID =  $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	ans := []int{}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, e.StacktraceError(err)
		}

		ans = append(ans, id)
	}
	return ans, nil
}

func (s *PostgreSqlStore) GetSocialTypeByChatID(chatID int) (string, error) {
	var studentID int32
	row := s.db.QueryRow(
		`SELECT studentID FROM chats WHERE id = $1;`, chatID)
	if err := row.Scan(&studentID); err != nil {
		return "", e.StacktraceError(err)
	}
	var type1 string
	row = s.db.QueryRow(
		`SELECT socialType FROM students WHERE id = $1;`, studentID)
	if err := row.Scan(&type1); err != nil {
		return "", e.StacktraceError(err)
	}
	return type1, nil
}
