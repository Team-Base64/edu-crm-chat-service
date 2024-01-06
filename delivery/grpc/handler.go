package grpc

import (
	"context"
	"io"
	"log"

	proto "main/delivery/grpc/proto"
	e "main/domain/errors"
	m "main/domain/model"
	u "main/usecase"
)

type ChatGrpcHandler struct {
	proto.UnimplementedBotChatServer
	uc u.UsecaseInterface
}

func NewChatGrpcHander(uc u.UsecaseInterface) proto.BotChatServer {
	return &ChatGrpcHandler{
		uc: uc,
	}
}

func (h *ChatGrpcHandler) StartChatVK(ch proto.BotChat_StartChatVKServer) error {
	log.Println("start chat vk")
	defer log.Println("end chat vk")
	var errSending error = nil
	go func() {
		for {
			// отправка из вебсокета в бота
			msg, err := h.uc.GetMsgForTG()

			if err != nil {
				errSending = e.StacktraceError(err)
				log.Println(errSending)
				break
			}

			if msg == nil {
				continue
			}

			log.Println("preparing msg to vk bot: ", msg)

			resp := proto.Message{
				Text:           msg.Text,
				ChatID:         int32(msg.ChatID),
				AttachmentURLs: msg.AttachmentURLs,
			}
			if err := ch.Send(&resp); err != nil {
				errSending = e.StacktraceError(err)
				log.Println(errSending)
				break
			}
		}
	}()
	for {
		//приём сообщений от бота
		if errSending != nil {
			log.Println(e.StacktraceError(errSending))
			return errSending
		}

		log.Println("waiting for msg from vk bot: ")

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

		if err = h.uc.SendMsgToClient(&m.Message{
			ChatID:         int(req.ChatID),
			Text:           req.Text,
			AttachmentURLs: req.AttachmentURLs,
		}, "vk"); err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}
	}
}

func (h *ChatGrpcHandler) StartChatTG(ch proto.BotChat_StartChatTGServer) error {
	log.Println("start chat tg")

	defer log.Println("end chat tg")
	var errSending error = nil

	go func() {
		for {
			// отправка из вебсокета в бота
			msg, err := h.uc.GetMsgForTG()

			if err != nil {
				errSending = e.StacktraceError(err)
				log.Println(errSending)
				break
			}

			if msg == nil {
				continue
			}

			log.Println("preparing msg to tg bot: ", msg)

			resp := proto.Message{
				Text:           msg.Text,
				ChatID:         int32(msg.ChatID),
				AttachmentURLs: msg.AttachmentURLs,
			}
			if err := ch.Send(&resp); err != nil {
				errSending = e.StacktraceError(err)
				log.Println(errSending)
				break
			}
		}
	}()

	for {
		//приём сообщений от бота
		if errSending != nil {
			log.Println(e.StacktraceError(errSending))
			return errSending
		}
		log.Println("waiting for msg from tg bot: ")

		req, err := ch.Recv()
		if err == io.EOF {
			log.Println(e.StacktraceError(err))
			return err
		}
		if err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}
		log.Println("received msg from tg bot: ", req)

		if err = h.uc.SendMsgToClient(&m.Message{
			ChatID:         int(req.ChatID),
			Text:           req.Text,
			AttachmentURLs: req.AttachmentURLs,
		}, "tg"); err != nil {
			log.Println(e.StacktraceError(err))
			return err
		}
	}
}

func (h *ChatGrpcHandler) UploadFile(ctx context.Context, req *proto.FileUploadRequest) (*proto.FileUploadResponse, error) {
	log.Println("called grpc UploadFile", req.Mimetype, req.FileURL)
	fileAddr, err := h.uc.LoadFile(req.Mimetype, req.FileURL, "chat")

	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.FileUploadResponse{InternalFileURL: ""}, err
	}
	return &proto.FileUploadResponse{InternalFileURL: fileAddr}, nil
}

func (h *ChatGrpcHandler) BroadcastMsg(ctx context.Context, req *proto.BroadcastMessage) (*proto.Nothing, error) {
	log.Println("called Broadcast Msg from main backend")

	if err := h.uc.SendBroadcastMsg(&m.BroadcastMessage{
		ChatID:         int(req.ClassID),
		Title:          req.Title,
		Description:    req.Description,
		AttachmentURLs: req.AttachmentURLs,
	}); err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}

	return &proto.Nothing{}, nil
}

