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
	//file := NewFile()
	//var fileSize uint32
	//fileSize := 0
	//url := "https://api.telegram.org/file/bot1290980811:AAEgopVWqb7o0I72cwdIGGZRsRyE0GGNkLA/photos/file_2285.jpg"

	// req, err := stream.Recv()
	// if err == io.EOF {
	// 	log.Println("exit upload stream")
	// 	return err
	// }
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }

	// s := string(req.ChatID) + time.Now().Format("2006.01.02 15:04:05")
	// h := sha256.New()
	// h.Write([]byte(s))
	// homeworkNum := hex.EncodeToString(h.Sum(nil))

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

	fileName := "./filestorage/homeworks/homework_" + homeworkNum + fileExt
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error create/open file")
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	defer f.Close()

	//resp, err := http.Get("https://api.telegram.org/file/bot1290980811:AAEgopVWqb7o0I72cwdIGGZRsRyE0GGNkLA/photos/file_2285.jpg")
	log.Println(req.FileURL)
	resp, err := http.Get(req.FileURL)
	if err != nil {
		log.Println(err)
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	defer resp.Body.Close()

	n, err := io.Copy(f, resp.Body)

	// _, err = f.Write(req.GetChunk())
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// fileSize += len(req.GetChunk())

	// for {
	// 	req, err := stream.Recv()
	// 	if err == io.EOF {
	// 		log.Println("error: end of file")
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	chunk := req.GetChunk()
	// 	if _, err := f.Write(chunk); err != nil {
	// 		log.Println(err)
	// 		return err
	// 	}
	// 	fileSize += len(chunk)
	// }
	log.Println("saved file:", fileName, "size: ", n)

	fileAddr := "http://127.0.0.1:8081/filestorage/homeworks/homework_" + homeworkNum + fileExt
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
	//return stream.SendAndClose(&proto.FileUploadResponse{FileName: "homework_" + homeworkNum + fileExt, Size: uint32(fileSize)})
}

func (sm *ChatManager) BroadcastMsg(ctx context.Context, req *proto.BroadcastMessage) (*proto.Nothing, error) {
	log.Println("called Broadcast Msg from main backend")
	ids, err := sm.store.GetChatsByClassID(int(req.ClassID))
	if err != nil {
		log.Println(err)
		return &proto.Nothing{}, err
	}
	for _, id := range *ids {
		sm.hub.MessagesToTGBot <- &m.MessageWebsocket{ChatID: int32(id), Text: req.Title + "\n" + req.Description, AttachmentURLs: req.AttachmentURLs}
	}

	return &proto.Nothing{}, nil
}
