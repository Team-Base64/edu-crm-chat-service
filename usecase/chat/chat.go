package chat

import (
	d "main/delivery"
	rep "main/repository"
	uc "main/usecase"
)

type ChatUsecase struct {
	hub             d.HubInterface
	dataStore       rep.DataStoreInterface
	fileStore       rep.FileStoreInterface
	calendar        d.CalendarInterface
	filestoragePath string
	urlDomain       string
}

func NewChatUsecase(
	hud d.HubInterface,
	store rep.DataStoreInterface,
	fs rep.FileStoreInterface,
	calendar d.CalendarInterface,
	fsPath string,
	urlDomain string,
) uc.UsecaseInterface {
	return &ChatUsecase{
		hub:             hud,
		dataStore:       store,
		fileStore:       fs,
		calendar:        calendar,
		filestoragePath: fsPath,
		urlDomain:       urlDomain,
	}
}
