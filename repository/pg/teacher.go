package pg

import e "main/domain/errors"

func (s *PostgreSqlStore) GetTeacherLoginById(id int) (string, error) {
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

func (s *PostgreSqlStore) GetTeacherLoginByChatId(id int) (string, error) {
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

func (s *PostgreSqlStore) GetTeacherLoginByHomeworkId(hwid int) (string, error) {
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
