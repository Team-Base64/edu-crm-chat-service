package repository

import (
	m "main/domain/model"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type StoreInterface interface {
	CheckSession(in string) (string, error)
	AddMessage(msg *m.CreateMessage) error
	GetChatsByClassID(classID int) ([]int, error)
	GetTypeByChatID(chatID int) (string, error)
	ValidateToken(tok string) (int, error)
	CreateStudent(student *m.Student) (int, error)
	GetTeacherLoginById(id int) (string, error)
	GetTeacherLoginByChatId(id int) (string, error)
	CreateChat(chat *m.Chat) (int, int, error)
	GetTeacherLoginByHomeworkId(hwid int) (string, error)
	GetTasksByHomeworkID(homeworkID int) ([]m.Task, error)
	GetHomeworksByChatID(classID int) ([]m.Homework, error)
	CreateSolution(sol *m.Solution) error
	GetAllUserChatIDs(teacherLogin string) ([]int32, error)
	GetTeacherIDByClassID(classID int) (int, error)
}
