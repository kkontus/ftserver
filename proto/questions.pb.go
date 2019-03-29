// Code generated by protoc-gen-go. DO NOT EDIT.
// source: questions.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Questipns
type LoadQuestionsList struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadQuestionsList) Reset()         { *m = LoadQuestionsList{} }
func (m *LoadQuestionsList) String() string { return proto.CompactTextString(m) }
func (*LoadQuestionsList) ProtoMessage()    {}
func (*LoadQuestionsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{0}
}

func (m *LoadQuestionsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadQuestionsList.Unmarshal(m, b)
}
func (m *LoadQuestionsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadQuestionsList.Marshal(b, m, deterministic)
}
func (m *LoadQuestionsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadQuestionsList.Merge(m, src)
}
func (m *LoadQuestionsList) XXX_Size() int {
	return xxx_messageInfo_LoadQuestionsList.Size(m)
}
func (m *LoadQuestionsList) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadQuestionsList.DiscardUnknown(m)
}

var xxx_messageInfo_LoadQuestionsList proto.InternalMessageInfo

type ReturnQuestionsList struct {
	Result               []*Questions `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReturnQuestionsList) Reset()         { *m = ReturnQuestionsList{} }
func (m *ReturnQuestionsList) String() string { return proto.CompactTextString(m) }
func (*ReturnQuestionsList) ProtoMessage()    {}
func (*ReturnQuestionsList) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{1}
}

func (m *ReturnQuestionsList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnQuestionsList.Unmarshal(m, b)
}
func (m *ReturnQuestionsList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnQuestionsList.Marshal(b, m, deterministic)
}
func (m *ReturnQuestionsList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnQuestionsList.Merge(m, src)
}
func (m *ReturnQuestionsList) XXX_Size() int {
	return xxx_messageInfo_ReturnQuestionsList.Size(m)
}
func (m *ReturnQuestionsList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnQuestionsList.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnQuestionsList proto.InternalMessageInfo

func (m *ReturnQuestionsList) GetResult() []*Questions {
	if m != nil {
		return m.Result
	}
	return nil
}

type Questions struct {
	Question             string   `protobuf:"bytes,2,opt,name=Question,proto3" json:"Question,omitempty"`
	Correct              string   `protobuf:"bytes,3,opt,name=Correct,proto3" json:"Correct,omitempty"`
	Answers              []string `protobuf:"bytes,4,rep,name=Answers,proto3" json:"Answers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Questions) Reset()         { *m = Questions{} }
func (m *Questions) String() string { return proto.CompactTextString(m) }
func (*Questions) ProtoMessage()    {}
func (*Questions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{2}
}

func (m *Questions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Questions.Unmarshal(m, b)
}
func (m *Questions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Questions.Marshal(b, m, deterministic)
}
func (m *Questions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Questions.Merge(m, src)
}
func (m *Questions) XXX_Size() int {
	return xxx_messageInfo_Questions.Size(m)
}
func (m *Questions) XXX_DiscardUnknown() {
	xxx_messageInfo_Questions.DiscardUnknown(m)
}

var xxx_messageInfo_Questions proto.InternalMessageInfo

func (m *Questions) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *Questions) GetCorrect() string {
	if m != nil {
		return m.Correct
	}
	return ""
}

func (m *Questions) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

// Answers
type SendUserAnswers struct {
	Answers              []string `protobuf:"bytes,1,rep,name=answers,proto3" json:"answers,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendUserAnswers) Reset()         { *m = SendUserAnswers{} }
func (m *SendUserAnswers) String() string { return proto.CompactTextString(m) }
func (*SendUserAnswers) ProtoMessage()    {}
func (*SendUserAnswers) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{3}
}

func (m *SendUserAnswers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendUserAnswers.Unmarshal(m, b)
}
func (m *SendUserAnswers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendUserAnswers.Marshal(b, m, deterministic)
}
func (m *SendUserAnswers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendUserAnswers.Merge(m, src)
}
func (m *SendUserAnswers) XXX_Size() int {
	return xxx_messageInfo_SendUserAnswers.Size(m)
}
func (m *SendUserAnswers) XXX_DiscardUnknown() {
	xxx_messageInfo_SendUserAnswers.DiscardUnknown(m)
}

var xxx_messageInfo_SendUserAnswers proto.InternalMessageInfo

func (m *SendUserAnswers) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *SendUserAnswers) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type UserResults struct {
	Result               []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	Answers              []string `protobuf:"bytes,2,rep,name=answers,proto3" json:"answers,omitempty"`
	Percentage           string   `protobuf:"bytes,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResults) Reset()         { *m = UserResults{} }
func (m *UserResults) String() string { return proto.CompactTextString(m) }
func (*UserResults) ProtoMessage()    {}
func (*UserResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{4}
}

func (m *UserResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResults.Unmarshal(m, b)
}
func (m *UserResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResults.Marshal(b, m, deterministic)
}
func (m *UserResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResults.Merge(m, src)
}
func (m *UserResults) XXX_Size() int {
	return xxx_messageInfo_UserResults.Size(m)
}
func (m *UserResults) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResults.DiscardUnknown(m)
}

