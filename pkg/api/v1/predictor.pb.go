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

type Meaning struct {
	Keywords             string   `protobuf:"bytes,1,opt,name=keywords,proto3" json:"keywords,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meaning) Reset()         { *m = Meaning{} }
func (m *Meaning) String() string { return proto.CompactTextString(m) }
func (*Meaning) ProtoMessage()    {}
func (*Meaning) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{1}
}

func (m *Meaning) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meaning.Unmarshal(m, b)
}
func (m *Meaning) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meaning.Marshal(b, m, deterministic)
}
func (m *Meaning) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meaning.Merge(m, src)
}
func (m *Meaning) XXX_Size() int {
	return xxx_messageInfo_Meaning.Size(m)
}
func (m *Meaning) XXX_DiscardUnknown() {
	xxx_messageInfo_Meaning.DiscardUnknown(m)
}

var xxx_messageInfo_Meaning proto.InternalMessageInfo

func (m *Meaning) GetKeywords() string {
	if m != nil {
		return m.Keywords
	}
	return ""
}

func (m *Meaning) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Card struct {
	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Rank  string `protobuf:"bytes,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Suit  string `protobuf:"bytes,3,opt,name=suit,proto3" json:"suit,omitempty"`
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Meta  string `protobuf:"bytes,5,opt,name=meta,proto3" json:"meta,omitempty"`
	// Only one meaning according to a context (e.g. longterm, pluto and so on)
	Meaning              *Meaning `protobuf:"bytes,6,opt,name=meaning,proto3" json:"meaning,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{2}
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

func (m *Card) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Card) GetRank() string {
	if m != nil {
		return m.Rank
	}
	return ""
}

func (m *Card) GetSuit() string {
	if m != nil {
		return m.Suit
	}
	return ""
}

func (m *Card) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Card) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *Card) GetMeaning() *Meaning {
	if m != nil {
		return m.Meaning
	}
	return nil
}

type Planet struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Symbol               string   `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Planet) Reset()         { *m = Planet{} }
func (m *Planet) String() string { return proto.CompactTextString(m) }
func (*Planet) ProtoMessage()    {}
func (*Planet) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{3}
}

func (m *Planet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Planet.Unmarshal(m, b)
}
func (m *Planet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Planet.Marshal(b, m, deterministic)
}
func (m *Planet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Planet.Merge(m, src)
}
func (m *Planet) XXX_Size() int {
	return xxx_messageInfo_Planet.Size(m)
}
func (m *Planet) XXX_DiscardUnknown() {
	xxx_messageInfo_Planet.DiscardUnknown(m)
}

var xxx_messageInfo_Planet proto.InternalMessageInfo

func (m *Planet) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Planet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Planet) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type PlanetCycleDate struct {
	Month                uint32   `protobuf:"varint,1,opt,name=month,proto3" json:"month,omitempty"`
	Day                  uint32   `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlanetCycleDate) Reset()         { *m = PlanetCycleDate{} }
func (m *PlanetCycleDate) String() string { return proto.CompactTextString(m) }
func (*PlanetCycleDate) ProtoMessage()    {}
func (*PlanetCycleDate) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{4}
}

func (m *PlanetCycleDate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanetCycleDate.Unmarshal(m, b)
}
func (m *PlanetCycleDate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanetCycleDate.Marshal(b, m, deterministic)
}
func (m *PlanetCycleDate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanetCycleDate.Merge(m, src)
}
func (m *PlanetCycleDate) XXX_Size() int {
	return xxx_messageInfo_PlanetCycleDate.Size(m)
}
func (m *PlanetCycleDate) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanetCycleDate.DiscardUnknown(m)
}

var xxx_messageInfo_PlanetCycleDate proto.InternalMessageInfo

func (m *PlanetCycleDate) GetMonth() uint32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *PlanetCycleDate) GetDay() uint32 {
	if m != nil {
		return m.Day
	}
	return 0
}

type PlanetCycle struct {
	Card                 *Card            `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
	Planet               *Planet          `protobuf:"bytes,2,opt,name=planet,proto3" json:"planet,omitempty"`
	Start                *PlanetCycleDate `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	End                  *PlanetCycleDate `protobuf:"bytes,4,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PlanetCycle) Reset()         { *m = PlanetCycle{} }
func (m *PlanetCycle) String() string { return proto.CompactTextString(m) }
func (*PlanetCycle) ProtoMessage()    {}
func (*PlanetCycle) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{5}
}

