// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/proto/chat.proto

package chat

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	ChatID               int32    `protobuf:"varint,1,opt,name=chatID,proto3" json:"chatID,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	AttachmentURLs       []string `protobuf:"bytes,3,rep,name=attachmentURLs,proto3" json:"attachmentURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetChatID() int32 {
	if m != nil {
		return m.ChatID
	}
	return 0
}

func (m *Message) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Message) GetAttachmentURLs() []string {
	if m != nil {
		return m.AttachmentURLs
	}
	return nil
}

type FileUploadRequest struct {
	Mimetype             string   `protobuf:"bytes,1,opt,name=mimetype,proto3" json:"mimetype,omitempty"`
	FileURL              string   `protobuf:"bytes,2,opt,name=fileURL,proto3" json:"fileURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileUploadRequest) Reset()         { *m = FileUploadRequest{} }
func (m *FileUploadRequest) String() string { return proto.CompactTextString(m) }
func (*FileUploadRequest) ProtoMessage()    {}
func (*FileUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{1}
}

func (m *FileUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileUploadRequest.Unmarshal(m, b)
}
func (m *FileUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileUploadRequest.Marshal(b, m, deterministic)
}
func (m *FileUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileUploadRequest.Merge(m, src)
}
func (m *FileUploadRequest) XXX_Size() int {
	return xxx_messageInfo_FileUploadRequest.Size(m)
}
func (m *FileUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileUploadRequest proto.InternalMessageInfo

func (m *FileUploadRequest) GetMimetype() string {
	if m != nil {
		return m.Mimetype
	}
	return ""
}

func (m *FileUploadRequest) GetFileURL() string {
	if m != nil {
		return m.FileURL
	}
	return ""
}

type FileUploadResponse struct {
	InternalFileURL      string   `protobuf:"bytes,1,opt,name=internalFileURL,proto3" json:"internalFileURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileUploadResponse) Reset()         { *m = FileUploadResponse{} }
func (m *FileUploadResponse) String() string { return proto.CompactTextString(m) }
func (*FileUploadResponse) ProtoMessage()    {}
func (*FileUploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{2}
}

func (m *FileUploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileUploadResponse.Unmarshal(m, b)
}
func (m *FileUploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileUploadResponse.Marshal(b, m, deterministic)
}
func (m *FileUploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileUploadResponse.Merge(m, src)
}
func (m *FileUploadResponse) XXX_Size() int {
	return xxx_messageInfo_FileUploadResponse.Size(m)
}
func (m *FileUploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FileUploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FileUploadResponse proto.InternalMessageInfo

func (m *FileUploadResponse) GetInternalFileURL() string {
	if m != nil {
		return m.InternalFileURL
	}
	return ""
}

type TaskData struct {
	Description          string   `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	AttachmentURLs       []string `protobuf:"bytes,2,rep,name=attachmentURLs,proto3" json:"attachmentURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskData) Reset()         { *m = TaskData{} }
func (m *TaskData) String() string { return proto.CompactTextString(m) }
func (*TaskData) ProtoMessage()    {}
func (*TaskData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{3}
}

func (m *TaskData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskData.Unmarshal(m, b)
}
func (m *TaskData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskData.Marshal(b, m, deterministic)
}
func (m *TaskData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskData.Merge(m, src)
}
func (m *TaskData) XXX_Size() int {
	return xxx_messageInfo_TaskData.Size(m)
}
func (m *TaskData) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskData.DiscardUnknown(m)
}

var xxx_messageInfo_TaskData proto.InternalMessageInfo

func (m *TaskData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TaskData) GetAttachmentURLs() []string {
	if m != nil {
		return m.AttachmentURLs
	}
	return nil
}

type HomeworkData struct {
	HomeworkID           int32       `protobuf:"varint,1,opt,name=homeworkID,proto3" json:"homeworkID,omitempty"`
	Title                string      `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string      `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tasks                []*TaskData `protobuf:"bytes,4,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HomeworkData) Reset()         { *m = HomeworkData{} }
func (m *HomeworkData) String() string { return proto.CompactTextString(m) }
func (*HomeworkData) ProtoMessage()    {}
func (*HomeworkData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{4}
}

func (m *HomeworkData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HomeworkData.Unmarshal(m, b)
}
func (m *HomeworkData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HomeworkData.Marshal(b, m, deterministic)
}
func (m *HomeworkData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HomeworkData.Merge(m, src)
}
func (m *HomeworkData) XXX_Size() int {
	return xxx_messageInfo_HomeworkData.Size(m)
}
func (m *HomeworkData) XXX_DiscardUnknown() {
	xxx_messageInfo_HomeworkData.DiscardUnknown(m)
}

var xxx_messageInfo_HomeworkData proto.InternalMessageInfo

func (m *HomeworkData) GetHomeworkID() int32 {
	if m != nil {
		return m.HomeworkID
	}
	return 0
}

func (m *HomeworkData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *HomeworkData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *HomeworkData) GetTasks() []*TaskData {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type GetHomeworksRequest struct {
	ClassID              int32    `protobuf:"varint,1,opt,name=classID,proto3" json:"classID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHomeworksRequest) Reset()         { *m = GetHomeworksRequest{} }
func (m *GetHomeworksRequest) String() string { return proto.CompactTextString(m) }
func (*GetHomeworksRequest) ProtoMessage()    {}
func (*GetHomeworksRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{5}
}

func (m *GetHomeworksRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHomeworksRequest.Unmarshal(m, b)
}
func (m *GetHomeworksRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHomeworksRequest.Marshal(b, m, deterministic)
}
func (m *GetHomeworksRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHomeworksRequest.Merge(m, src)
}
func (m *GetHomeworksRequest) XXX_Size() int {
	return xxx_messageInfo_GetHomeworksRequest.Size(m)
}
func (m *GetHomeworksRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHomeworksRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHomeworksRequest proto.InternalMessageInfo

func (m *GetHomeworksRequest) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

type GetHomeworksResponse struct {
	Homeworks            []*HomeworkData `protobuf:"bytes,1,rep,name=homeworks,proto3" json:"homeworks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetHomeworksResponse) Reset()         { *m = GetHomeworksResponse{} }
func (m *GetHomeworksResponse) String() string { return proto.CompactTextString(m) }
func (*GetHomeworksResponse) ProtoMessage()    {}
func (*GetHomeworksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{6}
}

func (m *GetHomeworksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHomeworksResponse.Unmarshal(m, b)
}
func (m *GetHomeworksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHomeworksResponse.Marshal(b, m, deterministic)
}
func (m *GetHomeworksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHomeworksResponse.Merge(m, src)
}
func (m *GetHomeworksResponse) XXX_Size() int {
	return xxx_messageInfo_GetHomeworksResponse.Size(m)
}
func (m *GetHomeworksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHomeworksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHomeworksResponse proto.InternalMessageInfo

func (m *GetHomeworksResponse) GetHomeworks() []*HomeworkData {
	if m != nil {
		return m.Homeworks
	}
	return nil
}

type SolutionData struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	AttachmentURLs       []string `protobuf:"bytes,2,rep,name=attachmentURLs,proto3" json:"attachmentURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SolutionData) Reset()         { *m = SolutionData{} }
func (m *SolutionData) String() string { return proto.CompactTextString(m) }
func (*SolutionData) ProtoMessage()    {}
func (*SolutionData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{7}
}

func (m *SolutionData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SolutionData.Unmarshal(m, b)
}
func (m *SolutionData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SolutionData.Marshal(b, m, deterministic)
}
func (m *SolutionData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SolutionData.Merge(m, src)
}
func (m *SolutionData) XXX_Size() int {
	return xxx_messageInfo_SolutionData.Size(m)
}
func (m *SolutionData) XXX_DiscardUnknown() {
	xxx_messageInfo_SolutionData.DiscardUnknown(m)
}

var xxx_messageInfo_SolutionData proto.InternalMessageInfo

func (m *SolutionData) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *SolutionData) GetAttachmentURLs() []string {
	if m != nil {
		return m.AttachmentURLs
	}
	return nil
}

type SendSolutionRequest struct {
	HomeworkID           int32         `protobuf:"varint,1,opt,name=homeworkID,proto3" json:"homeworkID,omitempty"`
	Solution             *SolutionData `protobuf:"bytes,2,opt,name=solution,proto3" json:"solution,omitempty"`
	StudentID            int32         `protobuf:"varint,3,opt,name=studentID,proto3" json:"studentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SendSolutionRequest) Reset()         { *m = SendSolutionRequest{} }
func (m *SendSolutionRequest) String() string { return proto.CompactTextString(m) }
func (*SendSolutionRequest) ProtoMessage()    {}
func (*SendSolutionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{8}
}

func (m *SendSolutionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSolutionRequest.Unmarshal(m, b)
}
func (m *SendSolutionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSolutionRequest.Marshal(b, m, deterministic)
}
func (m *SendSolutionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSolutionRequest.Merge(m, src)
}
func (m *SendSolutionRequest) XXX_Size() int {
	return xxx_messageInfo_SendSolutionRequest.Size(m)
}
func (m *SendSolutionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSolutionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendSolutionRequest proto.InternalMessageInfo

func (m *SendSolutionRequest) GetHomeworkID() int32 {
	if m != nil {
		return m.HomeworkID
	}
	return 0
}

func (m *SendSolutionRequest) GetSolution() *SolutionData {
	if m != nil {
		return m.Solution
	}
	return nil
}

func (m *SendSolutionRequest) GetStudentID() int32 {
	if m != nil {
		return m.StudentID
	}
	return 0
}

type SendSolutionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendSolutionResponse) Reset()         { *m = SendSolutionResponse{} }
func (m *SendSolutionResponse) String() string { return proto.CompactTextString(m) }
func (*SendSolutionResponse) ProtoMessage()    {}
func (*SendSolutionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{9}
}

func (m *SendSolutionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendSolutionResponse.Unmarshal(m, b)
}
func (m *SendSolutionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendSolutionResponse.Marshal(b, m, deterministic)
}
func (m *SendSolutionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendSolutionResponse.Merge(m, src)
}
func (m *SendSolutionResponse) XXX_Size() int {
	return xxx_messageInfo_SendSolutionResponse.Size(m)
}
func (m *SendSolutionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendSolutionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendSolutionResponse proto.InternalMessageInfo

type ValidateTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenRequest) Reset()         { *m = ValidateTokenRequest{} }
func (m *ValidateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenRequest) ProtoMessage()    {}
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{10}
}

func (m *ValidateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenRequest.Unmarshal(m, b)
}
func (m *ValidateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenRequest.Marshal(b, m, deterministic)
}
func (m *ValidateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenRequest.Merge(m, src)
}
func (m *ValidateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenRequest.Size(m)
}
func (m *ValidateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenRequest proto.InternalMessageInfo

func (m *ValidateTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ValidateTokenResponse struct {
	ClassID              int32    `protobuf:"varint,1,opt,name=classID,proto3" json:"classID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValidateTokenResponse) Reset()         { *m = ValidateTokenResponse{} }
func (m *ValidateTokenResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTokenResponse) ProtoMessage()    {}
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{11}
}

func (m *ValidateTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidateTokenResponse.Unmarshal(m, b)
}
func (m *ValidateTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidateTokenResponse.Marshal(b, m, deterministic)
}
func (m *ValidateTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTokenResponse.Merge(m, src)
}
func (m *ValidateTokenResponse) XXX_Size() int {
	return xxx_messageInfo_ValidateTokenResponse.Size(m)
}
func (m *ValidateTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTokenResponse proto.InternalMessageInfo

func (m *ValidateTokenResponse) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

type CreateStudentRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	AvatarURL            string   `protobuf:"bytes,3,opt,name=avatarURL,proto3" json:"avatarURL,omitempty"`
	ClassID              int32    `protobuf:"varint,4,opt,name=classID,proto3" json:"classID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStudentRequest) Reset()         { *m = CreateStudentRequest{} }
func (m *CreateStudentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateStudentRequest) ProtoMessage()    {}
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{12}
}

func (m *CreateStudentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStudentRequest.Unmarshal(m, b)
}
func (m *CreateStudentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStudentRequest.Marshal(b, m, deterministic)
}
func (m *CreateStudentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStudentRequest.Merge(m, src)
}
func (m *CreateStudentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateStudentRequest.Size(m)
}
func (m *CreateStudentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStudentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStudentRequest proto.InternalMessageInfo

func (m *CreateStudentRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateStudentRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateStudentRequest) GetAvatarURL() string {
	if m != nil {
		return m.AvatarURL
	}
	return ""
}

func (m *CreateStudentRequest) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

type CreateStudentResponse struct {
	StudentID            int32    `protobuf:"varint,1,opt,name=studentID,proto3" json:"studentID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateStudentResponse) Reset()         { *m = CreateStudentResponse{} }
func (m *CreateStudentResponse) String() string { return proto.CompactTextString(m) }
func (*CreateStudentResponse) ProtoMessage()    {}
func (*CreateStudentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{13}
}

func (m *CreateStudentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateStudentResponse.Unmarshal(m, b)
}
func (m *CreateStudentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateStudentResponse.Marshal(b, m, deterministic)
}
func (m *CreateStudentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateStudentResponse.Merge(m, src)
}
func (m *CreateStudentResponse) XXX_Size() int {
	return xxx_messageInfo_CreateStudentResponse.Size(m)
}
func (m *CreateStudentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateStudentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateStudentResponse proto.InternalMessageInfo

func (m *CreateStudentResponse) GetStudentID() int32 {
	if m != nil {
		return m.StudentID
	}
	return 0
}

type CreateChatRequest struct {
	StudentID            int32    `protobuf:"varint,1,opt,name=studentID,proto3" json:"studentID,omitempty"`
	ClassID              int32    `protobuf:"varint,2,opt,name=classID,proto3" json:"classID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatRequest) Reset()         { *m = CreateChatRequest{} }
func (m *CreateChatRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChatRequest) ProtoMessage()    {}
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{14}
}

func (m *CreateChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatRequest.Unmarshal(m, b)
}
func (m *CreateChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatRequest.Marshal(b, m, deterministic)
}
func (m *CreateChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatRequest.Merge(m, src)
}
func (m *CreateChatRequest) XXX_Size() int {
	return xxx_messageInfo_CreateChatRequest.Size(m)
}
func (m *CreateChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatRequest proto.InternalMessageInfo

func (m *CreateChatRequest) GetStudentID() int32 {
	if m != nil {
		return m.StudentID
	}
	return 0
}

func (m *CreateChatRequest) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

type CreateChatResponse struct {
	InternalChatID       int32    `protobuf:"varint,1,opt,name=internalChatID,proto3" json:"internalChatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatResponse) Reset()         { *m = CreateChatResponse{} }
func (m *CreateChatResponse) String() string { return proto.CompactTextString(m) }
func (*CreateChatResponse) ProtoMessage()    {}
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{15}
}

func (m *CreateChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatResponse.Unmarshal(m, b)
}
func (m *CreateChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatResponse.Marshal(b, m, deterministic)
}
func (m *CreateChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatResponse.Merge(m, src)
}
func (m *CreateChatResponse) XXX_Size() int {
	return xxx_messageInfo_CreateChatResponse.Size(m)
}
func (m *CreateChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatResponse proto.InternalMessageInfo

func (m *CreateChatResponse) GetInternalChatID() int32 {
	if m != nil {
		return m.InternalChatID
	}
	return 0
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{16}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

type BroadcastMessage struct {
	ClassID              int32    `protobuf:"varint,1,opt,name=classID,proto3" json:"classID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AttachmentURLs       []string `protobuf:"bytes,4,rep,name=attachmentURLs,proto3" json:"attachmentURLs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastMessage) Reset()         { *m = BroadcastMessage{} }
func (m *BroadcastMessage) String() string { return proto.CompactTextString(m) }
func (*BroadcastMessage) ProtoMessage()    {}
func (*BroadcastMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{17}
}

func (m *BroadcastMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastMessage.Unmarshal(m, b)
}
func (m *BroadcastMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastMessage.Marshal(b, m, deterministic)
}
func (m *BroadcastMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastMessage.Merge(m, src)
}
func (m *BroadcastMessage) XXX_Size() int {
	return xxx_messageInfo_BroadcastMessage.Size(m)
}
func (m *BroadcastMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastMessage proto.InternalMessageInfo

func (m *BroadcastMessage) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

func (m *BroadcastMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BroadcastMessage) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *BroadcastMessage) GetAttachmentURLs() []string {
	if m != nil {
		return m.AttachmentURLs
	}
	return nil
}

type EventData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartDate            string   `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate              string   `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventData) Reset()         { *m = EventData{} }
func (m *EventData) String() string { return proto.CompactTextString(m) }
func (*EventData) ProtoMessage()    {}
func (*EventData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{18}
}

func (m *EventData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventData.Unmarshal(m, b)
}
func (m *EventData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventData.Marshal(b, m, deterministic)
}
func (m *EventData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventData.Merge(m, src)
}
func (m *EventData) XXX_Size() int {
	return xxx_messageInfo_EventData.Size(m)
}
func (m *EventData) XXX_DiscardUnknown() {
	xxx_messageInfo_EventData.DiscardUnknown(m)
}

var xxx_messageInfo_EventData proto.InternalMessageInfo

func (m *EventData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventData) GetStartDate() string {
	if m != nil {
		return m.StartDate
	}
	return ""
}

func (m *EventData) GetEndDate() string {
	if m != nil {
		return m.EndDate
	}
	return ""
}

type GetEventsRequest struct {
	ClassID              int32    `protobuf:"varint,1,opt,name=classID,proto3" json:"classID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{19}
}

func (m *GetEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequest.Unmarshal(m, b)
}
func (m *GetEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequest.Merge(m, src)
}
func (m *GetEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequest.Size(m)
}
func (m *GetEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequest proto.InternalMessageInfo

func (m *GetEventsRequest) GetClassID() int32 {
	if m != nil {
		return m.ClassID
	}
	return 0
}

type GetEventsResponse struct {
	Events               []*EventData `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetEventsResponse) Reset()         { *m = GetEventsResponse{} }
func (m *GetEventsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEventsResponse) ProtoMessage()    {}
func (*GetEventsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3b146eda5bde7a7, []int{20}
}

func (m *GetEventsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsResponse.Unmarshal(m, b)
}
func (m *GetEventsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsResponse.Marshal(b, m, deterministic)
}
func (m *GetEventsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsResponse.Merge(m, src)
}
func (m *GetEventsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEventsResponse.Size(m)
}
func (m *GetEventsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsResponse proto.InternalMessageInfo

func (m *GetEventsResponse) GetEvents() []*EventData {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "chat.Message")
	proto.RegisterType((*FileUploadRequest)(nil), "chat.FileUploadRequest")
	proto.RegisterType((*FileUploadResponse)(nil), "chat.FileUploadResponse")
	proto.RegisterType((*TaskData)(nil), "chat.TaskData")
	proto.RegisterType((*HomeworkData)(nil), "chat.HomeworkData")
	proto.RegisterType((*GetHomeworksRequest)(nil), "chat.GetHomeworksRequest")
	proto.RegisterType((*GetHomeworksResponse)(nil), "chat.GetHomeworksResponse")
	proto.RegisterType((*SolutionData)(nil), "chat.SolutionData")
	proto.RegisterType((*SendSolutionRequest)(nil), "chat.SendSolutionRequest")
	proto.RegisterType((*SendSolutionResponse)(nil), "chat.SendSolutionResponse")
	proto.RegisterType((*ValidateTokenRequest)(nil), "chat.ValidateTokenRequest")
	proto.RegisterType((*ValidateTokenResponse)(nil), "chat.ValidateTokenResponse")
	proto.RegisterType((*CreateStudentRequest)(nil), "chat.CreateStudentRequest")
	proto.RegisterType((*CreateStudentResponse)(nil), "chat.CreateStudentResponse")
	proto.RegisterType((*CreateChatRequest)(nil), "chat.CreateChatRequest")
	proto.RegisterType((*CreateChatResponse)(nil), "chat.CreateChatResponse")
	proto.RegisterType((*Nothing)(nil), "chat.Nothing")
	proto.RegisterType((*BroadcastMessage)(nil), "chat.BroadcastMessage")
	proto.RegisterType((*EventData)(nil), "chat.EventData")
	proto.RegisterType((*GetEventsRequest)(nil), "chat.GetEventsRequest")
	proto.RegisterType((*GetEventsResponse)(nil), "chat.GetEventsResponse")
}

func init() {
	proto.RegisterFile("src/proto/chat.proto", fileDescriptor_d3b146eda5bde7a7)
}

var fileDescriptor_d3b146eda5bde7a7 = []byte{
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x4f, 0xdb, 0x4a,
	0x10, 0x97, 0xf3, 0xed, 0x21, 0x04, 0x58, 0x02, 0xf8, 0xf9, 0xa1, 0xa7, 0xc8, 0x7a, 0xe2, 0xe5,
	0x80, 0x12, 0x1e, 0x15, 0x37, 0x54, 0x24, 0xbe, 0x29, 0x94, 0x83, 0x13, 0x38, 0x54, 0xea, 0x61,
	0x1b, 0x2f, 0xc4, 0x22, 0xb1, 0x53, 0xef, 0x42, 0xdb, 0x73, 0x8f, 0xed, 0xdf, 0xd3, 0xbf, 0xaf,
	0xda, 0x2f, 0xc7, 0x76, 0x1c, 0x8a, 0xb8, 0x79, 0x66, 0x67, 0x76, 0x7e, 0xf3, 0x9b, 0x8f, 0x35,
	0x34, 0x69, 0x34, 0xe8, 0x4e, 0xa2, 0x90, 0x85, 0xdd, 0xc1, 0x10, 0xb3, 0x8e, 0xf8, 0x44, 0x25,
	0xfe, 0xed, 0x7c, 0x84, 0xea, 0x7b, 0x42, 0x29, 0xbe, 0x27, 0x68, 0x1d, 0x2a, 0x5c, 0x75, 0x71,
	0x6c, 0x19, 0x2d, 0xa3, 0x5d, 0x76, 0x95, 0x84, 0x10, 0x94, 0x18, 0xf9, 0xca, 0xac, 0x42, 0xcb,
	0x68, 0x9b, 0xae, 0xf8, 0x46, 0x5b, 0xd0, 0xc0, 0x8c, 0xe1, 0xc1, 0x70, 0x4c, 0x02, 0x76, 0xe3,
	0x5e, 0x51, 0xab, 0xd8, 0x2a, 0xb6, 0x4d, 0x37, 0xa3, 0x75, 0x2e, 0x60, 0xe5, 0xd4, 0x1f, 0x91,
	0x9b, 0xc9, 0x28, 0xc4, 0x9e, 0x4b, 0x3e, 0x3f, 0x12, 0xca, 0x90, 0x0d, 0xb5, 0xb1, 0x3f, 0x26,
	0xec, 0xdb, 0x84, 0x88, 0x50, 0xa6, 0x1b, 0xcb, 0xc8, 0x82, 0xea, 0x1d, 0x77, 0x70, 0xaf, 0x54,
	0x3c, 0x2d, 0x3a, 0x6f, 0x01, 0x25, 0xaf, 0xa2, 0x93, 0x30, 0xa0, 0x04, 0xb5, 0x61, 0xc9, 0x0f,
	0x18, 0x89, 0x02, 0x3c, 0x3a, 0x55, 0x7e, 0xf2, 0xca, 0xac, 0xda, 0xe9, 0x43, 0xad, 0x8f, 0xe9,
	0xc3, 0x31, 0x66, 0x18, 0xb5, 0x60, 0xc1, 0x23, 0x74, 0x10, 0xf9, 0x13, 0xe6, 0x87, 0x81, 0xf2,
	0x48, 0xaa, 0x72, 0x12, 0x2c, 0xe4, 0x26, 0xf8, 0xd3, 0x80, 0xfa, 0x79, 0x38, 0x26, 0x5f, 0xc2,
	0x48, 0x5e, 0xfd, 0x0f, 0xc0, 0x50, 0xc9, 0x31, 0x93, 0x09, 0x0d, 0x6a, 0x42, 0x99, 0xf9, 0x6c,
	0x44, 0x54, 0x7a, 0x52, 0xc8, 0x02, 0x2a, 0xce, 0x02, 0xfa, 0x17, 0xca, 0x0c, 0xd3, 0x07, 0x6a,
	0x95, 0x5a, 0xc5, 0xf6, 0xc2, 0x6e, 0xa3, 0x23, 0x4a, 0xa9, 0x33, 0x72, 0xe5, 0xa1, 0xd3, 0x85,
	0xd5, 0x33, 0xc2, 0x34, 0x20, 0xaa, 0x19, 0xb7, 0xa0, 0x3a, 0x18, 0x61, 0x4a, 0x63, 0x44, 0x5a,
	0x74, 0xce, 0xa1, 0x99, 0x76, 0x50, 0xbc, 0xee, 0x80, 0xa9, 0x41, 0x53, 0xcb, 0x10, 0x21, 0x91,
	0x0c, 0x99, 0xcc, 0xd6, 0x9d, 0x1a, 0x39, 0xef, 0xa0, 0xde, 0x0b, 0x47, 0x8f, 0x1c, 0xac, 0x20,
	0x42, 0xb7, 0x8d, 0xf1, 0x6c, 0xdb, 0xe4, 0xb3, 0xfa, 0xdd, 0x80, 0xd5, 0x1e, 0x09, 0x3c, 0x7d,
	0xa1, 0xce, 0xe3, 0x4f, 0xe4, 0x76, 0xa0, 0x46, 0x95, 0x8b, 0xe0, 0x37, 0x06, 0x9d, 0x44, 0xe6,
	0xc6, 0x36, 0x68, 0x13, 0x4c, 0xca, 0x1e, 0x3d, 0x12, 0xf0, 0xae, 0x2f, 0x8a, 0xeb, 0xa6, 0x0a,
	0x67, 0x1d, 0x9a, 0x69, 0x10, 0x92, 0x1b, 0x67, 0x1b, 0x9a, 0xb7, 0x78, 0xe4, 0x7b, 0x98, 0x91,
	0x7e, 0xf8, 0x40, 0x62, 0x74, 0xbc, 0xb4, 0x5c, 0x56, 0x29, 0x4b, 0xc1, 0xf9, 0x1f, 0xd6, 0x32,
	0xd6, 0x8a, 0xe2, 0xf9, 0x45, 0x79, 0x82, 0xe6, 0x51, 0x44, 0x30, 0x23, 0x3d, 0x89, 0x45, 0x07,
	0x40, 0x50, 0x0a, 0xf0, 0x58, 0x0f, 0x8d, 0xf8, 0x16, 0x34, 0xf3, 0x41, 0xd2, 0xd3, 0xc9, 0x87,
	0x68, 0x13, 0x4c, 0xfc, 0x84, 0x19, 0x8e, 0xf8, 0x38, 0xc8, 0x5e, 0x9a, 0x2a, 0x92, 0x71, 0x4b,
	0xe9, 0xb8, 0x7b, 0xb0, 0x96, 0x89, 0xab, 0xa0, 0xa6, 0x78, 0x32, 0xb2, 0x3c, 0x5d, 0xc2, 0x8a,
	0x74, 0x3b, 0x1a, 0xe2, 0x18, 0xeb, 0xb3, 0x2e, 0x49, 0x0c, 0x85, 0x34, 0x86, 0x7d, 0x40, 0xc9,
	0xcb, 0x14, 0x80, 0x2d, 0x68, 0xe8, 0x79, 0x3e, 0x4a, 0xee, 0xa8, 0x8c, 0xd6, 0x31, 0xa1, 0x7a,
	0x1d, 0xb2, 0xa1, 0x1f, 0xdc, 0xf3, 0xc9, 0x5c, 0x3e, 0x8c, 0x42, 0xec, 0x0d, 0x30, 0x65, 0x7a,
	0xc7, 0xcd, 0xe5, 0xfc, 0xd5, 0x73, 0x39, 0xdb, 0xd2, 0xa5, 0xdc, 0x96, 0xfe, 0x61, 0x80, 0x79,
	0xf2, 0x44, 0x02, 0x26, 0x86, 0xa3, 0x01, 0x05, 0xdf, 0x53, 0x75, 0x2c, 0xf8, 0xde, 0xab, 0xa3,
	0x0b, 0x96, 0x71, 0xc4, 0x2f, 0x25, 0xa2, 0x9a, 0xa6, 0x3b, 0x55, 0xf0, 0x6c, 0x49, 0xe0, 0x89,
	0xb3, 0xb2, 0x5c, 0xa6, 0x4a, 0x74, 0xb6, 0x61, 0xf9, 0x8c, 0x30, 0x81, 0xe7, 0x05, 0x4b, 0x62,
	0x1f, 0x56, 0x12, 0xd6, 0xaa, 0x24, 0xff, 0x41, 0x85, 0x08, 0x8d, 0x5a, 0x0f, 0x4b, 0x72, 0xd2,
	0xe2, 0x1c, 0x5d, 0x75, 0xbc, 0xfb, 0xab, 0x0c, 0xd5, 0xc3, 0x90, 0xf1, 0x0a, 0xa1, 0x3d, 0xa8,
	0x4f, 0x6b, 0x42, 0xef, 0xd1, 0xba, 0x74, 0xca, 0xd6, 0xc9, 0x5e, 0x94, 0x7a, 0x55, 0x4b, 0xb4,
	0x03, 0xcb, 0x7c, 0x12, 0xaf, 0x43, 0xe6, 0xdf, 0xf9, 0x03, 0x2c, 0x52, 0x57, 0x26, 0x73, 0x3c,
	0xba, 0xb0, 0xd0, 0xe3, 0x3c, 0xf0, 0xa8, 0xfd, 0xb3, 0x39, 0xc6, 0x4a, 0x6c, 0x1b, 0x3b, 0x46,
	0xca, 0xe1, 0xf6, 0xf2, 0x05, 0x0e, 0x07, 0x00, 0xf2, 0x2d, 0xe2, 0x0f, 0x0c, 0xda, 0x90, 0x06,
	0x33, 0x8f, 0x9d, 0x6d, 0xcd, 0x1e, 0x28, 0x02, 0xcf, 0x61, 0x31, 0xb5, 0x18, 0x90, 0x2d, 0x4d,
	0xf3, 0x76, 0x8b, 0xfd, 0x77, 0xee, 0x99, 0xba, 0xe9, 0x00, 0x60, 0x3a, 0x33, 0x1a, 0xca, 0xcc,
	0x48, 0x6a, 0x28, 0x39, 0xe3, 0x75, 0x02, 0xf5, 0xe4, 0x2b, 0x80, 0xfe, 0x92, 0x96, 0x39, 0x4f,
	0x89, 0x6d, 0xe7, 0x1d, 0x4d, 0x33, 0x4a, 0xed, 0x0f, 0x9d, 0x51, 0xde, 0x32, 0xd3, 0x19, 0xe5,
	0x2f, 0x9c, 0x13, 0xa8, 0x27, 0x57, 0xaf, 0x06, 0x94, 0xf3, 0x26, 0x68, 0x40, 0x79, 0x9b, 0x1a,
	0xed, 0x83, 0x19, 0x37, 0xae, 0xee, 0xb5, 0x6c, 0xdf, 0xdb, 0x1b, 0x33, 0x7a, 0xe9, 0x7d, 0x58,
	0xfb, 0x50, 0xe9, 0x88, 0x3f, 0xa6, 0x4f, 0x15, 0xf1, 0xcb, 0xf4, 0xe6, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xac, 0x4a, 0x28, 0x86, 0x4a, 0x09, 0x00, 0x00,
}
