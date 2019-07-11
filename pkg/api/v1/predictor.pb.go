// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/v1/predictor.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type Gender int32

const (
	Gender_OTHER  Gender = 0
	Gender_MALE   Gender = 1
	Gender_FEMALE Gender = 2
)

var Gender_name = map[int32]string{
	0: "OTHER",
	1: "MALE",
	2: "FEMALE",
}

var Gender_value = map[string]int32{
	"OTHER":  0,
	"MALE":   1,
	"FEMALE": 2,
}

func (x Gender) String() string {
	return proto.EnumName(Gender_name, int32(x))
}

func (Gender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{0}
}

type Features int32

const (
	Features_NOT_SPECIFIED  Features = 0
	Features_BUSINESS_OWNER Features = 1
	Features_CREATOR        Features = 2
)

var Features_name = map[int32]string{
	0: "NOT_SPECIFIED",
	1: "BUSINESS_OWNER",
	2: "CREATOR",
}

var Features_value = map[string]int32{
	"NOT_SPECIFIED":  0,
	"BUSINESS_OWNER": 1,
	"CREATOR":        2,
}

func (x Features) String() string {
	return proto.EnumName(Features_name, int32(x))
}

func (Features) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{1}
}

type PersonConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Gender   `protobuf:"varint,2,opt,name=gender,proto3,enum=v1.Gender" json:"gender,omitempty"`
	Birthday             string   `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Features             uint32   `protobuf:"varint,4,opt,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *PersonConfig) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_OTHER
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

type Person struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Gender               Gender                  `protobuf:"varint,2,opt,name=gender,proto3,enum=v1.Gender" json:"gender,omitempty"`
	Birthday             string                  `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	PlanetCycles         map[string]*PlanetCycle `protobuf:"bytes,4,rep,name=planet_cycles,json=planetCycles,proto3" json:"planet_cycles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BaseCards            map[string]*Card        `protobuf:"bytes,5,rep,name=base_cards,json=baseCards,proto3" json:"base_cards,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PersonalCards        map[string]*Card        `protobuf:"bytes,6,rep,name=personal_cards,json=personalCards,proto3" json:"personal_cards,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{1}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetGender() Gender {
	if m != nil {
		return m.Gender
	}
	return Gender_OTHER
}

func (m *Person) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *Person) GetPlanetCycles() map[string]*PlanetCycle {
	if m != nil {
		return m.PlanetCycles
	}
	return nil
}

func (m *Person) GetBaseCards() map[string]*Card {
	if m != nil {
		return m.BaseCards
	}
	return nil
}

func (m *Person) GetPersonalCards() map[string]*Card {
	if m != nil {
		return m.PersonalCards
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
	return fileDescriptor_76009c59104259e6, []int{2}
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
	// only one meaning according to a context (e.g. longterm, pluto and so on)
	Meaning              *Meaning `protobuf:"bytes,6,opt,name=meaning,proto3" json:"meaning,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{3}
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
	return fileDescriptor_76009c59104259e6, []int{4}
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
	return fileDescriptor_76009c59104259e6, []int{5}
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
	return fileDescriptor_76009c59104259e6, []int{6}
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

