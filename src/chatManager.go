package chat

import (
	"context"
	"io"
	"log"

	"sync"
	"time"

	"main/model"
	//chat "main/src"

	proto "main/src/proto"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatManager struct {
	proto.UnimplementedBotChatServer
	db  *pgxpool.Pool
	mu  sync.RWMutex
	hub *Hub
}

func NewChatManager(db *pgxpool.Pool, hub *Hub) *ChatManager {
	return &ChatManager{
		mu:  sync.RWMutex{},
		db:  db,
		hub: hub,
	}
}

func (sm *ChatManager) AddMessage(in *model.CreateMessage) error {
	_, err := sm.db.Query(context.Background(), `INSERT INTO messages (chatID, text, isAuthorTeacher, time) VALUES ($1, $2, $3, $4);`, in.ChatID, in.Text, in.IsAuthorTeacher, time.Now().Format("2006.01.02 15:04:05"))
	if err != nil {
		return err
	}
	return nil
}

func (sm *ChatManager) StartChatTG(ch proto.BotChat_StartChatTGServer) error {
	log.Println("start chat tg")
	var mockChatID int32 = 1
	go func() {
		for {
			// отправка из вебсокета в бота
			mes2 := <-sm.hub.MessagesToTGBot

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
			log.Println("preparing mes to tg bot: ", mes2)
			err := sm.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true})
			if err != nil {
				log.Println(err)
			}
		}
	}()
	for {
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			if err.Error() == "rpc error: code = Canceled desc = context canceled" {
				log.Println("breaking grpc stream")
				return nil
			}
			continue
		}
		log.Println("received mes from tg bot: ", req)
		mes := MessageWebsocket{Text: req.Text, ChatID: mockChatID, Channel: "chat"}
		sm.hub.Broadcast <- &mes
		sm.AddMessage(&model.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false})
	}
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
			sm.AddMessage(&model.CreateMessage{Text: resp.Text, ChatID: int(resp.ChatID), IsAuthorTeacher: true})
		}
	}()
	for {
		//приём сообщений от бота
		req, err := ch.Recv()
		if err == io.EOF {
			log.Println("exit")
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
		mes := MessageWebsocket{Text: req.Text, ChatID: mockChatID, Channel: "chat"}
		sm.hub.Broadcast <- &mes
		sm.AddMessage(&model.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: false})
	}
}