func (m *PlanetCycle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlanetCycle.Unmarshal(m, b)
}
func (m *PlanetCycle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlanetCycle.Marshal(b, m, deterministic)
}
func (m *PlanetCycle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlanetCycle.Merge(m, src)
}
func (m *PlanetCycle) XXX_Size() int {
	return xxx_messageInfo_PlanetCycle.Size(m)
}
func (m *PlanetCycle) XXX_DiscardUnknown() {
	xxx_messageInfo_PlanetCycle.DiscardUnknown(m)
}

var xxx_messageInfo_PlanetCycle proto.InternalMessageInfo

func (m *PlanetCycle) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *PlanetCycle) GetPlanet() *Planet {
	if m != nil {
		return m.Planet
	}
	return nil
}

func (m *PlanetCycle) GetStart() *PlanetCycleDate {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *PlanetCycle) GetEnd() *PlanetCycleDate {
	if m != nil {
		return m.End
	}
	return nil
}

type GeneralRequest struct {
	// TODO: request for specific year
	Api                  string        `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lang                 string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	PersonConfig         *PersonConfig `protobuf:"bytes,3,opt,name=personConfig,proto3" json:"personConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GeneralRequest) Reset()         { *m = GeneralRequest{} }
func (m *GeneralRequest) String() string { return proto.CompactTextString(m) }
func (*GeneralRequest) ProtoMessage()    {}
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{6}
}

