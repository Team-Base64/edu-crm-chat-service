// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: src/proto/chat.proto

// export PATH="$PATH:$(go env GOPATH)/bin"
// go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
// protoc --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative src/proto/chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BotChat_BroadcastMsg_FullMethodName  = "/chat.BotChat/BroadcastMsg"
	BotChat_StartChatTG_FullMethodName   = "/chat.BotChat/StartChatTG"
	BotChat_StartChatVK_FullMethodName   = "/chat.BotChat/StartChatVK"
	BotChat_UploadFile_FullMethodName    = "/chat.BotChat/UploadFile"
	BotChat_ValidateToken_FullMethodName = "/chat.BotChat/ValidateToken"
	BotChat_CreateChat_FullMethodName    = "/chat.BotChat/CreateChat"
	BotChat_GetHomeworks_FullMethodName  = "/chat.BotChat/GetHomeworks"
	BotChat_CreateStudent_FullMethodName = "/chat.BotChat/CreateStudent"
	BotChat_SendSolution_FullMethodName  = "/chat.BotChat/SendSolution"
)

// BotChatClient is the client API for BotChat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotChatClient interface {
	BroadcastMsg(ctx context.Context, in *BroadcastMessage, opts ...grpc.CallOption) (*Nothing, error)
	StartChatTG(ctx context.Context, opts ...grpc.CallOption) (BotChat_StartChatTGClient, error)
	StartChatVK(ctx context.Context, opts ...grpc.CallOption) (BotChat_StartChatVKClient, error)
	UploadFile(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error)
	ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
	CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error)
	GetHomeworks(ctx context.Context, in *GetHomeworksRequest, opts ...grpc.CallOption) (*GetHomeworksResponse, error)
	CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*CreateStudentResponse, error)
	SendSolution(ctx context.Context, in *SendSolutionRequest, opts ...grpc.CallOption) (*SendSolutionResponse, error)
}

type botChatClient struct {
	cc grpc.ClientConnInterface
}

func NewBotChatClient(cc grpc.ClientConnInterface) BotChatClient {
	return &botChatClient{cc}
}

