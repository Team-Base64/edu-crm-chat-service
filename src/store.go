package chat

import (
	"context"
	"time"

	e "main/domain/errors"
	m "main/domain/model"
	proto "main/src/proto"

	"github.com/jackc/pgx/v5"
)

type StoreInterface interface {
	AddMessage(in *m.CreateMessage) error
	GetChatsByClassID(chatID int) (*[]int, error)
	ValidateToken(tok string) (int, error)
	CreateStudent(in *proto.CreateStudentRequest) (int, error)
	CreateChat(in *proto.CreateChatRequest) (int, error)
	GetHomeworksByChatID(classID int) ([]*proto.HomeworkData, error)
	CreateSolution(in *proto.SendSolutionRequest) error
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
		return e.StacktraceError(err)
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
	return &ans, nil
}

func (s *Store) ValidateToken(tok string) (int, error) {
	var classID int = -1
	row := s.db.QueryRow(
		context.Background(),
		`SELECT id FROM classes WHERE inviteToken = $1;`, tok)
	if err := row.Scan(&classID); err != nil {
		return -1, e.StacktraceError(err)
	}
	return classID, nil
}

func (s *Store) CreateStudent(in *proto.CreateStudentRequest) (int, error) {
	var studID int = -1
	err := s.db.QueryRow(
		context.Background(),
		`INSERT INTO students (name, socialType) VALUES ($1, $2) RETURNING id;`,
		in.Name, in.Type,
	).Scan(&studID)
	if err != nil {
		return -1, e.StacktraceError(err)
	}
	return studID, nil
}

func (s *Store) CreateChat(in *proto.CreateChatRequest) (int, error) {
	var id int = -1
	var teacherID int = -1
	row := s.db.QueryRow(
		context.Background(),
		`SELECT teacherID FROM classes WHERE id = $1;`,
		in.ClassID,
	)
	if err := row.Scan(&teacherID); err != nil {
		return -1, e.StacktraceError(err)
	}

	err := s.db.QueryRow(
		context.Background(),
		`INSERT INTO chats (teacherID, studentID, classID) VALUES ($1, $2, $3) RETURNING id;`,
		teacherID, in.StudentID, in.ClassID,
	).Scan(&id)
	if err != nil {
		return -1, e.StacktraceError(err)
	}
	return id, nil
}

func (s *Store) GetHomeworksByChatID(classID int) ([]*proto.HomeworkData, error) {
	hws := []*proto.HomeworkData{}
	rows, err := s.db.Query(
		context.Background(),
		`SELECT (id, title, description, file) FROM homeworks WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()
	for rows.Next() {
		tmp := proto.HomeworkData{}
		tmpFileString := ""
		if err := rows.Scan(&tmp.HomeworkID, &tmp.Title, &tmp.Description, &tmpFileString); err != nil {
			return nil, e.StacktraceError(err)
		}
		tmp.AttachmentURLs = append(tmp.AttachmentURLs, tmpFileString)
		hws = append(hws, &tmp)
	}
	return hws, nil
}

func (s *Store) CreateSolution(in *proto.SendSolutionRequest) error {
	_, err := s.db.Exec(
		context.Background(),
		`INSERT INTO solutions (hwID, studentID, text, time, file) VALUES ($1, $2, $3, $4, $5);`,
		in.HomeworkID, in.StudentID, in.Solution.Text, time.Now().Format("2006.01.02 15:04:05"), in.Solution.AttachmentURLs[0],
	)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}
