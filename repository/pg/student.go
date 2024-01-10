package pg

import (
	e "main/domain/errors"
	m "main/domain/model"
)

func (s *PostgreSqlStore) CreateStudent(student *m.Student) (int, error) {
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
