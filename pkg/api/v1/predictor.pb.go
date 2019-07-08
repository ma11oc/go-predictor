// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/v1/predictor.proto

package predictor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type PersonConfig_Gender int32

const (
	PersonConfig_OTHER  PersonConfig_Gender = 0
	PersonConfig_MALE   PersonConfig_Gender = 1
	PersonConfig_FEMALE PersonConfig_Gender = 2
)

var PersonConfig_Gender_name = map[int32]string{
	0: "OTHER",
	1: "MALE",
	2: "FEMALE",
}

var PersonConfig_Gender_value = map[string]int32{
	"OTHER":  0,
	"MALE":   1,
	"FEMALE": 2,
}

func (x PersonConfig_Gender) String() string {
	return proto.EnumName(PersonConfig_Gender_name, int32(x))
}

func (PersonConfig_Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{0, 0}
}

type PersonConfig_Features int32

const (
	PersonConfig_NOT_SPECIFIED  PersonConfig_Features = 0
	PersonConfig_BUSINESS_OWNER PersonConfig_Features = 1
	PersonConfig_CREATOR        PersonConfig_Features = 2
)

var PersonConfig_Features_name = map[int32]string{
	0: "NOT_SPECIFIED",
	1: "BUSINESS_OWNER",
	2: "CREATOR",
}

var PersonConfig_Features_value = map[string]int32{
	"NOT_SPECIFIED":  0,
	"BUSINESS_OWNER": 1,
	"CREATOR":        2,
}

func (x PersonConfig_Features) String() string {
	return proto.EnumName(PersonConfig_Features_name, int32(x))
}

func (PersonConfig_Features) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{0, 1}
}