var xxx_messageInfo_UserResults proto.InternalMessageInfo

func (m *UserResults) GetResult() []string {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *UserResults) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *UserResults) GetPercentage() string {
	if m != nil {
		return m.Percentage
	}
	return ""
}

// Rankings
type LoadUserRanking struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadUserRanking) Reset()         { *m = LoadUserRanking{} }
func (m *LoadUserRanking) String() string { return proto.CompactTextString(m) }
func (*LoadUserRanking) ProtoMessage()    {}
func (*LoadUserRanking) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{5}
}

func (m *LoadUserRanking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadUserRanking.Unmarshal(m, b)
}
func (m *LoadUserRanking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadUserRanking.Marshal(b, m, deterministic)
}
func (m *LoadUserRanking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadUserRanking.Merge(m, src)
}
func (m *LoadUserRanking) XXX_Size() int {
	return xxx_messageInfo_LoadUserRanking.Size(m)
}
func (m *LoadUserRanking) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadUserRanking.DiscardUnknown(m)
}

var xxx_messageInfo_LoadUserRanking proto.InternalMessageInfo

func (m *LoadUserRanking) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type ReturnUserRanking struct {
	Score                string   `protobuf:"bytes,1,opt,name=score,proto3" json:"score,omitempty"`
	ScoreOverall         string   `protobuf:"bytes,2,opt,name=scoreOverall,proto3" json:"scoreOverall,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReturnUserRanking) Reset()         { *m = ReturnUserRanking{} }
func (m *ReturnUserRanking) String() string { return proto.CompactTextString(m) }
func (*ReturnUserRanking) ProtoMessage()    {}
func (*ReturnUserRanking) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e4cafc9881e02ba, []int{6}
}

func (m *ReturnUserRanking) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReturnUserRanking.Unmarshal(m, b)
}
func (m *ReturnUserRanking) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReturnUserRanking.Marshal(b, m, deterministic)
}
func (m *ReturnUserRanking) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReturnUserRanking.Merge(m, src)
}
func (m *ReturnUserRanking) XXX_Size() int {
	return xxx_messageInfo_ReturnUserRanking.Size(m)
}
func (m *ReturnUserRanking) XXX_DiscardUnknown() {
	xxx_messageInfo_ReturnUserRanking.DiscardUnknown(m)
}

var xxx_messageInfo_ReturnUserRanking proto.InternalMessageInfo

func (m *ReturnUserRanking) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *ReturnUserRanking) GetScoreOverall() string {
	if m != nil {
		return m.ScoreOverall
	}
	return ""
}

func init() {
	proto.RegisterType((*LoadQuestionsList)(nil), "proto.LoadQuestionsList")
	proto.RegisterType((*ReturnQuestionsList)(nil), "proto.ReturnQuestionsList")
	proto.RegisterType((*Questions)(nil), "proto.Questions")
	proto.RegisterType((*SendUserAnswers)(nil), "proto.SendUserAnswers")
	proto.RegisterType((*UserResults)(nil), "proto.UserResults")
	proto.RegisterType((*LoadUserRanking)(nil), "proto.LoadUserRanking")
	proto.RegisterType((*ReturnUserRanking)(nil), "proto.ReturnUserRanking")
}

func init() { proto.RegisterFile("questions.proto", fileDescriptor_8e4cafc9881e02ba) }

var fileDescriptor_8e4cafc9881e02ba = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xf3, 0x40,
	0x10, 0x26, 0xfd, 0x7a, 0xdf, 0x4e, 0x85, 0xb4, 0x5b, 0x29, 0x4b, 0x0e, 0x52, 0x16, 0x84, 0x9e,
	0x7a, 0xa8, 0x77, 0x4b, 0xa9, 0xd0, 0x4b, 0x45, 0x4c, 0xf1, 0xe4, 0x41, 0x62, 0x3a, 0xd4, 0xd0,
	0xb0, 0xa9, 0xbb, 0x9b, 0xfa, 0x7b, 0xfd, 0x27, 0x92, 0xcd, 0x24, 0x6e, 0xd4, 0xd3, 0xee, 0xf3,
	0xb1, 0xb3, 0xcf, 0xcc, 0x80, 0xff, 0x9e, 0xa3, 0x36, 0x49, 0x26, 0xf5, 0xfc, 0xa4, 0x32, 0x93,
	0xb1, 0xae, 0x3d, 0xc4, 0x18, 0x46, 0xdb, 0x2c, 0xda, 0x3f, 0x56, 0xea, 0x36, 0xd1, 0x46, 0x2c,
	0x61, 0x1c, 0xa2, 0xc9, 0x95, 0x6c, 0xd0, 0x6c, 0x06, 0x3d, 0x85, 0x3a, 0x4f, 0x0d, 0xf7, 0xa6,
	0xed, 0xd9, 0x60, 0x31, 0x2c, 0x4b, 0xcd, 0x6b, 0x57, 0x48, 0xba, 0x78, 0x86, 0x7e, 0x4d, 0xb2,
	0x00, 0xfe, 0x57, 0x80, 0xb7, 0xa6, 0xde, 0xac, 0x1f, 0xd6, 0x98, 0x71, 0xf8, 0xb7, 0xce, 0x94,
	0xc2, 0xd8, 0xf0, 0xb6, 0x95, 0x2a, 0x58, 0x28, 0x2b, 0xa9, 0x3f, 0x50, 0x69, 0xde, 0x99, 0xb6,
	0x0b, 0x85, 0xa0, 0x58, 0x82, 0xbf, 0x43, 0xb9, 0x7f, 0xd2, 0xa8, 0x88, 0x2a, 0xcc, 0x11, 0x99,
	0xbd, 0xd2, 0x4c, 0x90, 0x31, 0xe8, 0xe4, 0x1a, 0x15, 0x7d, 0x6c, 0xef, 0xe2, 0x05, 0x06, 0xc5,
	0xe3, 0xd0, 0x66, 0xd5, 0x6c, 0xd2, 0x68, 0xab, 0x5f, 0x35, 0xe1, 0x16, 0x6d, 0x35, 0x8b, 0x5e,
	0x01, 0x9c, 0x50, 0xc5, 0x28, 0x4d, 0x74, 0x40, 0x0a, 0xee, 0x30, 0xe2, 0x1a, 0xfc, 0x62, 0xa8,
	0xf6, 0x93, 0x48, 0x1e, 0x13, 0x79, 0xa8, 0x73, 0x78, 0x4e, 0x8e, 0x7b, 0x18, 0x95, 0x63, 0x76,
	0x8d, 0x97, 0xd0, 0xd5, 0x71, 0xa6, 0x90, 0x9c, 0x25, 0x60, 0x02, 0x2e, 0xec, 0xe5, 0xe1, 0x8c,
	0x2a, 0x4a, 0x53, 0x6a, 0xa7, 0xc1, 0x2d, 0x3e, 0x3d, 0x18, 0xd6, 0x53, 0xdf, 0xa1, 0x3a, 0x27,
	0x31, 0xb2, 0x0d, 0xf8, 0x1b, 0x34, 0xab, 0x34, 0xfd, 0xde, 0x07, 0xa7, 0xb5, 0xfd, 0xda, 0x7b,
	0x10, 0x90, 0xf2, 0xd7, 0xf2, 0x6f, 0x61, 0xb8, 0x7e, 0xc3, 0xf8, 0xe8, 0x8e, 0x7d, 0x42, 0xfe,
	0x1f, 0xeb, 0x08, 0x18, 0xf1, 0xee, 0x94, 0xef, 0x9c, 0xf7, 0x55, 0xaf, 0x13, 0x27, 0x89, 0xc3,
	0x07, 0xbc, 0x91, 0xc3, 0x51, 0x5e, 0x7b, 0x56, 0xb8, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe5,
	0x22, 0xce, 0x84, 0xcf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuestionsServiceClient is the client API for QuestionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuestionsServiceClient interface {
	GetAllQuestions(ctx context.Context, in *LoadQuestionsList, opts ...grpc.CallOption) (*ReturnQuestionsList, error)
	CheckUserAnswers(ctx context.Context, in *SendUserAnswers, opts ...grpc.CallOption) (*UserResults, error)
	CheckUserRanking(ctx context.Context, in *LoadUserRanking, opts ...grpc.CallOption) (*ReturnUserRanking, error)
}

type questionsServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuestionsServiceClient(cc *grpc.ClientConn) QuestionsServiceClient {
	return &questionsServiceClient{cc}
}

func (c *questionsServiceClient) GetAllQuestions(ctx context.Context, in *LoadQuestionsList, opts ...grpc.CallOption) (*ReturnQuestionsList, error) {
	out := new(ReturnQuestionsList)
	err := c.cc.Invoke(ctx, "/proto.QuestionsService/GetAllQuestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionsServiceClient) CheckUserAnswers(ctx context.Context, in *SendUserAnswers, opts ...grpc.CallOption) (*UserResults, error) {
	out := new(UserResults)
	err := c.cc.Invoke(ctx, "/proto.QuestionsService/CheckUserAnswers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *questionsServiceClient) CheckUserRanking(ctx context.Context, in *LoadUserRanking, opts ...grpc.CallOption) (*ReturnUserRanking, error) {
	out := new(ReturnUserRanking)
	err := c.cc.Invoke(ctx, "/proto.QuestionsService/CheckUserRanking", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuestionsServiceServer is the server API for QuestionsService service.
type QuestionsServiceServer interface {
	GetAllQuestions(context.Context, *LoadQuestionsList) (*ReturnQuestionsList, error)
	CheckUserAnswers(context.Context, *SendUserAnswers) (*UserResults, error)
	CheckUserRanking(context.Context, *LoadUserRanking) (*ReturnUserRanking, error)
}

// UnimplementedQuestionsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQuestionsServiceServer struct {
}

func (*UnimplementedQuestionsServiceServer) GetAllQuestions(ctx context.Context, req *LoadQuestionsList) (*ReturnQuestionsList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllQuestions not implemented")
}
func (*UnimplementedQuestionsServiceServer) CheckUserAnswers(ctx context.Context, req *SendUserAnswers) (*UserResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserAnswers not implemented")
}
func (*UnimplementedQuestionsServiceServer) CheckUserRanking(ctx context.Context, req *LoadUserRanking) (*ReturnUserRanking, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUserRanking not implemented")
}

func RegisterQuestionsServiceServer(s *grpc.Server, srv QuestionsServiceServer) {
	s.RegisterService(&_QuestionsService_serviceDesc, srv)
}

func _QuestionsService_GetAllQuestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadQuestionsList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).GetAllQuestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuestionsService/GetAllQuestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).GetAllQuestions(ctx, req.(*LoadQuestionsList))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionsService_CheckUserAnswers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendUserAnswers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).CheckUserAnswers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuestionsService/CheckUserAnswers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).CheckUserAnswers(ctx, req.(*SendUserAnswers))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuestionsService_CheckUserRanking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadUserRanking)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuestionsServiceServer).CheckUserRanking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.QuestionsService/CheckUserRanking",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuestionsServiceServer).CheckUserRanking(ctx, req.(*LoadUserRanking))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuestionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.QuestionsService",
	HandlerType: (*QuestionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllQuestions",
			Handler:    _QuestionsService_GetAllQuestions_Handler,
		},
		{
			MethodName: "CheckUserAnswers",
			Handler:    _QuestionsService_CheckUserAnswers_Handler,
		},
		{
			MethodName: "CheckUserRanking",
			Handler:    _QuestionsService_CheckUserRanking_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "questions.proto",
}