func (m *GeneralRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GeneralRequest.Unmarshal(m, b)
}
func (m *GeneralRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GeneralRequest.Marshal(b, m, deterministic)
}
func (m *GeneralRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeneralRequest.Merge(m, src)
}
func (m *GeneralRequest) XXX_Size() int {
	return xxx_messageInfo_GeneralRequest.Size(m)
}
func (m *GeneralRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeneralRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GeneralRequest proto.InternalMessageInfo

func (m *GeneralRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *GeneralRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *GeneralRequest) GetPersonConfig() *PersonConfig {
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
	return fileDescriptor_76009c59104259e6, []int{7}
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

type PredictionResponse struct {
	Api                  string                  `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lang                 string                  `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	PlanetCycles         map[string]*PlanetCycle `protobuf:"bytes,3,rep,name=planetCycles,proto3" json:"planetCycles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Cards                map[string]*Card        `protobuf:"bytes,4,rep,name=cards,proto3" json:"cards,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *PredictionResponse) Reset()         { *m = PredictionResponse{} }
func (m *PredictionResponse) String() string { return proto.CompactTextString(m) }
func (*PredictionResponse) ProtoMessage()    {}
func (*PredictionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{8}
}

func (m *PredictionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PredictionResponse.Unmarshal(m, b)
}
func (m *PredictionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PredictionResponse.Marshal(b, m, deterministic)
}
func (m *PredictionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PredictionResponse.Merge(m, src)
}
func (m *PredictionResponse) XXX_Size() int {
	return xxx_messageInfo_PredictionResponse.Size(m)
}
func (m *PredictionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PredictionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PredictionResponse proto.InternalMessageInfo

func (m *PredictionResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *PredictionResponse) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *PredictionResponse) GetPlanetCycles() map[string]*PlanetCycle {
	if m != nil {
		return m.PlanetCycles
	}
	return nil
}

func (m *PredictionResponse) GetCards() map[string]*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto.RegisterEnum("predictor.PersonConfig_Gender", PersonConfig_Gender_name, PersonConfig_Gender_value)
	proto.RegisterEnum("predictor.PersonConfig_Features", PersonConfig_Features_name, PersonConfig_Features_value)
	proto.RegisterType((*PersonConfig)(nil), "predictor.PersonConfig")
	proto.RegisterType((*Meaning)(nil), "predictor.Meaning")
	proto.RegisterType((*Card)(nil), "predictor.Card")
	proto.RegisterType((*Planet)(nil), "predictor.Planet")
	proto.RegisterType((*PlanetCycleDate)(nil), "predictor.PlanetCycleDate")
	proto.RegisterType((*PlanetCycle)(nil), "predictor.PlanetCycle")
	proto.RegisterType((*GeneralRequest)(nil), "predictor.GeneralRequest")
	proto.RegisterType((*CardResponse)(nil), "predictor.CardResponse")
	proto.RegisterType((*PredictionResponse)(nil), "predictor.PredictionResponse")
	proto.RegisterMapType((map[string]*Card)(nil), "predictor.PredictionResponse.CardsEntry")
	proto.RegisterMapType((map[string]*PlanetCycle)(nil), "predictor.PredictionResponse.PlanetCyclesEntry")
}

func init() { proto.RegisterFile("api/proto/v1/predictor.proto", fileDescriptor_76009c59104259e6) }

var fileDescriptor_76009c59104259e6 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0xd1, 0x6a, 0xeb, 0x46,
	0x10, 0x8d, 0x64, 0x5b, 0x8e, 0x47, 0x8e, 0xa3, 0x2c, 0x21, 0x55, 0x4d, 0x5a, 0x8c, 0x4a, 0xc1,
	0x29, 0x21, 0x6e, 0x5d, 0x28, 0x4d, 0x0a, 0x81, 0xd4, 0x51, 0x5c, 0x43, 0x13, 0x9b, 0x75, 0x4a,
	0xe8, 0x53, 0xd8, 0x58, 0x1b, 0x47, 0x58, 0x5e, 0xa9, 0xab, 0xb5, 0x8b, 0xa1, 0x4f, 0xfd, 0x85,
	0x42, 0xbf, 0xa5, 0x2f, 0xfd, 0x8a, 0xfe, 0x42, 0xff, 0xe0, 0xfe, 0xc0, 0x65, 0x57, 0x92, 0x2d,
	0xe3, 0xc4, 0xdc, 0xfb, 0x36, 0xb3, 0x73, 0xf6, 0xcc, 0x99, 0xb3, 0x23, 0x04, 0xc7, 0x24, 0xf2,
	0x5b, 0x11, 0x0f, 0x45, 0xd8, 0x9a, 0x7f, 0xd3, 0x8a, 0x38, 0xf5, 0xfc, 0x91, 0x08, 0xf9, 0x99,
	0x3a, 0x42, 0x95, 0xe5, 0x41, 0xfd, 0x78, 0x1c, 0x86, 0xe3, 0x80, 0xb6, 0x24, 0x9e, 0x30, 0x16,
	0x0a, 0x22, 0xfc, 0x90, 0xc5, 0x09, 0xd0, 0xf9, 0x47, 0x87, 0xea, 0x80, 0xf2, 0x38, 0x64, 0x9d,
	0x90, 0x3d, 0xfb, 0x63, 0x84, 0xa0, 0xc8, 0xc8, 0x94, 0xda, 0x85, 0x86, 0xd6, 0xac, 0x60, 0x15,
	0xa3, 0xef, 0xc0, 0x18, 0x53, 0xe6, 0x51, 0x6e, 0x17, 0x1b, 0x5a, 0xb3, 0xd6, 0xfe, 0xfc, 0x6c,
	0xd5, 0x2f, 0x7f, 0xf9, 0xac, 0xab, 0x50, 0x38, 0x45, 0xa3, 0x3a, 0xec, 0x3e, 0xf9, 0x5c, 0xbc,
	0x78, 0x64, 0x61, 0x97, 0x14, 0xdf, 0x32, 0x97, 0xb5, 0x67, 0x4a, 0xc4, 0x8c, 0xd3, 0xd8, 0x36,
	0x1a, 0x5a, 0x73, 0x0f, 0x2f, 0x73, 0x74, 0x0e, 0x26, 0x65, 0x73, 0x9f, 0x87, 0x6c, 0x4a, 0x99,
	0xb0, 0xcb, 0x8d, 0x42, 0xd3, 0x6c, 0x7f, 0xf2, 0x46, 0x53, 0x9c, 0xc7, 0x3a, 0x27, 0x60, 0x24,
	0x22, 0x50, 0x05, 0x4a, 0xfd, 0xfb, 0x9f, 0x5c, 0x6c, 0xed, 0xa0, 0x5d, 0x28, 0xde, 0x5e, 0xfd,
	0xec, 0x5a, 0x1a, 0x02, 0x30, 0x6e, 0x5c, 0x15, 0xeb, 0xce, 0x25, 0xec, 0xde, 0x64, 0x1d, 0x0f,
	0x60, 0xef, 0xae, 0x7f, 0xff, 0x38, 0x1c, 0xb8, 0x9d, 0xde, 0x4d, 0xcf, 0xbd, 0xb6, 0x76, 0x10,
	0x82, 0xda, 0x8f, 0xbf, 0x0c, 0x7b, 0x77, 0xee, 0x70, 0xf8, 0xd8, 0x7f, 0xb8, 0x73, 0xb1, 0xa5,
	0x21, 0x13, 0xca, 0x1d, 0xec, 0x5e, 0xdd, 0xf7, 0xb1, 0xa5, 0x3b, 0x5d, 0x28, 0xdf, 0x52, 0xc2,
	0x7c, 0x36, 0x96, 0xc3, 0x4c, 0xe8, 0xe2, 0xf7, 0x90, 0x7b, 0xb1, 0xad, 0x25, 0x83, 0x66, 0x39,
	0x6a, 0x80, 0xe9, 0xd1, 0x78, 0xc4, 0xfd, 0x48, 0xfa, 0x6e, 0xeb, 0xaa, 0x9c, 0x3f, 0x72, 0xfe,
	0xd6, 0xa0, 0xd8, 0x21, 0xdc, 0x43, 0x35, 0xd0, 0x7d, 0x4f, 0x11, 0xec, 0x61, 0xdd, 0xf7, 0xe4,
	0x5b, 0x70, 0xc2, 0x26, 0xe9, 0x1d, 0x15, 0xcb, 0xb3, 0x78, 0xe6, 0x8b, 0xec, 0x7d, 0x64, 0x8c,
	0x0e, 0xa1, 0x24, 0x7c, 0x11, 0x50, 0xf5, 0x3c, 0x15, 0x9c, 0x24, 0x12, 0x39, 0xa5, 0x82, 0xa4,
	0xce, 0xab, 0x18, 0x9d, 0x42, 0x79, 0x9a, 0x68, 0x56, 0xa6, 0x9b, 0x6d, 0x94, 0x73, 0x35, 0x9d,
	0x06, 0x67, 0x10, 0xe7, 0x1a, 0x8c, 0x41, 0x40, 0x18, 0x15, 0xaf, 0x29, 0x53, 0x5b, 0xa2, 0xe7,
	0xb6, 0xe4, 0x08, 0x8c, 0x78, 0x31, 0x7d, 0x0a, 0x83, 0x54, 0x5b, 0x9a, 0x39, 0xe7, 0xb0, 0x9f,
	0xb0, 0x74, 0x16, 0xa3, 0x80, 0x5e, 0x13, 0x41, 0xa5, 0xe0, 0x69, 0xc8, 0xc4, 0x4b, 0xca, 0x98,
	0x24, 0xc8, 0x82, 0x82, 0xdc, 0x14, 0x5d, 0x9d, 0xc9, 0xd0, 0xf9, 0x57, 0x03, 0x33, 0x77, 0x17,
	0x7d, 0x01, 0xc5, 0x11, 0xe1, 0x89, 0x10, 0xb3, 0xbd, 0x9f, 0xd3, 0x2e, 0xfd, 0xc3, 0xaa, 0x88,
	0x4e, 0xc0, 0x88, 0xd4, 0x1d, 0xc5, 0x64, 0xb6, 0x0f, 0xf2, 0x8b, 0xa3, 0x0a, 0x38, 0x05, 0xa0,
	0xaf, 0xa1, 0x14, 0x0b, 0xc2, 0x13, 0x37, 0xcd, 0x76, 0x7d, 0x03, 0xb9, 0x94, 0x8c, 0x13, 0x20,
	0x3a, 0x85, 0x02, 0x65, 0x9e, 0x32, 0x7a, 0x3b, 0x5e, 0xc2, 0x9c, 0x18, 0x6a, 0x5d, 0xca, 0x28,
	0x27, 0x01, 0xa6, 0xbf, 0xcd, 0x68, 0x2c, 0xe4, 0x8c, 0x24, 0xf2, 0xd3, 0x25, 0x91, 0xa1, 0xb4,
	0x32, 0x20, 0x6c, 0x9c, 0x59, 0x29, 0x63, 0xf4, 0x03, 0x54, 0xa3, 0xdc, 0x8a, 0xa7, 0xf2, 0xde,
	0xfc, 0x02, 0xd6, 0xc0, 0xce, 0xaf, 0x50, 0x55, 0x6e, 0xd0, 0x38, 0x0a, 0x59, 0x4c, 0x3f, 0xb0,
	0x65, 0x66, 0x6d, 0x61, 0x8b, 0xb5, 0xce, 0x3b, 0x1d, 0xd0, 0x20, 0x29, 0xf8, 0x21, 0xfb, 0xc8,
	0x0e, 0x43, 0xa8, 0x46, 0x2b, 0x93, 0x62, 0xbb, 0xa0, 0x3e, 0xeb, 0x56, 0x7e, 0xa8, 0x0d, 0xea,
	0xbc, 0xad, 0xb1, 0xcb, 0x04, 0x5f, 0xe0, 0x35, 0x12, 0x74, 0x09, 0x25, 0xa9, 0x2c, 0xb6, 0x8b,
	0x8a, 0xad, 0xb9, 0x9d, 0x4d, 0x8e, 0x92, 0xd2, 0x24, 0xd7, 0xea, 0x0f, 0x70, 0xb0, 0xd1, 0x42,
	0xce, 0x33, 0xa1, 0x8b, 0x6c, 0x9e, 0x09, 0x5d, 0xa0, 0x53, 0x28, 0xcd, 0x49, 0x30, 0xa3, 0xe9,
	0x4a, 0x1d, 0xbd, 0xfe, 0xf0, 0x38, 0x01, 0x5d, 0xe8, 0xdf, 0x6b, 0xf5, 0x1e, 0xc0, 0xaa, 0xdb,
	0x2b, 0x8c, 0x5f, 0xae, 0x33, 0x6e, 0x18, 0xbe, 0xa2, 0x6a, 0xff, 0x01, 0xd6, 0x20, 0x2b, 0x0e,
	0x29, 0x9f, 0xfb, 0x23, 0x8a, 0x5e, 0xe0, 0xb0, 0x4b, 0x45, 0xba, 0x5c, 0xab, 0x49, 0xd1, 0xa7,
	0x39, 0x9e, 0xf5, 0xd5, 0xab, 0x7f, 0xb6, 0xd5, 0x1b, 0xe7, 0xe8, 0xcf, 0xff, 0xfe, 0xff, 0x4b,
	0xb7, 0x1c, 0x33, 0xf7, 0x3f, 0xb9, 0xd0, 0xbe, 0x7a, 0x32, 0xd4, 0x8f, 0xe2, 0xdb, 0xf7, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x20, 0x74, 0xfd, 0xe8, 0x71, 0x06, 0x00, 0x00,
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
	GetGeneralPrediction(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*PredictionResponse, error)
}

type predictorServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictorServiceClient(cc *grpc.ClientConn) PredictorServiceClient {
	return &predictorServiceClient{cc}
}

func (c *predictorServiceClient) GetGeneralPrediction(ctx context.Context, in *GeneralRequest, opts ...grpc.CallOption) (*PredictionResponse, error) {
	out := new(PredictionResponse)
	err := c.cc.Invoke(ctx, "/predictor.PredictorService/GetGeneralPrediction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictorServiceServer is the server API for PredictorService service.
type PredictorServiceServer interface {
	GetGeneralPrediction(context.Context, *GeneralRequest) (*PredictionResponse, error)
}

func RegisterPredictorServiceServer(s *grpc.Server, srv PredictorServiceServer) {
	s.RegisterService(&_PredictorService_serviceDesc, srv)
}

func _PredictorService_GetGeneralPrediction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeneralRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServiceServer).GetGeneralPrediction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/predictor.PredictorService/GetGeneralPrediction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServiceServer).GetGeneralPrediction(ctx, req.(*GeneralRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "predictor.PredictorService",
	HandlerType: (*PredictorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGeneralPrediction",
			Handler:    _PredictorService_GetGeneralPrediction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/predictor.proto",
}
