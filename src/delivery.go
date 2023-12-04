package chat

import (
	"encoding/json"
	"net/http"

	e "main/domain/errors"
	"main/domain/model"
)

// @title TCRA API
// @version 1.0
// @description EDUCRM back chat server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081
// @BasePath  /apichat

type Handler struct {
	store StoreInterface
	hub   *Hub
}

func NewHandler(store StoreInterface, hub *Hub) *Handler {
	return &Handler{
		store: store,
		hub:   hub,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}