func (h *ChatGrpcHandler) ValidateToken(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	log.Println("called ValidateToken " + req.Token)

	id, err := h.uc.ValidateToken(req.Token)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.ValidateTokenResponse{ClassID: -1}, err
	}

	return &proto.ValidateTokenResponse{ClassID: int32(id)}, nil
}

func (h *ChatGrpcHandler) CreateStudent(ctx context.Context, req *proto.CreateStudentRequest) (*proto.CreateStudentResponse, error) {
	log.Println("called CreateStudent "+req.Name, req.Type)

	id, err := h.uc.CreateStudent(&m.Student{
		Name:      req.Name,
		Type:      req.Type,
		AvatarURL: req.AvatarURL,
	})
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateStudentResponse{StudentID: -1}, err
	}

	return &proto.CreateStudentResponse{StudentID: int32(id)}, nil
}

func (h *ChatGrpcHandler) CreateChat(ctx context.Context, req *proto.CreateChatRequest) (*proto.CreateChatResponse, error) {
	log.Println("called CreateChat ")

	chatID, err := h.uc.CreateChat(&m.Chat{
		ClassID:   int(req.ClassID),
		StudentID: int(req.StudentID),
	})
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.CreateChatResponse{InternalChatID: -1}, err
	}

	return &proto.CreateChatResponse{InternalChatID: int32(chatID)}, nil
}

func (h *ChatGrpcHandler) GetHomeworks(ctx context.Context, req *proto.GetHomeworksRequest) (*proto.GetHomeworksResponse, error) {
	log.Println("called GetHomeworks ")

	hws, err := h.uc.GetHomeworks(int(req.ClassID))
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.GetHomeworksResponse{Homeworks: nil}, err
	}
	protoHws := []*proto.HomeworkData{}
	for _, hw := range hws {
		protoHw := proto.HomeworkData{
			HomeworkID:   int32(hw.HomeworkID),
			Title:        hw.Title,
			Description:  hw.Description,
			CreateDate:   hw.CreateDate.String(),
			DeadlineDate: hw.DeadlineDate.String(),
			Tasks:        []*proto.TaskData{},
		}

		for _, task := range hw.Tasks {
			protoTask := proto.TaskData{
				Description:    task.Description,
				AttachmentURLs: task.AttachmentURLs,
			}
			protoHw.Tasks = append(protoHw.Tasks, &protoTask)
		}

		protoHws = append(protoHws, &protoHw)
	}

	return &proto.GetHomeworksResponse{Homeworks: protoHws}, nil
}

func (h *ChatGrpcHandler) SendSolution(ctx context.Context, req *proto.SendSolutionRequest) (*proto.Nothing, error) {
	log.Println("called SendSolution ", req.HomeworkID, req.Solution.Text, req.Solution.AttachmentURLs)
	if err := h.uc.SendSolution(&m.Solution{
		HomeworkID:     int(req.HomeworkID),
		Text:           req.Solution.Text,
		AttachmentURLs: req.Solution.AttachmentURLs,
	}); err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.Nothing{}, err
	}

	return &proto.Nothing{}, nil
}

func (h *ChatGrpcHandler) SendNotification(ctx context.Context, req *proto.Message) (*proto.Nothing, error) {
	log.Println("called Send Msg from main backend")

	if err := h.uc.SendNotification(&m.Message{
		ChatID:         int(req.ChatID),
		Text:           req.Text,
		AttachmentURLs: req.AttachmentURLs,
	}); err != nil {
		return &proto.Nothing{}, e.StacktraceError(err)
	}

	return &proto.Nothing{}, nil
}

func (h *ChatGrpcHandler) GetEvents(ctx context.Context, req *proto.GetEventsRequest) (*proto.GetEventsResponse, error) {
	log.Println("called GetEvents from bots")
	events, err := h.uc.GetEvents(int(req.ClassID))
	if err != nil {
		log.Println(e.StacktraceError(err))
		return &proto.GetEventsResponse{}, err
	}

	protoEvents := []*proto.EventData{}
	for _, event := range events {
		protoEvents = append(protoEvents, &proto.EventData{
			Id:          event.ID,
			Title:       event.Title,
			StartDate:   event.StartDate.String(),
			EndDate:     event.EndDate.String(),
			Description: event.Description,
			ClassID:     int32(event.ClassID),
		})
	}

	return &proto.GetEventsResponse{Events: protoEvents}, nil
}
