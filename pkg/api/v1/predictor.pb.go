// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/v1/predictor.proto

package predictor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

type Date struct {
	Year                 uint32   `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month                uint32   `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day                  uint32   `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Date) Reset()         { *m = Date{} }
func (m *Date) String() string { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()    {}
func (*Date) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{0}
}

func (m *Date) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Date.Unmarshal(m, b)
}
func (m *Date) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Date.Marshal(b, m, deterministic)
}
func (m *Date) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Date.Merge(m, src)
}
func (m *Date) XXX_Size() int {
	return xxx_messageInfo_Date.Size(m)
}
func (m *Date) XXX_DiscardUnknown() {
	xxx_messageInfo_Date.DiscardUnknown(m)
}

var xxx_messageInfo_Date proto.InternalMessageInfo

func (m *Date) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *Date) GetMonth() uint32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *Date) GetDay() uint32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type Card struct {
	Number               uint32   `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Suite                string   `protobuf:"bytes,2,opt,name=suite,proto3" json:"suite,omitempty"`
	Rank                 string   `protobuf:"bytes,3,opt,name=rank,proto3" json:"rank,omitempty"`
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

func (m *Card) GetNumber() uint32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Card) GetSuite() string {
	if m != nil {
		return m.Suite
	}
	return ""
}

func (m *Card) GetRank() string {
	if m != nil {
		return m.Rank
	}
	return ""
}

