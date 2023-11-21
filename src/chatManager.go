package chat

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	e "main/domain/errors"
	"main/domain/model"
	m "main/domain/model"

	proto "main/src/proto"

	"github.com/google/uuid"
)

type ChatManager struct {
	proto.UnimplementedBotChatServer
	store           StoreInterface
	hub             *Hub
	filestoragePath string
	urlDomain       string
}

func NewChatManager(store StoreInterface, hub *Hub, fp string, ud string) *ChatManager {
	return &ChatManager{
		store:           store,
		hub:             hub,
		filestoragePath: fp,
		urlDomain:       ud,
	}
}

func (sm *ChatManager) StartChatTG(ch proto.BotChat_StartChatTGServer) error {
	log.Println("start chat tg")

	defer log.Println("end chat tg")
	var errSending error = nil

	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToTGBot

			if mes2.Text == "" && mes2.AttachmentURLs == nil {
				continue
			}

			resp := proto.Message{Text: mes2.Text, ChatID: mes2.ChatID, AttachmentURLs: mes2.AttachmentURLs}
			if !mes2.IsSavedToDB {
				log.Println("writing mes to db: ", mes2)
				err := sm.store.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true, IsRead: true, AttachmentURLs: resp.AttachmentURLs})
				if err != nil {
					errSending = e.StacktraceError(err)
				  log.Println(errSending)
					break
				}
			}

			log.Println("preparing mes to tg bot: ", mes2)
			if err := ch.Send(&resp); err != nil {
				errSending = e.StacktraceError(err)
				log.Println(errSending)
				break
			}
		}
	}()
	for {
		if errSending != nil {
			log.Println(e.StacktraceError(errSending))
			return errSending
		}

		log.Println("waiting for mes from tg bot: ")
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println(e.StacktraceError(err))
			return err
		}
		if err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}
		log.Println("received mes from tg bot: ", req)
		mes := m.MessageWebsocket{Text: req.Text, ChatID: req.ChatID, Channel: "chat", AttachmentURLs: req.AttachmentURLs, CreateTime: time.Now(), IsSavedToDB: true}

		log.Println("writing mes to db: ", mes)
		err = sm.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false, IsRead: false, AttachmentURLs: mes.AttachmentURLs})
		if err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}

		if sm.hub.chats[mes.ChatID] != nil {
			sm.hub.Broadcast <- &mes
			log.Println("routing mes from tg bot to hub + added to broadcast: ", req)
		}
	}
}

func (sm *ChatManager) StartChatVK(ch proto.BotChat_StartChatVKServer) error {
	log.Println("start chat vk")
	defer log.Println("end chat vk")
	var errSending error = nil
	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToVKBot
			if mes2.Text == "" && mes2.AttachmentURLs == nil {
				continue
			}
			resp := proto.Message{Text: mes2.Text, ChatID: mes2.ChatID, AttachmentURLs: mes2.AttachmentURLs}

			if !mes2.IsSavedToDB {
				log.Println("writing mes to db: ", mes2)
				err := sm.store.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true, IsRead: true, AttachmentURLs: resp.AttachmentURLs})
				if err != nil {
					errSending = e.StacktraceError(err)
				  log.Println(errSending)
					break
				}
			}
			log.Println("preparing mes to vk bot: ", mes2)
			if err := ch.Send(&resp); err != nil {
				errSending = e.StacktraceError(err)
				log.Println(errSending)
				break
			}
		}
	}()
	for {
		if errSending.Error() != "Empty" {
			log.Println(e.StacktraceError(errSending))
			return errSending
		}

		log.Println("waiting for mes from vk bot: ")
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println(e.StacktraceError(err))
			return err
		}
		if err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}
		log.Println("received mes from vk bot: ", req)
		mes := m.MessageWebsocket{Text: req.Text, ChatID: req.ChatID, Channel: "chat", AttachmentURLs: req.AttachmentURLs, CreateTime: time.Now(), IsSavedToDB: true}

		log.Println("writing mes to db: ", mes)
		err = sm.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false, IsRead: false, AttachmentURLs: mes.AttachmentURLs})
		if err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}

		if sm.hub.chats[mes.ChatID] != nil {
			sm.hub.Broadcast <- &mes
			log.Println("routing mes from tg bot to hub + added to broadcast: ", req)
		}
	}
}

