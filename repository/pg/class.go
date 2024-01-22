package pg

import (
	e "main/domain/errors"
	m "main/domain/model"

	"github.com/lib/pq"
)

func (s *PostgreSqlStore) ValidateToken(tok string) (int, error) {
	var classID int = -1
	row := s.db.QueryRow(
		`SELECT id FROM classes WHERE inviteToken = $1;`, tok)
	if err := row.Scan(&classID); err != nil {
		return -1, e.StacktraceError(err)
	}
	return classID, nil
}

func (s *PostgreSqlStore) GetTasksByHomeworkID(homeworkID int) ([]m.Task, error) {
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

func (s *PostgreSqlStore) GetHomeworksByChatID(classID int) ([]m.Homework, error) {
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

func (s *PostgreSqlStore) CreateSolution(sol *m.CreateSolution) error {
	if sol.AttachmentURLs == nil {
		sol.AttachmentURLs = []string{}
	}

	_, err := s.db.Exec(
		`INSERT INTO solutions (homeworkID, studentID, text, createTime, files) VALUES ($1, $2, $3, $4, $5);`,
		sol.HomeworkID, sol.StudentID, sol.Text, sol.CreateDate, (*pq.StringArray)(&sol.AttachmentURLs),
	)
	if err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (s *PostgreSqlStore) GetTeacherIDByClassID(classID int) (int, error) {
	tID := 0
	row := s.db.QueryRow(`SELECT teacherID FROM classes WHERE id = $1;`, classID)
	if err := row.Scan(&tID); err != nil {
		return -1, e.StacktraceError(err)
	}
	return tID, nil
}