type Matrix struct {
	M                    *_struct.ListValue `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Matrix) Reset()         { *m = Matrix{} }
func (m *Matrix) String() string { return proto.CompactTextString(m) }
func (*Matrix) ProtoMessage()    {}
func (*Matrix) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{2}
}

func (m *Matrix) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Matrix.Unmarshal(m, b)
}
func (m *Matrix) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Matrix.Marshal(b, m, deterministic)
}
func (m *Matrix) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Matrix.Merge(m, src)
}
func (m *Matrix) XXX_Size() int {
	return xxx_messageInfo_Matrix.Size(m)
}
func (m *Matrix) XXX_DiscardUnknown() {
	xxx_messageInfo_Matrix.DiscardUnknown(m)
}

var xxx_messageInfo_Matrix proto.InternalMessageInfo

func (m *Matrix) GetM() *_struct.ListValue {
	if m != nil {
		return m.M
	}
	return nil
}

func init() {
	proto.RegisterType((*Date)(nil), "predictor.Date")
	proto.RegisterType((*Card)(nil), "predictor.Card")
	proto.RegisterType((*Matrix)(nil), "predictor.Matrix")
}

func init() { proto.RegisterFile("api/proto/v1/predictor.proto", fileDescriptor_76009c59104259e6) }

var fileDescriptor_76009c59104259e6 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xe5, 0xb6, 0x7f, 0xf5, 0xc7, 0xa8, 0x82, 0x1a, 0x54, 0x55, 0xa1, 0x0b, 0x94, 0x55,
	0xd5, 0x45, 0x22, 0xca, 0xae, 0xcb, 0x00, 0x2a, 0x42, 0x20, 0x50, 0x16, 0xec, 0xa7, 0x89, 0x69,
	0x2d, 0x1a, 0x3b, 0x72, 0x26, 0x88, 0x6c, 0xb9, 0x02, 0x97, 0xe0, 0x3e, 0x5c, 0x81, 0x83, 0x20,
	0xdb, 0x0d, 0x94, 0xb2, 0x9b, 0x79, 0xf3, 0xfc, 0x4d, 0xe6, 0x85, 0x8e, 0xa0, 0x10, 0x51, 0xa1,
	0x15, 0xaa, 0xe8, 0xf9, 0x34, 0x2a, 0x34, 0xcf, 0x44, 0x8a, 0x4a, 0x87, 0x56, 0x62, 0xde, 0xb7,
	0xe0, 0x1f, 0x2f, 0x95, 0x5a, 0xae, 0xb9, 0xf3, 0x2e, 0xaa, 0xc7, 0x88, 0xe7, 0x05, 0xd6, 0xce,
	0xe7, 0x8f, 0x76, 0x87, 0x25, 0xea, 0x2a, 0xc5, 0x9d, 0xa9, 0x59, 0x05, 0x52, 0x2a, 0x04, 0x14,
	0x4a, 0x96, 0x6e, 0x1a, 0xc4, 0xb4, 0x73, 0x01, 0xc8, 0x19, 0xa3, 0x9d, 0x9a, 0x83, 0x1e, 0x92,
	0x13, 0x32, 0xee, 0x25, 0xb6, 0x66, 0x47, 0xf4, 0x5f, 0xae, 0x24, 0xae, 0x86, 0x2d, 0x2b, 0xba,
	0x86, 0x1d, 0xd0, 0x76, 0x06, 0xf5, 0xb0, 0x6d, 0x35, 0x53, 0x06, 0x57, 0xb4, 0x73, 0x0e, 0x3a,
	0x63, 0x03, 0xda, 0x95, 0x55, 0xbe, 0xe0, 0x0d, 0x65, 0xd3, 0x19, 0x4e, 0x59, 0x09, 0xe4, 0x96,
	0xe3, 0x25, 0xae, 0x31, 0x1b, 0x35, 0xc8, 0x27, 0x0b, 0xf2, 0x12, 0x5b, 0x07, 0x53, 0xda, 0xbd,
	0x05, 0xd4, 0xe2, 0x85, 0x8d, 0x29, 0xc9, 0x2d, 0x66, 0x6f, 0xea, 0x87, 0xee, 0x82, 0xb0, 0xb9,
	0x2f, 0xbc, 0x11, 0x25, 0x3e, 0xc0, 0xba, 0xe2, 0x09, 0xc9, 0xa7, 0xef, 0x84, 0x7a, 0xf7, 0x4d,
	0x50, 0xec, 0x8e, 0xf6, 0xe6, 0x1c, 0x63, 0x28, 0xf9, 0x06, 0x34, 0xf8, 0xf3, 0xfa, 0xd2, 0x44,
	0xe7, 0xf7, 0xc3, 0x9f, 0xb8, 0x9d, 0x35, 0x38, 0x7c, 0xfd, 0xf8, 0x7c, 0x6b, 0xf5, 0x82, 0xff,
	0xe6, 0x67, 0xf0, 0x74, 0xa5, 0x66, 0x64, 0xc2, 0xae, 0x69, 0x7f, 0xce, 0xd1, 0xdc, 0x17, 0xd7,
	0xb1, 0xd0, 0xb8, 0xca, 0xa0, 0x66, 0xfb, 0x5b, 0x8f, 0x4d, 0x7c, 0xfe, 0xb6, 0x60, 0xbc, 0xbf,
	0x59, 0x29, 0xe8, 0x6c, 0x46, 0x26, 0x8b, 0xae, 0xfd, 0x86, 0xb3, 0xaf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x25, 0xfe, 0x3b, 0xf7, 0xf7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PredictorClient is the client API for Predictor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PredictorClient interface {
	GetBaseMatrix(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Matrix, error)
	FindCardByBirthday(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Card, error)
}

type predictorClient struct {
	cc *grpc.ClientConn
}

func NewPredictorClient(cc *grpc.ClientConn) PredictorClient {
	return &predictorClient{cc}
}

func (c *predictorClient) GetBaseMatrix(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Matrix, error) {
	out := new(Matrix)
	err := c.cc.Invoke(ctx, "/predictor.Predictor/GetBaseMatrix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *predictorClient) FindCardByBirthday(ctx context.Context, in *Date, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, "/predictor.Predictor/FindCardByBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictorServer is the server API for Predictor service.
type PredictorServer interface {
	GetBaseMatrix(context.Context, *empty.Empty) (*Matrix, error)
	FindCardByBirthday(context.Context, *Date) (*Card, error)
}

func RegisterPredictorServer(s *grpc.Server, srv PredictorServer) {
	s.RegisterService(&_Predictor_serviceDesc, srv)
}

func _Predictor_GetBaseMatrix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServer).GetBaseMatrix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predictor.Predictor/GetBaseMatrix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServer).GetBaseMatrix(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Predictor_FindCardByBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Date)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServer).FindCardByBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predictor.Predictor/FindCardByBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServer).FindCardByBirthday(ctx, req.(*Date))
	}
	return interceptor(ctx, in, info, handler)
}

var _Predictor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "predictor.Predictor",
	HandlerType: (*PredictorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBaseMatrix",
			Handler:    _Predictor_GetBaseMatrix_Handler,
		},
		{
			MethodName: "FindCardByBirthday",
			Handler:    _Predictor_FindCardByBirthday_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/predictor.proto",
}
