package repository

import (
	m "main/domain/model"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type DataStoreInterface interface {
	// CHAT
	CreateChat(chat *m.Chat) (int, int, error)
	AddMessage(msg *m.CreateMessage) error
	GetChatsByClassID(classID int) ([]int, error)
	GetSocialTypeByChatID(chatID int) (string, error)
	// CLASS
	ValidateToken(tok string) (int, error)
	GetTasksByHomeworkID(homeworkID int) ([]m.Task, error)
	GetHomeworksByChatID(classID int) ([]m.Homework, error)
	CreateSolution(sol *m.CreateSolution) error
	// STUDENT
	CreateStudent(student *m.Student) (int, error)
	// TEACHER
	GetTeacherLoginById(id int) (string, error)
	GetTeacherLoginByChatId(id int) (string, error)
	GetTeacherLoginByHomeworkId(hwid int) (string, error)
	GetTeacherIDByClassID(classID int) (int, error)
}

type FileStoreInterface interface {
	UploadFile(file *m.Attach) (string, error)
}
