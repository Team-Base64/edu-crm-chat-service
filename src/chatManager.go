package chat

import (
	"errors"
	"io"
	"log"

	m "main/domain/model"
	//chat "main/src"

	proto "main/src/proto"

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
			resp := proto.Message{Text: mes2.Text, ChatID: mockChatID}
			if err := ch.Send(&resp); err != nil {
				log.Println("!!!!!!!!!")
				log.Println(err)
				if err.Error() == "rpc error: code = Canceled desc = context canceled" {
					log.Println("breaking grpc stream")

					//break
					//return nil
				}
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
			log.Println("2!!!!!!!!!")
			log.Println(err)
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

func (sm *ChatManager) StartChatVK(ch proto.BotChat_StartChatVKServer) error {
	log.Println("start chat vk")
	defer log.Println("close chat vk")
	var mockChatID int32 = 2
	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToVKBot
			resp := proto.Message{Text: mes2.Text, ChatID: mockChatID}
			if err := ch.Send(&resp); err != nil {
				log.Println(err)
				if err.Error() == "rpc error: code = Canceled desc = context canceled" {
					log.Println("breaking grpc stream")
					break
					//return nil
				}
				continue
			}
			log.Println("preparing mes to vk bot: ", mes2)
			sm.store.AddMessage(&m.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true, IsRead: true})
		}
	}()
	for {
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println("exit vk stream")
			return nil
		}
		if err != nil {
			log.Println(err)
			if err.Error() == "rpc error: code = Canceled desc = context canceled" {
				log.Println("breaking grpc stream")
				return nil
			}
			//continue
			return nil
		}
		log.Println("received mes from vk bot: ", req)
		mes := m.MessageWebsocket{Text: req.Text, ChatID: mockChatID, Channel: "chat"}
		sm.hub.Broadcast <- &mes
		sm.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false, IsRead: false})
	}
}
