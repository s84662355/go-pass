// Code generated by protoc-gen-go. DO NOT EDIT.
// source: convention/protobuf/TcpProtocol结构体.proto

package protocol

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

type ProtocolTag int32

const (
	ProtocolTag_Server ProtocolTag = 0
	ProtocolTag_Client ProtocolTag = 1
)

var ProtocolTag_name = map[int32]string{
	0: "Server",
	1: "Client",
}

var ProtocolTag_value = map[string]int32{
	"Server": 0,
	"Client": 1,
}

func (x ProtocolTag) String() string {
	return proto.EnumName(ProtocolTag_name, int32(x))
}

func (ProtocolTag) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_87e796a72177dad1, []int{0}
}

type TcpProtocol struct {
	Date                 string      `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Tag                  ProtocolTag `protobuf:"varint,2,opt,name=tag,proto3,enum=protocol.ProtocolTag" json:"tag,omitempty"`
	Address              string      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	ECode                string      `protobuf:"bytes,4,opt,name=eCode,proto3" json:"eCode,omitempty"`
	ExtraData            []byte      `protobuf:"bytes,5,opt,name=extraData,proto3" json:"extraData,omitempty"`
	BodyBytes            []byte      `protobuf:"bytes,6,opt,name=bodyBytes,proto3" json:"bodyBytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TcpProtocol) Reset()         { *m = TcpProtocol{} }
func (m *TcpProtocol) String() string { return proto.CompactTextString(m) }
func (*TcpProtocol) ProtoMessage()    {}
func (*TcpProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_87e796a72177dad1, []int{0}
}

func (m *TcpProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TcpProtocol.Unmarshal(m, b)
}
func (m *TcpProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TcpProtocol.Marshal(b, m, deterministic)
}
func (m *TcpProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TcpProtocol.Merge(m, src)
}
func (m *TcpProtocol) XXX_Size() int {
	return xxx_messageInfo_TcpProtocol.Size(m)
}
func (m *TcpProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_TcpProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_TcpProtocol proto.InternalMessageInfo

func (m *TcpProtocol) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *TcpProtocol) GetTag() ProtocolTag {
	if m != nil {
		return m.Tag
	}
	return ProtocolTag_Server
}

func (m *TcpProtocol) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TcpProtocol) GetECode() string {
	if m != nil {
		return m.ECode
	}
	return ""
}

func (m *TcpProtocol) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *TcpProtocol) GetBodyBytes() []byte {
	if m != nil {
		return m.BodyBytes
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.ProtocolTag", ProtocolTag_name, ProtocolTag_value)
	proto.RegisterType((*TcpProtocol)(nil), "protocol.TcpProtocol")
}

func init() {
	proto.RegisterFile("convention/protobuf/TcpProtocol结构体.proto", fileDescriptor_87e796a72177dad1)
}

var fileDescriptor_87e796a72177dad1 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xce, 0xcf, 0x2b,
	0x4b, 0xcd, 0x2b, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3,
	0x0f, 0x49, 0x2e, 0x08, 0x00, 0xb1, 0x93, 0xf3, 0x73, 0x9e, 0xef, 0x9e, 0xfc, 0x6c, 0x5e, 0xcb,
	0x93, 0xbd, 0x93, 0xf5, 0xc0, 0xb2, 0x42, 0x1c, 0x05, 0x50, 0x09, 0xa5, 0x6d, 0x8c, 0x5c, 0xdc,
	0x48, 0x0a, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x52, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0xc0, 0x6c, 0x21, 0x75, 0x2e, 0xe6, 0x92, 0xc4, 0x74, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e,
	0x23, 0x51, 0x3d, 0x98, 0x5e, 0x3d, 0x98, 0xa6, 0x90, 0xc4, 0xf4, 0x20, 0x90, 0x0a, 0x21, 0x09,
	0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x66, 0xb0, 0x7e, 0x18, 0x57, 0x48,
	0x84, 0x8b, 0x35, 0xd5, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x05, 0x2c, 0x0e, 0xe1, 0x08, 0xc9, 0x70,
	0x71, 0xa6, 0x56, 0x94, 0x14, 0x25, 0xba, 0x24, 0x96, 0x24, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0xf0,
	0x04, 0x21, 0x04, 0x40, 0xb2, 0x49, 0xf9, 0x29, 0x95, 0x4e, 0x95, 0x25, 0xa9, 0xc5, 0x12, 0x6c,
	0x10, 0x59, 0xb8, 0x80, 0x96, 0x2a, 0x17, 0x37, 0x92, 0xfd, 0x42, 0x5c, 0x5c, 0x6c, 0xc1, 0xa9,
	0x45, 0x65, 0xa9, 0x45, 0x02, 0x0c, 0x20, 0xb6, 0x73, 0x4e, 0x66, 0x6a, 0x5e, 0x89, 0x00, 0x63,
	0x12, 0x1b, 0xd8, 0xb5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0xb8, 0x0a, 0x8d, 0x22,
	0x01, 0x00, 0x00,
}
