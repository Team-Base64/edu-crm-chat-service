package pg

import (
	e "main/domain/errors"
	m "main/domain/model"
	rep "main/repository"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) rep.StoreInterface {
	return &Store{
		db: db,
	}
}

func (s *Store) CheckSession(in string) (string, error) {
	teacherLogin := ""
	if err := s.db.QueryRow(`SELECT teacherLogin FROM sessions WHERE id = $1;`, in).Scan(&teacherLogin); err != nil {
		return "", e.StacktraceError(err)
	}
	return teacherLogin, nil
}

func (s *Store) AddMessage(msg *m.CreateMessage) error {
	_, err := s.db.Exec(`INSERT INTO messages (chatID, text, isAuthorTeacher, createtime, isRead, attaches) VALUES ($1, $2, $3, $4, $5, $6);`,
		msg.ChatID, msg.Text, msg.IsAuthorTeacher, msg.CreateTime, msg.IsRead, (*pq.StringArray)(&msg.AttachmentURLs))
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) GetChatsByClassID(classID int) ([]int, error) {
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

func (s *Store) GetTypeByChatID(chatID int) (string, error) {
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

func (s *Store) ValidateToken(tok string) (int, error) {
	var classID int = -1
	row := s.db.QueryRow(
		`SELECT id FROM classes WHERE inviteToken = $1;`, tok)
	if err := row.Scan(&classID); err != nil {
		return -1, e.StacktraceError(err)
	}
	return classID, nil
}

func (s *Store) CreateStudent(student *m.Student) (int, error) {
	var studID int = -1
	err := s.db.QueryRow(
		`INSERT INTO students (name, socialType, avatar) VALUES ($1, $2, $3) RETURNING id;`,
		student.Name, student.Type, student.AvatarURL,
	).Scan(&studID)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	return studID, nil
}

func (s *Store) GetTeacherLoginById(id int) (string, error) {
	login := ""
	row := s.db.QueryRow(
		`SELECT login FROM teachers WHERE id = $1;`,
		id,
	)
	if err := row.Scan(&login); err != nil {
		return "", e.StacktraceError(err)
	}
	return login, nil
}

func (s *Store) GetTeacherLoginByChatId(id int) (string, error) {
	teacherId := 0
	row := s.db.QueryRow(
		`SELECT teacherID FROM chats WHERE id = $1;`,
		id,
	)
	if err := row.Scan(&teacherId); err != nil {
		return "", e.StacktraceError(err)
	}
	return s.GetTeacherLoginById(teacherId)
}

func (s *Store) CreateChat(chat *m.Chat) (int, int, error) {
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

func (s *Store) GetTeacherLoginByHomeworkId(hwid int) (string, error) {
	classId := 0
	row := s.db.QueryRow(
		`SELECT classID FROM homeworks WHERE id = $1;`,
		hwid,
	)
	if err := row.Scan(&classId); err != nil {
		return "", e.StacktraceError(err)
	}
	tId := 0
	row = s.db.QueryRow(
		`SELECT teacherID FROM classes WHERE id = $1;`,
		classId,
	)
	if err := row.Scan(&tId); err != nil {
		return "", e.StacktraceError(err)
	}
	return s.GetTeacherLoginById(tId)
}

func (s *Store) GetTasksByHomeworkID(homeworkID int) ([]m.Task, error) {
	rows, err := s.db.Query(
		`SELECT t.description, t.attaches
		 FROM tasks t
		 JOIN homeworks_tasks ht ON t.id = ht.taskID
		 WHERE ht.homeworkID = $1
		 ORDER BY ht.rank;`,
		homeworkID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()

	tasks := []m.Task{}
	for rows.Next() {
		var task m.Task
		err := rows.Scan(&task.Description, (*pq.StringArray)(&task.AttachmentURLs))
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *Store) GetHomeworksByChatID(classID int) ([]m.Homework, error) {
	hws := []m.Homework{}
	rows, err := s.db.Query(
		`SELECT id, title, description, createTime, deadlineTime
		 FROM homeworks WHERE classID = $1;`,
		classID,
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}
	defer rows.Close()
	for rows.Next() {
		tmp := m.Homework{}
		if err := rows.Scan(
			&tmp.HomeworkID, &tmp.Title, &tmp.Description,
			&tmp.CreateDate, &tmp.DeadlineDate,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		tmp.Tasks, err = s.GetTasksByHomeworkID(tmp.HomeworkID)
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		hws = append(hws, tmp)
	}
	return hws, nil
}

func (s *Store) CreateSolution(sol *m.Solution) error {
	_, err := s.db.Exec(
		`INSERT INTO solutions (homeworkID, studentID, text, createTime, files) VALUES ($1, $2, $3, $4, $5);`,
		sol.HomeworkID, sol.StudentID, sol.Text, sol.CreateDate, (*pq.StringArray)(&sol.AttachmentURLs),
	)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) GetAllUserChatIDs(teacherLogin string) ([]int32, error) {
	var teacherID int = -1
	tIDs := []int32{}
	row := s.db.QueryRow(`SELECT id FROM teachers WHERE login = $1;`, teacherLogin)
	if err := row.Scan(&teacherID); err != nil {
		return []int32{}, e.StacktraceError(err)
	}

	rows, err := s.db.Query(`SELECT id FROM chats WHERE teacherID = $1;`, teacherID)
	if err != nil {
		return []int32{}, e.StacktraceError(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tmp int32 = 0
		if err := rows.Scan(&tmp); err != nil {
			return []int32{}, e.StacktraceError(err)
		}
		tIDs = append(tIDs, tmp)
	}
	return tIDs, nil
}

func (s *Store) GetTeacherIDByClassID(classID int) (int, error) {
	tID := 0
	row := s.db.QueryRow(`SELECT teacherID FROM classes WHERE id = $1;`, classID)
	if err := row.Scan(&tID); err != nil {
		return -1, e.StacktraceError(err)
	}
	return tID, nil
}
