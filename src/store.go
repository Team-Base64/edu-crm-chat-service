package chat

import (
	"time"

	e "main/domain/errors"
	"main/domain/model"
	m "main/domain/model"
	proto "main/src/proto"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/lib/pq"
)

type StoreInterface interface {
	CheckSession(in string) (string, error)
	AddMessage(in *m.CreateMessage) error
	GetChatsByClassID(classID int) (*[]int, error)
	GetTypeByChatID(chatID int) (string, error)
	ValidateToken(tok string) (int, error)
	CreateStudent(in *proto.CreateStudentRequest) (int, error)
	GetTeacherLoginById(id int) (string, error)
	GetTeacherLoginByChatId(id int) (string, error)
	CreateChat(in *proto.CreateChatRequest) (int, int, error)
	GetTeacherLoginByHomeworkId(hwid int) (string, error)
	GetHomeworksByChatID(classID int) ([]*proto.HomeworkData, error)
	CreateSolution(in *proto.SendSolutionRequest) error
	GetAllUserChatIDs(teacherLogin string) ([]int32, error)

	GetTeacherIDByClassID(classID int) (int, error)
	GetCalendarDB(teacherID int) (*model.CalendarParams, error)
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) StoreInterface {
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

func (s *Store) AddMessage(in *m.CreateMessage) error {
	modelAtt := []string{}
	if in.AttachmentURLs != nil {
		modelAtt = in.AttachmentURLs
	}
	_, err := s.db.Exec(`INSERT INTO messages (chatID, text, isAuthorTeacher, createtime, isRead, attaches) VALUES ($1, $2, $3, $4, $5, $6);`,
		in.ChatID, in.Text, in.IsAuthorTeacher, time.Now().Format("2006.01.02 15:04:05"), in.IsRead, (*pq.StringArray)(&modelAtt))
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *Store) GetChatsByClassID(classID int) (*[]int, error) {
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
	return &ans, nil
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

func (s *Store) CreateStudent(in *proto.CreateStudentRequest) (int, error) {
	var studID int = -1
	err := s.db.QueryRow(
		`INSERT INTO students (name, socialType, avatar) VALUES ($1, $2, $3) RETURNING id;`,
		in.Name, in.Type, in.AvatarURL,
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

func (s *Store) CreateChat(in *proto.CreateChatRequest) (int, int, error) {
	var id int = -1
	var teacherID int = -1
	row := s.db.QueryRow(
		`SELECT teacherID FROM classes WHERE id = $1;`,
		in.ClassID,
	)
	if err := row.Scan(&teacherID); err != nil {
		return -1, -1, e.StacktraceError(err)
	}

	err := s.db.QueryRow(
		`INSERT INTO chats (teacherID, studentID, classID) VALUES ($1, $2, $3) RETURNING id;`,
		teacherID, in.StudentID, in.ClassID,
	).Scan(&id)
	if err != nil {
		return -1, -1, e.StacktraceError(err)
	}
	return teacherID, id, nil
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

func (s *Store) GetTasksByHomeworkID(homeworkID int) ([]*proto.TaskData, error) {
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

	tasks := []*proto.TaskData{}
	for rows.Next() {
		var task proto.TaskData
		err := rows.Scan(&task.Description, (*pq.StringArray)(&task.AttachmentURLs))
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (s *Store) GetHomeworksByChatID(classID int) ([]*proto.HomeworkData, error) {
	hws := []*proto.HomeworkData{}
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
		tmp := proto.HomeworkData{}
		var createTime, deadlineTime time.Time
		if err := rows.Scan(
			&tmp.HomeworkID, &tmp.Title, &tmp.Description,
			&createTime, &deadlineTime,
		); err != nil {
			return nil, e.StacktraceError(err)
		}

		tmp.CreateDate = createTime.String()
		tmp.DeadlineDate = deadlineTime.String()

		tmp.Tasks, err = s.GetTasksByHomeworkID(int(tmp.GetHomeworkID()))
		if err != nil {
			return nil, e.StacktraceError(err)
		}
		hws = append(hws, &tmp)
	}
	return hws, nil
}

func (s *Store) CreateSolution(in *proto.SendSolutionRequest) error {
	tmpAttach := []string{}
	if in.Solution.AttachmentURLs != nil {
		tmpAttach = in.Solution.AttachmentURLs
	}
	_, err := s.db.Exec(
		`INSERT INTO solutions (homeworkID, studentID, text, createTime, files) VALUES ($1, $2, $3, $4, $5);`,
		in.HomeworkID, in.StudentID, in.Solution.Text, time.Now().Format("2006.01.02 15:04:05"), (*pq.StringArray)(&tmpAttach),
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

func (s *Store) GetCalendarDB(teacherID int) (*model.CalendarParams, error) {
	ans := model.CalendarParams{}
	row := s.db.QueryRow(`SELECT id, idInGoogle FROM calendars WHERE teacherID = $1;`, teacherID)
	if err := row.Scan(&ans.ID, &ans.IDInGoogle); err != nil {
		return nil, e.StacktraceError(err)
	}
	return &ans, nil
}