type PersonConfig struct {
	Name                 string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Gender               PersonConfig_Gender `protobuf:"varint,4,opt,name=gender,proto3,enum=predictor.PersonConfig_Gender" json:"gender,omitempty"`
	Birthday             string              `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Features             uint32              `protobuf:"varint,6,opt,name=features,proto3" json:"features,omitempty"`
	Environment          []*PersonConfig     `protobuf:"bytes,7,rep,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PersonConfig) Reset()         { *m = PersonConfig{} }
func (m *PersonConfig) String() string { return proto.CompactTextString(m) }
func (*PersonConfig) ProtoMessage()    {}
func (*PersonConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{0}
}

func (m *PersonConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonConfig.Unmarshal(m, b)
}
func (m *PersonConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonConfig.Marshal(b, m, deterministic)
}
func (m *PersonConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonConfig.Merge(m, src)
}
func (m *PersonConfig) XXX_Size() int {
	return xxx_messageInfo_PersonConfig.Size(m)
}
func (m *PersonConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PersonConfig proto.InternalMessageInfo

func (m *PersonConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PersonConfig) GetGender() PersonConfig_Gender {
	if m != nil {
		return m.Gender
	}
	return PersonConfig_OTHER
}

func (m *PersonConfig) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *PersonConfig) GetFeatures() uint32 {
	if m != nil {
		return m.Features
	}
	return 0
}

func (m *PersonConfig) GetEnvironment() []*PersonConfig {
	if m != nil {
		return m.Environment
	}
	return nil
}

type Card struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{1}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

type CardRequest struct {
	Api                  string        `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lang                 string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	PersonConfig         *PersonConfig `protobuf:"bytes,3,opt,name=personConfig,proto3" json:"personConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CardRequest) Reset()         { *m = CardRequest{} }
func (m *CardRequest) String() string { return proto.CompactTextString(m) }
func (*CardRequest) ProtoMessage()    {}
func (*CardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{2}
}

func (m *CardRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardRequest.Unmarshal(m, b)
}
func (m *CardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardRequest.Marshal(b, m, deterministic)
}
func (m *CardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardRequest.Merge(m, src)
}
func (m *CardRequest) XXX_Size() int {
	return xxx_messageInfo_CardRequest.Size(m)
}
func (m *CardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CardRequest proto.InternalMessageInfo

func (m *CardRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CardRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *CardRequest) GetPersonConfig() *PersonConfig {
	if m != nil {
		return m.PersonConfig
	}
	return nil
}

type CardResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Card                 *Card    `protobuf:"bytes,3,opt,name=card,proto3" json:"card,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CardResponse) Reset()         { *m = CardResponse{} }
func (m *CardResponse) String() string { return proto.CompactTextString(m) }
func (*CardResponse) ProtoMessage()    {}
func (*CardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{3}
}

func (m *CardResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardResponse.Unmarshal(m, b)
}
func (m *CardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardResponse.Marshal(b, m, deterministic)
}
func (m *CardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardResponse.Merge(m, src)
}
func (m *CardResponse) XXX_Size() int {
	return xxx_messageInfo_CardResponse.Size(m)
}
func (m *CardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CardResponse proto.InternalMessageInfo

func (m *CardResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *CardResponse) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *CardResponse) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func init() {
	proto.RegisterEnum("predictor.PersonConfig_Gender", PersonConfig_Gender_name, PersonConfig_Gender_value)
	proto.RegisterEnum("predictor.PersonConfig_Features", PersonConfig_Features_name, PersonConfig_Features_value)
	proto.RegisterType((*PersonConfig)(nil), "predictor.PersonConfig")
	proto.RegisterType((*Card)(nil), "predictor.Card")
	proto.RegisterType((*CardRequest)(nil), "predictor.CardRequest")
	proto.RegisterType((*CardResponse)(nil), "predictor.CardResponse")
}

func init() { proto.RegisterFile("api/proto/v1/predictor.proto", fileDescriptor_76009c59104259e6) }

var fileDescriptor_76009c59104259e6 = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x97, 0x34, 0xcb, 0xda, 0xd3, 0x76, 0x84, 0x83, 0xc4, 0xa2, 0x6a, 0x42, 0x55, 0xb8,
	0x29, 0x5c, 0xb4, 0xa2, 0x48, 0x48, 0x80, 0x84, 0xb4, 0x96, 0x14, 0x2a, 0x41, 0x5b, 0x39, 0x45,
	0x08, 0x6e, 0x26, 0xaf, 0xf1, 0x82, 0xa5, 0xcd, 0x0e, 0x8e, 0x57, 0x69, 0xb7, 0xbc, 0x02, 0x2f,
	0xc3, 0x7b, 0xf0, 0x0a, 0x3c, 0x08, 0xb2, 0xd3, 0x96, 0x00, 0x02, 0xed, 0x2a, 0xc7, 0xbf, 0xff,
	0x9c, 0xcf, 0xff, 0xb1, 0xe1, 0x98, 0xe6, 0x7c, 0x90, 0x2b, 0xa9, 0xe5, 0x60, 0xfd, 0x68, 0x90,
	0x2b, 0x96, 0xf2, 0x95, 0x96, 0xaa, 0x6f, 0x25, 0x6c, 0xec, 0x84, 0xce, 0x71, 0x26, 0x65, 0x76,
	0xc1, 0x06, 0xc6, 0x4f, 0x85, 0x90, 0x9a, 0x6a, 0x2e, 0x45, 0x51, 0x1a, 0xa3, 0x6f, 0x2e, 0xb4,
	0x16, 0x4c, 0x15, 0x52, 0x8c, 0xa5, 0x38, 0xe7, 0x19, 0x22, 0x78, 0x82, 0x5e, 0xb2, 0xb0, 0xd6,
	0x75, 0x7a, 0x0d, 0x62, 0x6b, 0x7c, 0x02, 0x7e, 0xc6, 0x44, 0xca, 0x54, 0xe8, 0x75, 0x9d, 0xde,
	0xe1, 0xf0, 0x5e, 0xff, 0x17, 0xaf, 0xfa, 0x73, 0xff, 0x95, 0x75, 0x91, 0x8d, 0x1b, 0x3b, 0x50,
	0x3f, 0xe3, 0x4a, 0x7f, 0x4a, 0xe9, 0x75, 0xb8, 0x6f, 0xfb, 0xed, 0xd6, 0x66, 0xef, 0x9c, 0x51,
	0x7d, 0xa5, 0x58, 0x11, 0xfa, 0x5d, 0xa7, 0xd7, 0x26, 0xbb, 0x35, 0x3e, 0x85, 0x26, 0x13, 0x6b,
	0xae, 0xa4, 0xb8, 0x64, 0x42, 0x87, 0x07, 0xdd, 0x5a, 0xaf, 0x39, 0x3c, 0xfa, 0x07, 0x94, 0x54,
	0xbd, 0xd1, 0x03, 0xf0, 0xcb, 0x43, 0x60, 0x03, 0xf6, 0xe7, 0xcb, 0xd7, 0x31, 0x09, 0xf6, 0xb0,
	0x0e, 0xde, 0xdb, 0x93, 0x37, 0x71, 0xe0, 0x20, 0x80, 0x3f, 0x89, 0x6d, 0xed, 0x46, 0x2f, 0xa0,
	0x3e, 0xd9, 0x12, 0x6f, 0x43, 0x7b, 0x36, 0x5f, 0x9e, 0x26, 0x8b, 0x78, 0x3c, 0x9d, 0x4c, 0xe3,
	0x97, 0xc1, 0x1e, 0x22, 0x1c, 0x8e, 0xde, 0x25, 0xd3, 0x59, 0x9c, 0x24, 0xa7, 0xf3, 0xf7, 0xb3,
	0x98, 0x04, 0x0e, 0x36, 0xe1, 0x60, 0x4c, 0xe2, 0x93, 0xe5, 0x9c, 0x04, 0x6e, 0xe4, 0x83, 0x37,
	0xa6, 0x2a, 0x8d, 0x72, 0x68, 0x9a, 0x2f, 0x61, 0x9f, 0xaf, 0x58, 0xa1, 0x31, 0x80, 0x1a, 0xcd,
	0x79, 0xe8, 0xd8, 0xbc, 0xa6, 0x34, 0x23, 0xbd, 0xa0, 0x22, 0x0b, 0xdd, 0x72, 0xa4, 0xa6, 0xc6,
	0xe7, 0xd0, 0xca, 0x2b, 0x21, 0xec, 0xb8, 0xff, 0x93, 0xf1, 0x37, 0x73, 0xf4, 0x01, 0x5a, 0x25,
	0xb1, 0xc8, 0xa5, 0x28, 0xd8, 0x0d, 0x91, 0xf7, 0xc1, 0x5b, 0x51, 0x95, 0x6e, 0x50, 0xb7, 0x2a,
	0x28, 0xdb, 0xcc, 0x6e, 0x0e, 0x05, 0x04, 0x8b, 0xad, 0x9e, 0x30, 0xb5, 0xe6, 0x2b, 0x86, 0x1f,
	0x01, 0x27, 0x5c, 0xa4, 0xc6, 0x35, 0xba, 0x1e, 0x6d, 0x2f, 0xf0, 0xee, 0x9f, 0x0d, 0xca, 0xfc,
	0x9d, 0xa3, 0xbf, 0xf4, 0xf2, 0x94, 0xd1, 0x9d, 0x2f, 0xdf, 0x7f, 0x7c, 0x75, 0xdb, 0x51, 0xdd,
	0x3c, 0x58, 0x03, 0x7b, 0xe6, 0x3c, 0x3c, 0xf3, 0xed, 0x33, 0x7c, 0xfc, 0x33, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0xd9, 0x72, 0xa4, 0xcf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictorServiceClient is the client API for PredictorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictorServiceClient interface {
	FindCardByBirthday(ctx context.Context, in *CardRequest, opts ...grpc.CallOption) (*CardResponse, error)
}

type predictorServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictorServiceClient(cc *grpc.ClientConn) PredictorServiceClient {
	return &predictorServiceClient{cc}
}

func (c *predictorServiceClient) FindCardByBirthday(ctx context.Context, in *CardRequest, opts ...grpc.CallOption) (*CardResponse, error) {
	out := new(CardResponse)
	err := c.cc.Invoke(ctx, "/predictor.PredictorService/FindCardByBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictorServiceServer is the server API for PredictorService service.
type PredictorServiceServer interface {
	FindCardByBirthday(context.Context, *CardRequest) (*CardResponse, error)
}

func RegisterPredictorServiceServer(s *grpc.Server, srv PredictorServiceServer) {
	s.RegisterService(&_PredictorService_serviceDesc, srv)
}

func _PredictorService_FindCardByBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServiceServer).FindCardByBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predictor.PredictorService/FindCardByBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServiceServer).FindCardByBirthday(ctx, req.(*CardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "predictor.PredictorService",
	HandlerType: (*PredictorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindCardByBirthday",
			Handler:    _PredictorService_FindCardByBirthday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/predictor.proto",
}
