package chat

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	m "main/domain/model"

	//chat "main/src"

	proto "main/src/proto"

	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type ChatManager struct {
	proto.UnimplementedBotChatServer
	store StoreInterface
	hub   *Hub
}

func NewChatManager(store StoreInterface, hub *Hub) *ChatManager {
	return &ChatManager{
		store: store,
		hub:   hub,
	}
}

func (sm *ChatManager) StartChatTG(ch proto.BotChat_StartChatTGServer) error {
	log.Println("start chat tg")
	var mockChatID int32 = 1
	defer log.Println("end chat tg")
	errSending := errors.New("Empty")
	//defer return errors.New("GRPC Consume: message channel closed")

	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToTGBot
			if mes2.Text == "" {
				continue
			}
			log.Println("preparing mes to tg bot: ", mes2)
			resp := proto.Message{Text: mes2.Text, ChatID: mockChatID, AttachmentURLs: mes2.AttachmentURLs}
			if err := ch.Send(&resp); err != nil {
				log.Println("1!!!!!!!!! error: ", err)
				// if err.Error() == "rpc error: code = Canceled desc = context canceled" {
				// 	log.Println("breaking grpc stream")
				// }
				log.Println("breaking grpc stream")
				errSending = err
				break
				//continue
				//return err
			}
			// err := sm.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true, IsRead: true})
			// if err != nil {
			// 	log.Println(err)
			// 	errSending = err
			// 	break
			// }
		}
	}()
	for {
		if errSending.Error() != "Empty" {
			log.Println("err on sending goroutine: ", errSending)
			return errSending
		}

		log.Println("waiting for mes from tg bot: ")
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println("exit tg stream")
			return err
		}
		if err != nil {
			log.Println("2!!!!!!!!! error: ", err)
			if err.Error() == "rpc error: code = Canceled desc = context canceled" {
				log.Println("breaking grpc stream")
				return err
			}
			//continue
			//break
			return err
		}
		log.Println("received mes from tg bot: ", req)
		mes := m.MessageWebsocket{Text: req.Text, ChatID: mockChatID, Channel: "chat"}
		if sm.hub.chats[mes.ChatID] != nil {
			log.Println("routing mes from tg bot to hub: ", req)
			sm.hub.Broadcast <- &mes
			log.Println("routing mes from tg bot to hub + added to broadcast: ", req)
			err = sm.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false, IsRead: false})
			if err != nil {
				log.Println(err)
				return err
			}
			log.Println("routing mes from tg bot to hub + added to db: ", req)
		}

		//break
	}
	//return nil
	//return errors.New("GRPC Consume: message channel closed")
}

// func (sm *ChatManager) StartChatVK(ch proto.BotChat_StartChatVKServer) error {
// 	log.Println("start chat vk")
// 	defer log.Println("close chat vk")
// 	var mockChatID int32 = 2
// 	go func() {
// 		for {
// 			// отправка из вебсокета в бота
// 			mes2 := <-sm.hub.MessagesToVKBot
// 			resp := proto.Message{Text: mes2.Text, ChatID: mockChatID}
// 			if err := ch.Send(&resp); err != nil {
// 				log.Println(err)
// 				if err.Error() == "rpc error: code = Canceled desc = context canceled" {
// 					log.Println("breaking grpc stream")
// 					break
// 					//return nil
// 				}
// 				continue
// 			}
// 			log.Println("preparing mes to vk bot: ", mes2)
// 			sm.store.AddMessage(&m.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true, IsRead: true})
// 		}
// 	}()
// 	for {
// 		//приём сообщений от бота
// 		req, err := ch.Recv()
// 		if err == io.EOF {
// 			log.Println("exit vk stream")
// 			return nil
// 		}
// 		if err != nil {
// 			log.Println(err)
// 			if err.Error() == "rpc error: code = Canceled desc = context canceled" {
// 				log.Println("breaking grpc stream")
// 				return nil
// 			}
// 			//continue
// 			return nil
// 		}
// 		log.Println("received mes from vk bot: ", req)
// 		mes := m.MessageWebsocket{Text: req.Text, ChatID: mockChatID, Channel: "chat"}
// 		sm.hub.Broadcast <- &mes
// 		sm.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false, IsRead: false})
// 	}
// }