func (c *botChatClient) BroadcastMsg(ctx context.Context, in *BroadcastMessage, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, BotChat_BroadcastMsg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botChatClient) StartChatTG(ctx context.Context, opts ...grpc.CallOption) (BotChat_StartChatTGClient, error) {
	stream, err := c.cc.NewStream(ctx, &BotChat_ServiceDesc.Streams[0], BotChat_StartChatTG_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &botChatStartChatTGClient{stream}
	return x, nil
}

type BotChat_StartChatTGClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type botChatStartChatTGClient struct {
	grpc.ClientStream
}

func (x *botChatStartChatTGClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *botChatStartChatTGClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *botChatClient) StartChatVK(ctx context.Context, opts ...grpc.CallOption) (BotChat_StartChatVKClient, error) {
	stream, err := c.cc.NewStream(ctx, &BotChat_ServiceDesc.Streams[1], BotChat_StartChatVK_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &botChatStartChatVKClient{stream}
	return x, nil
}

type BotChat_StartChatVKClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type botChatStartChatVKClient struct {
	grpc.ClientStream
}

func (x *botChatStartChatVKClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *botChatStartChatVKClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *botChatClient) UploadFile(ctx context.Context, in *FileUploadRequest, opts ...grpc.CallOption) (*FileUploadResponse, error) {
	out := new(FileUploadResponse)
	err := c.cc.Invoke(ctx, BotChat_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botChatClient) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	out := new(ValidateTokenResponse)
	err := c.cc.Invoke(ctx, BotChat_ValidateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botChatClient) CreateChat(ctx context.Context, in *CreateChatRequest, opts ...grpc.CallOption) (*CreateChatResponse, error) {
	out := new(CreateChatResponse)
	err := c.cc.Invoke(ctx, BotChat_CreateChat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botChatClient) GetHomeworks(ctx context.Context, in *GetHomeworksRequest, opts ...grpc.CallOption) (*GetHomeworksResponse, error) {
	out := new(GetHomeworksResponse)
	err := c.cc.Invoke(ctx, BotChat_GetHomeworks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botChatClient) CreateStudent(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*CreateStudentResponse, error) {
	out := new(CreateStudentResponse)
	err := c.cc.Invoke(ctx, BotChat_CreateStudent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botChatClient) SendSolution(ctx context.Context, in *SendSolutionRequest, opts ...grpc.CallOption) (*SendSolutionResponse, error) {
	out := new(SendSolutionResponse)
	err := c.cc.Invoke(ctx, BotChat_SendSolution_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotChatServer is the server API for BotChat service.
// All implementations must embed UnimplementedBotChatServer
// for forward compatibility
type BotChatServer interface {
	BroadcastMsg(context.Context, *BroadcastMessage) (*Nothing, error)
	StartChatTG(BotChat_StartChatTGServer) error
	StartChatVK(BotChat_StartChatVKServer) error
	UploadFile(context.Context, *FileUploadRequest) (*FileUploadResponse, error)
	ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error)
	CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error)
	GetHomeworks(context.Context, *GetHomeworksRequest) (*GetHomeworksResponse, error)
	CreateStudent(context.Context, *CreateStudentRequest) (*CreateStudentResponse, error)
	SendSolution(context.Context, *SendSolutionRequest) (*SendSolutionResponse, error)
	mustEmbedUnimplementedBotChatServer()
}

// UnimplementedBotChatServer must be embedded to have forward compatible implementations.
type UnimplementedBotChatServer struct {
}

func (UnimplementedBotChatServer) BroadcastMsg(context.Context, *BroadcastMessage) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastMsg not implemented")
}
func (UnimplementedBotChatServer) StartChatTG(BotChat_StartChatTGServer) error {
	return status.Errorf(codes.Unimplemented, "method StartChatTG not implemented")
}
func (UnimplementedBotChatServer) StartChatVK(BotChat_StartChatVKServer) error {
	return status.Errorf(codes.Unimplemented, "method StartChatVK not implemented")
}
func (UnimplementedBotChatServer) UploadFile(context.Context, *FileUploadRequest) (*FileUploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedBotChatServer) ValidateToken(context.Context, *ValidateTokenRequest) (*ValidateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateToken not implemented")
}
func (UnimplementedBotChatServer) CreateChat(context.Context, *CreateChatRequest) (*CreateChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChat not implemented")
}
func (UnimplementedBotChatServer) GetHomeworks(context.Context, *GetHomeworksRequest) (*GetHomeworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeworks not implemented")
}
func (UnimplementedBotChatServer) CreateStudent(context.Context, *CreateStudentRequest) (*CreateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStudent not implemented")
}
func (UnimplementedBotChatServer) SendSolution(context.Context, *SendSolutionRequest) (*SendSolutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSolution not implemented")
}
func (UnimplementedBotChatServer) mustEmbedUnimplementedBotChatServer() {}

// UnsafeBotChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotChatServer will
// result in compilation errors.
type UnsafeBotChatServer interface {
	mustEmbedUnimplementedBotChatServer()
}

func RegisterBotChatServer(s grpc.ServiceRegistrar, srv BotChatServer) {
	s.RegisterService(&BotChat_ServiceDesc, srv)
}

func _BotChat_BroadcastMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).BroadcastMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_BroadcastMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).BroadcastMsg(ctx, req.(*BroadcastMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotChat_StartChatTG_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotChatServer).StartChatTG(&botChatStartChatTGServer{stream})
}

type BotChat_StartChatTGServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type botChatStartChatTGServer struct {
	grpc.ServerStream
}

func (x *botChatStartChatTGServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *botChatStartChatTGServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BotChat_StartChatVK_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BotChatServer).StartChatVK(&botChatStartChatVKServer{stream})
}

type BotChat_StartChatVKServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type botChatStartChatVKServer struct {
	grpc.ServerStream
}

func (x *botChatStartChatVKServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *botChatStartChatVKServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BotChat_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).UploadFile(ctx, req.(*FileUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotChat_ValidateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).ValidateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_ValidateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).ValidateToken(ctx, req.(*ValidateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotChat_CreateChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).CreateChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_CreateChat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).CreateChat(ctx, req.(*CreateChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotChat_GetHomeworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).GetHomeworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_GetHomeworks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).GetHomeworks(ctx, req.(*GetHomeworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotChat_CreateStudent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).CreateStudent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_CreateStudent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).CreateStudent(ctx, req.(*CreateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotChat_SendSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSolutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotChatServer).SendSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotChat_SendSolution_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotChatServer).SendSolution(ctx, req.(*SendSolutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotChat_ServiceDesc is the grpc.ServiceDesc for BotChat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotChat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.BotChat",
	HandlerType: (*BotChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BroadcastMsg",
			Handler:    _BotChat_BroadcastMsg_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _BotChat_UploadFile_Handler,
		},
		{
			MethodName: "ValidateToken",
			Handler:    _BotChat_ValidateToken_Handler,
		},
		{
			MethodName: "CreateChat",
			Handler:    _BotChat_CreateChat_Handler,
		},
		{
			MethodName: "GetHomeworks",
			Handler:    _BotChat_GetHomeworks_Handler,
		},
		{
			MethodName: "CreateStudent",
			Handler:    _BotChat_CreateStudent_Handler,
		},
		{
			MethodName: "SendSolution",
			Handler:    _BotChat_SendSolution_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartChatTG",
			Handler:       _BotChat_StartChatTG_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "StartChatVK",
			Handler:       _BotChat_StartChatVK_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/proto/chat.proto",
}