func (sm *ChatManager) UploadFile(ctx context.Context, req *proto.FileUploadRequest) (*proto.FileUploadResponse, error) {
	log.Println("called grpc UploadFile", req.Mimetype, req.FileURL)

	homeworkNum := uuid.New().String()

	fileExt := ""
	switch req.Mimetype {
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/png":
		fileExt = ".png"
	case "image/svg+xml":
		fileExt = ".svg"
	case "application/pdf":
		fileExt = ".pdf"
	default:
		err := errors.New("error: " + fileExt + " is not allowed file extension")
		log.Println(err)
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}

	fileName := sm.filestoragePath + "/chat/" + homeworkNum + fileExt
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	defer f.Close()

	resp, err := http.Get(req.FileURL)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	defer resp.Body.Close()

	n, err := io.Copy(f, resp.Body)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	log.Println("saved file:", fileName, "size: ", n)

	fileAddr := sm.urlDomain + "/filestorage/chat/" + homeworkNum + fileExt
	return &proto.FileUploadResponse{InternalFileURL: fileAddr}, nil
}

func (sm *ChatManager) BroadcastMsg(ctx context.Context, req *proto.BroadcastMessage) (*proto.Nothing, error) {
	log.Println("called Broadcast Msg from main backend")
	ids, err := sm.store.GetChatsByClassID(int(req.ClassID))
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}
	log.Println(" Broadcast for chats: ", ids)

	for _, id := range *ids {
		type1, err := sm.store.GetTypeByChatID(id)
		if err != nil {
			log.Println(e.StacktraceError(err), errors.New("chatid: "+strconv.Itoa(id)))
		}
		switch type1 {
		case "tg":
			sm.hub.MessagesToTGBot <- &m.MessageWebsocket{ChatID: int32(id), Text: req.Title + "\n" + req.Description, AttachmentURLs: req.AttachmentURLs, IsSavedToDB: false}
		case "vk":
			sm.hub.MessagesToVKBot <- &m.MessageWebsocket{ChatID: int32(id), Text: req.Title + "\n" + req.Description, AttachmentURLs: req.AttachmentURLs, IsSavedToDB: false}
		default:
		}
	}
	return &proto.Nothing{}, nil
}

func (sm *ChatManager) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	log.Println("called ValidateToken " + req.Token)
	id, err := sm.store.ValidateToken(req.Token)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.ValidateTokenResponse{ClassID: -1}, err
	}
	return &proto.ValidateTokenResponse{ClassID: int32(id)}, nil
}

func (sm *ChatManager) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateStudentResponse, error) {
	log.Println("called CreateStudent "+req.Name, req.Type)
	id, err := sm.store.CreateStudent(req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateStudentResponse{StudentID: -1}, err
	}
	return &proto.CreateStudentResponse{StudentID: int32(id)}, nil
}

func (sm *ChatManager) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {
	log.Println("called CreateChat ")
	id, err := sm.store.CreateChat(req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateChatResponse{InternalChatID: -1}, err
	}
	return &proto.CreateChatResponse{InternalChatID: int32(id)}, nil
}

func (sm *ChatManager) GetHomeworks(ctx context.Context, req *proto.GetHomeworksRequest) (*proto.GetHomeworksResponse, error) {
	log.Println("called GetHomeworks ")
	hws, err := sm.store.GetHomeworksByChatID(int(req.ClassID))
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.GetHomeworksResponse{Homeworks: nil}, err
	}
	return &proto.GetHomeworksResponse{Homeworks: hws}, nil
}

func (sm *ChatManager) SendSolution(ctx context.Context, req *proto.SendSolutionRequest) (*proto.SendSolutionResponse, error) {
	log.Println("called SendSolution ", req.HomeworkID, req.Solution.Text, req.Solution.AttachmentURLs)
	err := sm.store.CreateSolution(req)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.SendSolutionResponse{}, err
	}
	return &proto.SendSolutionResponse{}, nil
}

func (sm *ChatManager) SendMsg(ctx context.Context, req *proto.Message) (*proto.Nothing, error) {
	log.Println("called Send Msg from main backend")

	socialType, err := sm.store.GetTypeByChatID(int(req.ChatID))
	if err != nil {
		log.Println(e.StacktraceError(err), errors.New("chatid: "+strconv.Itoa(int(req.ChatID))))
		return &proto.Nothing{}, err
	}
	switch socialType {
	case "tg":
		sm.hub.MessagesToTGBot <- &m.MessageWebsocket{ChatID: req.ChatID, Text: req.Text, AttachmentURLs: req.AttachmentURLs}
	case "vk":
		sm.hub.MessagesToVKBot <- &m.MessageWebsocket{ChatID: req.ChatID, Text: req.Text, AttachmentURLs: req.AttachmentURLs}
	default:
	}

	return &proto.Nothing{}, nil
}