type PersonRequest struct {
	Api                  string        `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lang                 string        `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Year                 uint32        `protobuf:"varint,3,opt,name=year,proto3" json:"year,omitempty"`
	PersonConfig         *PersonConfig `protobuf:"bytes,4,opt,name=person_config,json=personConfig,proto3" json:"person_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *PersonRequest) Reset()         { *m = PersonRequest{} }
func (m *PersonRequest) String() string { return proto.CompactTextString(m) }
func (*PersonRequest) ProtoMessage()    {}
func (*PersonRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{7}
}

func (m *PersonRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonRequest.Unmarshal(m, b)
}
func (m *PersonRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonRequest.Marshal(b, m, deterministic)
}
func (m *PersonRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonRequest.Merge(m, src)
}
func (m *PersonRequest) XXX_Size() int {
	return xxx_messageInfo_PersonRequest.Size(m)
}
func (m *PersonRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PersonRequest proto.InternalMessageInfo

func (m *PersonRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *PersonRequest) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *PersonRequest) GetYear() uint32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *PersonRequest) GetPersonConfig() *PersonConfig {
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
	return fileDescriptor_76009c59104259e6, []int{8}
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

type PersonResponse struct {
	Api                  string   `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Lang                 string   `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Person               *Person  `protobuf:"bytes,3,opt,name=person,proto3" json:"person,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonResponse) Reset()         { *m = PersonResponse{} }
func (m *PersonResponse) String() string { return proto.CompactTextString(m) }
func (*PersonResponse) ProtoMessage()    {}
func (*PersonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_76009c59104259e6, []int{9}
}

func (m *PersonResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonResponse.Unmarshal(m, b)
}
func (m *PersonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonResponse.Marshal(b, m, deterministic)
}
func (m *PersonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonResponse.Merge(m, src)
}
func (m *PersonResponse) XXX_Size() int {
	return xxx_messageInfo_PersonResponse.Size(m)
}
func (m *PersonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PersonResponse proto.InternalMessageInfo

func (m *PersonResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *PersonResponse) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *PersonResponse) GetPerson() *Person {
	if m != nil {
		return m.Person
	}
	return nil
}

func init() {
	proto.RegisterEnum("v1.Gender", Gender_name, Gender_value)
	proto.RegisterEnum("v1.Features", Features_name, Features_value)
	proto.RegisterType((*PersonConfig)(nil), "v1.PersonConfig")
	proto.RegisterType((*Person)(nil), "v1.Person")
	proto.RegisterMapType((map[string]*Card)(nil), "v1.Person.BaseCardsEntry")
	proto.RegisterMapType((map[string]*Card)(nil), "v1.Person.PersonalCardsEntry")
	proto.RegisterMapType((map[string]*PlanetCycle)(nil), "v1.Person.PlanetCyclesEntry")
	proto.RegisterType((*Meaning)(nil), "v1.Meaning")
	proto.RegisterType((*Card)(nil), "v1.Card")
	proto.RegisterType((*Planet)(nil), "v1.Planet")
	proto.RegisterType((*PlanetCycleDate)(nil), "v1.PlanetCycleDate")
	proto.RegisterType((*PlanetCycle)(nil), "v1.PlanetCycle")
	proto.RegisterType((*PersonRequest)(nil), "v1.PersonRequest")
	proto.RegisterType((*CardResponse)(nil), "v1.CardResponse")
	proto.RegisterType((*PersonResponse)(nil), "v1.PersonResponse")
}

func init() { proto.RegisterFile("api/proto/v1/predictor.proto", fileDescriptor_76009c59104259e6) }

var fileDescriptor_76009c59104259e6 = []byte{
	// 961 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x51, 0x6f, 0x23, 0x35,
	0x10, 0xbe, 0xdd, 0x24, 0xdb, 0x76, 0xd2, 0xe4, 0x52, 0x83, 0x4e, 0xb9, 0xa8, 0xa0, 0xd5, 0x4a,
	0x95, 0x7a, 0x15, 0x97, 0x5c, 0x43, 0x91, 0x8e, 0x22, 0x21, 0xda, 0x34, 0x3d, 0x0a, 0x5c, 0x5b,
	0xb9, 0x45, 0x48, 0xbc, 0x54, 0xee, 0xee, 0x5c, 0xb2, 0xd7, 0xc4, 0x5e, 0x6c, 0x27, 0x25, 0x42,
	0x20, 0x84, 0xc4, 0x33, 0x12, 0xbc, 0xf1, 0xc6, 0xef, 0xe1, 0x81, 0x07, 0xfe, 0x02, 0x3f, 0x04,
	0xd9, 0xde, 0xa4, 0x7b, 0xb4, 0x48, 0x80, 0xc4, 0xd3, 0xce, 0x7c, 0x33, 0xf3, 0xcd, 0x78, 0x66,
	0xd6, 0x86, 0x75, 0x96, 0xa5, 0x9d, 0x4c, 0x0a, 0x2d, 0x3a, 0xd3, 0xed, 0x4e, 0x26, 0x31, 0x49,
	0x63, 0x2d, 0x64, 0xdb, 0x42, 0xc4, 0x9f, 0x6e, 0xb7, 0xd6, 0x07, 0x42, 0x0c, 0x46, 0xd8, 0x31,
	0x8e, 0x8c, 0x73, 0xa1, 0x99, 0x4e, 0x05, 0x57, 0xce, 0xa3, 0xf5, 0x96, 0xfd, 0xc4, 0x8f, 0x07,
	0xc8, 0x1f, 0xab, 0x6b, 0x36, 0x18, 0xa0, 0xec, 0x88, 0xcc, 0x7a, 0xdc, 0xf6, 0x8e, 0xbe, 0x81,
	0xd5, 0x53, 0x94, 0x4a, 0xf0, 0x9e, 0xe0, 0x2f, 0xd2, 0x01, 0x21, 0x50, 0xe6, 0x6c, 0x8c, 0x4d,
	0x2f, 0xf4, 0x36, 0x57, 0xa8, 0x95, 0x49, 0x04, 0xc1, 0x00, 0x79, 0x82, 0xb2, 0xe9, 0x87, 0xde,
	0x66, 0xbd, 0x0b, 0xed, 0xe9, 0x76, 0xfb, 0x99, 0x45, 0x68, 0x6e, 0x21, 0x2d, 0x58, 0xbe, 0x4c,
	0xa5, 0x1e, 0x26, 0x6c, 0xd6, 0x2c, 0xd9, 0xd8, 0x85, 0x6e, 0x6c, 0x2f, 0x90, 0xe9, 0x89, 0x44,
	0xd5, 0x2c, 0x87, 0xde, 0x66, 0x8d, 0x2e, 0xf4, 0xe8, 0xfb, 0x32, 0x04, 0xae, 0x80, 0xff, 0x25,
	0xf5, 0x1e, 0xd4, 0xb2, 0x11, 0xe3, 0xa8, 0x2f, 0xe2, 0x59, 0x3c, 0xb2, 0xf9, 0x4b, 0x9b, 0xd5,
	0xee, 0xba, 0xa1, 0x71, 0x69, 0xdb, 0xa7, 0xd6, 0xde, 0xb3, 0xe6, 0x3e, 0xd7, 0x72, 0x46, 0x57,
	0xb3, 0x02, 0x44, 0x9e, 0x02, 0x5c, 0x32, 0x85, 0x17, 0x31, 0x93, 0x89, 0x6a, 0x56, 0x6c, 0xfc,
	0xc3, 0x42, 0xfc, 0x3e, 0x53, 0xd8, 0x33, 0x36, 0x17, 0xbc, 0x72, 0x39, 0xd7, 0xc9, 0x01, 0xd4,
	0x33, 0xeb, 0xc3, 0x46, 0x79, 0x74, 0x60, 0xa3, 0xdf, 0x28, 0x66, 0xcf, 0x1d, 0x0a, 0x0c, 0xb5,
	0xac, 0x88, 0xb5, 0x4e, 0x61, 0xed, 0x56, 0x89, 0xa4, 0x01, 0xa5, 0x2b, 0x9c, 0xe5, 0xad, 0x32,
	0x22, 0xd9, 0x80, 0xca, 0x94, 0x8d, 0x26, 0x68, 0x1b, 0x55, 0xed, 0xde, 0xb7, 0x39, 0x6e, 0xe2,
	0xa8, 0xb3, 0xee, 0xfa, 0x4f, 0xbd, 0xd6, 0x21, 0xd4, 0x5f, 0x2d, 0xfa, 0x0e, 0xba, 0x37, 0x5f,
	0xa5, 0x5b, 0x36, 0x74, 0x26, 0xa0, 0xc8, 0xf3, 0x11, 0x90, 0xdb, 0xe5, 0xff, 0x37, 0xae, 0xe8,
	0x19, 0x2c, 0x3d, 0x47, 0xc6, 0x53, 0x3e, 0x30, 0xf3, 0xbc, 0xc2, 0xd9, 0xb5, 0x30, 0x0d, 0x73,
	0x2c, 0x0b, 0x9d, 0x84, 0x50, 0x4d, 0x50, 0xc5, 0x32, 0xb5, 0x0b, 0x6d, 0x09, 0x57, 0x68, 0x11,
	0x8a, 0x7e, 0xf0, 0xa0, 0x6c, 0xc8, 0x49, 0x1d, 0xfc, 0x34, 0xb1, 0x04, 0x35, 0xea, 0xa7, 0x89,
	0x59, 0x2f, 0xc9, 0xf8, 0x55, 0x1e, 0x63, 0x65, 0x83, 0xa9, 0x49, 0xaa, 0xf3, 0xb5, 0xb1, 0x32,
	0x79, 0x1d, 0x2a, 0x3a, 0xd5, 0x23, 0xb4, 0xab, 0xba, 0x42, 0x9d, 0x62, 0x3c, 0xc7, 0xa8, 0x59,
	0xb3, 0xe2, 0x3c, 0x8d, 0x4c, 0x36, 0x60, 0x69, 0xec, 0x6a, 0x6e, 0x06, 0xf6, 0x64, 0x55, 0x73,
	0xb2, 0xfc, 0x18, 0x74, 0x6e, 0x8b, 0x0e, 0x20, 0x70, 0x83, 0xb8, 0xab, 0x24, 0xbb, 0xf1, 0x7e,
	0x61, 0xe3, 0x1f, 0x40, 0xa0, 0x66, 0xe3, 0x4b, 0x31, 0xca, 0x8b, 0xca, 0xb5, 0xe8, 0x5d, 0xb8,
	0x5f, 0x18, 0xe7, 0x01, 0xd3, 0x68, 0x2a, 0x1d, 0x0b, 0xae, 0x87, 0x39, 0xa3, 0x53, 0x4c, 0xff,
	0xcd, 0x9f, 0xe0, 0x5b, 0xcc, 0x88, 0xd1, 0x2f, 0x1e, 0x54, 0x0b, 0xb1, 0x64, 0x1d, 0xca, 0x66,
	0x1d, 0x6d, 0x58, 0x71, 0x1c, 0x16, 0x35, 0xbf, 0x9c, 0xdb, 0xff, 0x7c, 0x5c, 0x70, 0xb3, 0x49,
	0x34, 0xb7, 0x90, 0x47, 0x50, 0x51, 0x9a, 0x49, 0xd7, 0xb8, 0x6a, 0xf7, 0xb5, 0xbf, 0x2c, 0x9b,
	0xa9, 0x8e, 0x3a, 0x0f, 0xb2, 0x01, 0x25, 0xe4, 0x89, 0x6d, 0xe6, 0xdf, 0x38, 0x1a, 0x7b, 0xf4,
	0xad, 0x07, 0x35, 0xb7, 0x4c, 0x14, 0xbf, 0x98, 0xa0, 0xd2, 0xe6, 0x1c, 0x2c, 0x4b, 0xe7, 0x7b,
	0xc4, 0xb2, 0xd4, 0xb4, 0x6b, 0xc4, 0xf8, 0x60, 0xde, 0x2e, 0x23, 0x1b, 0x6c, 0x86, 0x4c, 0xda,
	0x42, 0x6a, 0xd4, 0xca, 0xe4, 0x1d, 0xc8, 0x7f, 0xa1, 0x8b, 0xd8, 0x5e, 0x6a, 0x79, 0xf2, 0xc6,
	0xcd, 0x6f, 0xe7, 0x2e, 0x3b, 0xba, 0x9a, 0x15, 0xb4, 0x88, 0xc2, 0xaa, 0x6d, 0x03, 0xaa, 0x4c,
	0x70, 0x85, 0xff, 0xb0, 0x80, 0x79, 0x33, 0x4b, 0x77, 0x35, 0x33, 0xfa, 0x1c, 0xea, 0xf3, 0x53,
	0xfd, 0x2b, 0x56, 0x33, 0x04, 0x1b, 0x97, 0xf3, 0xc2, 0x4d, 0xed, 0x34, 0xb7, 0x6c, 0x3d, 0x82,
	0xc0, 0xdd, 0x84, 0x64, 0x05, 0x2a, 0x27, 0xe7, 0x1f, 0xf6, 0x69, 0xe3, 0x1e, 0x59, 0x86, 0xf2,
	0xf3, 0xbd, 0x4f, 0xfa, 0x0d, 0x8f, 0x00, 0x04, 0x87, 0x7d, 0x2b, 0xfb, 0x5b, 0xef, 0xc3, 0xf2,
	0x61, 0x7e, 0xe3, 0x92, 0x35, 0xa8, 0x1d, 0x9f, 0x9c, 0x5f, 0x9c, 0x9d, 0xf6, 0x7b, 0x47, 0x87,
	0x47, 0xfd, 0x83, 0xc6, 0x3d, 0x42, 0xa0, 0xbe, 0xff, 0xe9, 0xd9, 0xd1, 0x71, 0xff, 0xec, 0xec,
	0xe2, 0xe4, 0xb3, 0xe3, 0x3e, 0x6d, 0x78, 0xa4, 0x0a, 0x4b, 0x3d, 0xda, 0xdf, 0x3b, 0x3f, 0xa1,
	0x0d, 0xbf, 0x3b, 0x84, 0xc6, 0xe9, 0xfc, 0x21, 0x3a, 0x43, 0x39, 0x4d, 0x63, 0x24, 0xe7, 0x50,
	0xeb, 0x89, 0x71, 0x36, 0xd1, 0x98, 0xdf, 0xdf, 0x6b, 0x85, 0x1a, 0xdd, 0x0c, 0x5b, 0xa4, 0x08,
	0xb9, 0x06, 0x44, 0xeb, 0xdf, 0xfd, 0xfe, 0xc7, 0x4f, 0xfe, 0x83, 0x68, 0xcd, 0x3e, 0x6f, 0xd6,
	0xd6, 0xf9, 0xca, 0xcc, 0xed, 0xeb, 0x5d, 0x6f, 0x6b, 0xff, 0x37, 0xef, 0xc7, 0xbd, 0x5f, 0x3d,
	0x82, 0xb0, 0xb6, 0x48, 0x18, 0xe6, 0x19, 0xa3, 0x8f, 0xe1, 0xe1, 0xe2, 0x39, 0x0c, 0x95, 0x03,
	0xc3, 0x4c, 0x8a, 0x97, 0x18, 0x6b, 0xd2, 0x1a, 0x6a, 0x9d, 0xa9, 0xdd, 0x4e, 0xe1, 0xc5, 0x94,
	0xa8, 0x50, 0xb7, 0xd5, 0xb0, 0x45, 0x16, 0xd8, 0x07, 0x73, 0xac, 0x5b, 0x7a, 0xd2, 0xde, 0xde,
	0xf2, 0xbc, 0x6e, 0x83, 0x65, 0xd9, 0x28, 0x8d, 0xed, 0xa3, 0xd8, 0x79, 0xa9, 0x04, 0xdf, 0xbd,
	0x85, 0xd0, 0xf7, 0xa0, 0xb4, 0xf3, 0x64, 0x87, 0xec, 0xc0, 0x16, 0x45, 0x3d, 0x91, 0x1c, 0x93,
	0xf0, 0x7a, 0x88, 0x3c, 0xd4, 0x43, 0x0c, 0x25, 0x2a, 0x31, 0x91, 0x31, 0x86, 0x89, 0x40, 0x15,
	0x72, 0xa1, 0x43, 0xfc, 0x32, 0x55, 0xba, 0x4d, 0x02, 0x28, 0xff, 0xec, 0x7b, 0x4b, 0x97, 0x81,
	0x7d, 0x68, 0xdf, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0x51, 0xab, 0x1f, 0x17, 0xd8, 0x07, 0x00,
	0x00,
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
	ComputePerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error)
}

type predictorServiceClient struct {
	cc *grpc.ClientConn
}

func NewPredictorServiceClient(cc *grpc.ClientConn) PredictorServiceClient {
	return &predictorServiceClient{cc}
}

func (c *predictorServiceClient) ComputePerson(ctx context.Context, in *PersonRequest, opts ...grpc.CallOption) (*PersonResponse, error) {
	out := new(PersonResponse)
	err := c.cc.Invoke(ctx, "/v1.PredictorService/ComputePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PredictorServiceServer is the server API for PredictorService service.
type PredictorServiceServer interface {
	ComputePerson(context.Context, *PersonRequest) (*PersonResponse, error)
}

func RegisterPredictorServiceServer(s *grpc.Server, srv PredictorServiceServer) {
	s.RegisterService(&_PredictorService_serviceDesc, srv)
}

func _PredictorService_ComputePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PredictorServiceServer).ComputePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.PredictorService/ComputePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PredictorServiceServer).ComputePerson(ctx, req.(*PersonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PredictorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.PredictorService",
	HandlerType: (*PredictorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputePerson",
			Handler:    _PredictorService_ComputePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/predictor.proto",
}