func (sm *ChatManager) UploadFile(ctx context.Context, req *proto.FileUploadRequest) (*proto.FileUploadResponse, error) {
	log.Println("called UploadFile")

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

	fileName := "./filestorage/chat/attach_" + homeworkNum + fileExt
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error create/open file")
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	defer f.Close()

	//resp, err := http.Get("https://api.telegram.org/file/bot1290980811:AAEgopVWqb7o0I72cwdIGGZRsRyE0GGNkLA/photos/file_2285.jpg")
	//resp, err := http.Get("https://vk.com/doc211427710_672529050?hash=mh0QoNOXWeSDKQqSdLmQcvzYGYlZX0BtMSHIE1L9hwg&dl=KGB2X4QArOzAZEZWxZuYjwrx9RbVVrf2FkTMZ8hHklH&api=1&no_preview=1")
	log.Println(req.FileURL)
	resp, err := http.Get(req.FileURL)
	if err != nil {
		log.Println(err)
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	defer resp.Body.Close()

	n, err := io.Copy(f, resp.Body)
	log.Println("saved file:", fileName, "size: ", n)

	fileAddr := "http://127.0.0.1:8081/filestorage/chat/attach_" + homeworkNum + fileExt
	// mes := m.MessageWebsocket{Text: req.Text + "\n" + fileAddr, ChatID: 1, Channel: "chat"}
	// if sm.hub.chats[mes.ChatID] != nil {
	// 	log.Println("routing mes with attach from tg bot to hub: ", req)
	// 	sm.hub.Broadcast <- &mes
	// 	log.Println("routing mes with attach from tg bot to hub + added to broadcast: ", req)
	// 	err = sm.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false, IsRead: false})
	// 	if err != nil {
	// 		log.Println(err)
	// 		return &proto.Status{IsSuccessful: false}, err
	// 	}
	// 	log.Println("routing mes with attach from tg bot to hub + added to db: ", req)
	// }
	return &proto.FileUploadResponse{InternalFileURL: fileAddr}, nil
}

func (sm *ChatManager) BroadcastMsg(ctx context.Context, req *proto.BroadcastMessage) (*proto.Nothing, error) {
	log.Println("called Broadcast Msg from main backend")
	ids, err := sm.store.GetChatsByClassID(int(req.ClassID))
	if err != nil {
		log.Println(err)
		return &proto.Nothing{}, err
	}
	// ИСПРАВИТЬ!!!
	for _, id := range *ids {
		sm.hub.MessagesToTGBot <- &m.MessageWebsocket{ChatID: int32(id), Text: req.Title + "\n" + req.Description, AttachmentURLs: req.AttachmentURLs}
		sm.hub.MessagesToVKBot <- &m.MessageWebsocket{ChatID: int32(id), Text: req.Title + "\n" + req.Description, AttachmentURLs: req.AttachmentURLs}
	}

	return &proto.Nothing{}, nil
}

func (sm *ChatManager) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	log.Println("called ValidateToken " + req.Token)
	id, err := sm.store.ValidateToken(req.Token)
	if err != nil {
		log.Println(err)
		return &proto.ValidateTokenResponse{ClassID: -1}, err
	}
	return &proto.ValidateTokenResponse{ClassID: int32(id)}, nil
}

func (sm *ChatManager) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateStudentResponse, error) {
	log.Println("called CreateStudent " + req.Name)
	id, err := sm.store.CreateStudent(req)
	if err != nil {
		log.Println(err)
		return &proto.CreateStudentResponse{StudentID: -1}, err
	}
	return &proto.CreateStudentResponse{StudentID: int32(id)}, nil
}

func (sm *ChatManager) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {
	log.Println("called CreateChat ")
	id, err := sm.store.CreateChat(req)
	if err != nil {
		log.Println(err)
		return &proto.CreateChatResponse{InternalChatID: -1}, err
	}
	return &proto.CreateChatResponse{InternalChatID: int32(id)}, nil
}

func (sm *ChatManager) GetHomeworks(ctx context.Context, req *proto.GetHomeworksRequest) (*proto.GetHomeworksResponse, error) {
	log.Println("called GetHomeworks ")
	hws, err := sm.store.GetHomeworksByChatID(int(req.ClassID))
	if err != nil {
		log.Println(err)
		return &proto.GetHomeworksResponse{Homeworks: nil}, err
	}
	return &proto.GetHomeworksResponse{Homeworks: hws}, nil
}

func (sm *ChatManager) SendSolution(ctx context.Context, req *proto.SendSolutionRequest) (*proto.SendSolutionResponse, error) {
	log.Println("called GetHomeworks ")
	err := sm.store.CreateSolution(req)
	if err != nil {
		log.Println(err)
		return &proto.SendSolutionResponse{}, err
	}
	return &proto.SendSolutionResponse{}, nil
}
