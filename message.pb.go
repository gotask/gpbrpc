// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package gpbrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtocolRequest struct {
	CmdId                uint32   `protobuf:"varint,1,opt,name=CmdId,proto3" json:"CmdId,omitempty"`
	CmdSeq               uint32   `protobuf:"varint,2,opt,name=CmdSeq,proto3" json:"CmdSeq,omitempty"`
	CmdData              string   `protobuf:"bytes,3,opt,name=CmdData,proto3" json:"CmdData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtocolRequest) Reset()         { *m = ProtocolRequest{} }
func (m *ProtocolRequest) String() string { return proto.CompactTextString(m) }
func (*ProtocolRequest) ProtoMessage()    {}
func (*ProtocolRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8b6b6ca31979fe00, []int{0}
}
func (m *ProtocolRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolRequest.Unmarshal(m, b)
}
func (m *ProtocolRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolRequest.Marshal(b, m, deterministic)
}
func (dst *ProtocolRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolRequest.Merge(dst, src)
}
func (m *ProtocolRequest) XXX_Size() int {
	return xxx_messageInfo_ProtocolRequest.Size(m)
}
func (m *ProtocolRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolRequest proto.InternalMessageInfo

func (m *ProtocolRequest) GetCmdId() uint32 {
	if m != nil {
		return m.CmdId
	}
	return 0
}

func (m *ProtocolRequest) GetCmdSeq() uint32 {
	if m != nil {
		return m.CmdSeq
	}
	return 0
}

func (m *ProtocolRequest) GetCmdData() string {
	if m != nil {
		return m.CmdData
	}
	return ""
}

type ProtocolResponse struct {
	CmdId                uint32   `protobuf:"varint,1,opt,name=CmdId,proto3" json:"CmdId,omitempty"`
	CmdSeq               uint32   `protobuf:"varint,2,opt,name=CmdSeq,proto3" json:"CmdSeq,omitempty"`
	PushSeq              uint32   `protobuf:"varint,3,opt,name=PushSeq,proto3" json:"PushSeq,omitempty"`
	ErrorCode            uint32   `protobuf:"varint,4,opt,name=ErrorCode,proto3" json:"ErrorCode,omitempty"`
	CmdData              string   `protobuf:"bytes,5,opt,name=CmdData,proto3" json:"CmdData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtocolResponse) Reset()         { *m = ProtocolResponse{} }
func (m *ProtocolResponse) String() string { return proto.CompactTextString(m) }
func (*ProtocolResponse) ProtoMessage()    {}
func (*ProtocolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8b6b6ca31979fe00, []int{1}
}
func (m *ProtocolResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtocolResponse.Unmarshal(m, b)
}
func (m *ProtocolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtocolResponse.Marshal(b, m, deterministic)
}
func (dst *ProtocolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtocolResponse.Merge(dst, src)
}
func (m *ProtocolResponse) XXX_Size() int {
	return xxx_messageInfo_ProtocolResponse.Size(m)
}
func (m *ProtocolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtocolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProtocolResponse proto.InternalMessageInfo

func (m *ProtocolResponse) GetCmdId() uint32 {
	if m != nil {
		return m.CmdId
	}
	return 0
}

func (m *ProtocolResponse) GetCmdSeq() uint32 {
	if m != nil {
		return m.CmdSeq
	}
	return 0
}

func (m *ProtocolResponse) GetPushSeq() uint32 {
	if m != nil {
		return m.PushSeq
	}
	return 0
}

func (m *ProtocolResponse) GetErrorCode() uint32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *ProtocolResponse) GetCmdData() string {
	if m != nil {
		return m.CmdData
	}
	return ""
}

type RPCRequestPacket struct {
	IsOneWay             bool              `protobuf:"varint,1,opt,name=IsOneWay,proto3" json:"IsOneWay,omitempty"`
	RequestId            uint32            `protobuf:"varint,2,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	ServiceName          string            `protobuf:"bytes,3,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	FuncName             string            `protobuf:"bytes,4,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	ReqPayload           string            `protobuf:"bytes,5,opt,name=ReqPayload,proto3" json:"ReqPayload,omitempty"`
	Timeout              uint32            `protobuf:"varint,6,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Context              map[string]string `protobuf:"bytes,7,rep,name=Context,proto3" json:"Context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RPCRequestPacket) Reset()         { *m = RPCRequestPacket{} }
func (m *RPCRequestPacket) String() string { return proto.CompactTextString(m) }
func (*RPCRequestPacket) ProtoMessage()    {}
func (*RPCRequestPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8b6b6ca31979fe00, []int{2}
}
func (m *RPCRequestPacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCRequestPacket.Unmarshal(m, b)
}
func (m *RPCRequestPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCRequestPacket.Marshal(b, m, deterministic)
}
func (dst *RPCRequestPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCRequestPacket.Merge(dst, src)
}
func (m *RPCRequestPacket) XXX_Size() int {
	return xxx_messageInfo_RPCRequestPacket.Size(m)
}
func (m *RPCRequestPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCRequestPacket.DiscardUnknown(m)
}

var xxx_messageInfo_RPCRequestPacket proto.InternalMessageInfo

func (m *RPCRequestPacket) GetIsOneWay() bool {
	if m != nil {
		return m.IsOneWay
	}
	return false
}

func (m *RPCRequestPacket) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *RPCRequestPacket) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *RPCRequestPacket) GetFuncName() string {
	if m != nil {
		return m.FuncName
	}
	return ""
}

func (m *RPCRequestPacket) GetReqPayload() string {
	if m != nil {
		return m.ReqPayload
	}
	return ""
}

func (m *RPCRequestPacket) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RPCRequestPacket) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

type RPCResponsePacket struct {
	RPCRetCode           int32             `protobuf:"varint,1,opt,name=RPCRetCode,proto3" json:"RPCRetCode,omitempty"`
	RequestId            uint32            `protobuf:"varint,2,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	RspPayload           string            `protobuf:"bytes,3,opt,name=RspPayload,proto3" json:"RspPayload,omitempty"`
	Context              map[string]string `protobuf:"bytes,4,rep,name=Context,proto3" json:"Context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RPCResponsePacket) Reset()         { *m = RPCResponsePacket{} }
func (m *RPCResponsePacket) String() string { return proto.CompactTextString(m) }
func (*RPCResponsePacket) ProtoMessage()    {}
func (*RPCResponsePacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_message_8b6b6ca31979fe00, []int{3}
}
func (m *RPCResponsePacket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCResponsePacket.Unmarshal(m, b)
}
func (m *RPCResponsePacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCResponsePacket.Marshal(b, m, deterministic)
}
func (dst *RPCResponsePacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCResponsePacket.Merge(dst, src)
}
func (m *RPCResponsePacket) XXX_Size() int {
	return xxx_messageInfo_RPCResponsePacket.Size(m)
}
func (m *RPCResponsePacket) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCResponsePacket.DiscardUnknown(m)
}

var xxx_messageInfo_RPCResponsePacket proto.InternalMessageInfo

func (m *RPCResponsePacket) GetRPCRetCode() int32 {
	if m != nil {
		return m.RPCRetCode
	}
	return 0
}

func (m *RPCResponsePacket) GetRequestId() uint32 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *RPCResponsePacket) GetRspPayload() string {
	if m != nil {
		return m.RspPayload
	}
	return ""
}

func (m *RPCResponsePacket) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func init() {
	proto.RegisterType((*ProtocolRequest)(nil), "gpbrpc.ProtocolRequest")
	proto.RegisterType((*ProtocolResponse)(nil), "gpbrpc.ProtocolResponse")
	proto.RegisterType((*RPCRequestPacket)(nil), "gpbrpc.RPCRequestPacket")
	proto.RegisterMapType((map[string]string)(nil), "gpbrpc.RPCRequestPacket.ContextEntry")
	proto.RegisterType((*RPCResponsePacket)(nil), "gpbrpc.RPCResponsePacket")
	proto.RegisterMapType((map[string]string)(nil), "gpbrpc.RPCResponsePacket.ContextEntry")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_message_8b6b6ca31979fe00) }

var fileDescriptor_message_8b6b6ca31979fe00 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6a, 0xdb, 0x40,
	0x10, 0x86, 0x91, 0x65, 0xcb, 0xd6, 0xb8, 0xa6, 0xee, 0x52, 0x8a, 0x30, 0xa5, 0x08, 0x43, 0x8b,
	0x4f, 0x3a, 0xb4, 0x97, 0xe2, 0x4b, 0x0b, 0xaa, 0x0b, 0xbe, 0xb4, 0x62, 0x1d, 0x08, 0x39, 0xae,
	0xa5, 0xc1, 0x31, 0xb6, 0xb4, 0xb2, 0x76, 0x65, 0xe2, 0xd7, 0xc8, 0x1b, 0xe5, 0x89, 0xf2, 0x0a,
	0x61, 0x57, 0x2b, 0x4b, 0xce, 0x25, 0x04, 0x72, 0xd3, 0xff, 0xcf, 0xee, 0xec, 0xff, 0xcd, 0x20,
	0x18, 0xa5, 0x28, 0x04, 0xdb, 0x60, 0x90, 0x17, 0x5c, 0x72, 0xe2, 0x6c, 0xf2, 0x75, 0x91, 0xc7,
	0xd3, 0x1b, 0x78, 0x1f, 0x29, 0x23, 0xe6, 0x7b, 0x8a, 0x87, 0x12, 0x85, 0x24, 0x1f, 0xa1, 0x17,
	0xa6, 0xc9, 0x32, 0xf1, 0x2c, 0xdf, 0x9a, 0x8d, 0x68, 0x25, 0xc8, 0x27, 0x70, 0xc2, 0x34, 0x59,
	0xe1, 0xc1, 0xeb, 0x68, 0xdb, 0x28, 0xe2, 0x41, 0x3f, 0x4c, 0x93, 0x3f, 0x4c, 0x32, 0xcf, 0xf6,
	0xad, 0x99, 0x4b, 0x6b, 0x39, 0xbd, 0xb7, 0x60, 0xdc, 0xf4, 0x16, 0x39, 0xcf, 0x04, 0xbe, 0xbe,
	0x79, 0x54, 0x8a, 0x5b, 0x55, 0xb0, 0x75, 0xa1, 0x96, 0xe4, 0x33, 0xb8, 0x8b, 0xa2, 0xe0, 0x45,
	0xc8, 0x13, 0xf4, 0xba, 0xba, 0xd6, 0x18, 0xed, 0x50, 0xbd, 0xcb, 0x50, 0x0f, 0x1d, 0x18, 0xd3,
	0x28, 0x34, 0xac, 0x11, 0x8b, 0x77, 0x28, 0xc9, 0x04, 0x06, 0x4b, 0xf1, 0x3f, 0xc3, 0x6b, 0x76,
	0xd2, 0xb9, 0x06, 0xf4, 0xac, 0xd5, 0x43, 0xe6, 0xf0, 0x32, 0x31, 0xe9, 0x1a, 0x83, 0xf8, 0x30,
	0x5c, 0x61, 0x71, 0xdc, 0xc6, 0xf8, 0x8f, 0xa5, 0x68, 0x26, 0xd0, 0xb6, 0x54, 0xef, 0xbf, 0x65,
	0x16, 0xeb, 0x72, 0x57, 0x97, 0xcf, 0x9a, 0x7c, 0x01, 0xa0, 0x78, 0x88, 0xd8, 0x69, 0xcf, 0x59,
	0x62, 0x92, 0xb6, 0x1c, 0x85, 0x71, 0xb5, 0x4d, 0x91, 0x97, 0xd2, 0x73, 0x2a, 0x7c, 0x23, 0xc9,
	0x2f, 0xe8, 0x87, 0x3c, 0x93, 0x78, 0x27, 0xbd, 0xbe, 0x6f, 0xcf, 0x86, 0xdf, 0xbf, 0x06, 0xd5,
	0x42, 0x83, 0xe7, 0x70, 0x81, 0x39, 0xb7, 0xc8, 0x64, 0x71, 0xa2, 0xf5, 0xad, 0xc9, 0x1c, 0xde,
	0xb5, 0x0b, 0x64, 0x0c, 0xf6, 0x0e, 0x2b, 0x7a, 0x97, 0xaa, 0x4f, 0xb5, 0xa9, 0x23, 0xdb, 0x97,
	0xa8, 0xa1, 0x5d, 0x5a, 0x89, 0x79, 0xe7, 0xa7, 0x35, 0x7d, 0xb4, 0xe0, 0x83, 0x7e, 0xa6, 0xda,
	0xa9, 0x19, 0xa2, 0x82, 0x51, 0xa6, 0xd4, 0x2b, 0x51, 0x8d, 0x7a, 0xb4, 0xe5, 0xbc, 0x30, 0x48,
	0x75, 0x5b, 0xe4, 0xf5, 0x28, 0x6c, 0x33, 0x8a, 0xb3, 0x43, 0x7e, 0x37, 0xc0, 0x5d, 0x0d, 0xfc,
	0xed, 0x02, 0xb8, 0x9d, 0xe4, 0xed, 0x89, 0xd7, 0x8e, 0xfe, 0x69, 0x7e, 0x3c, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x62, 0x66, 0x9f, 0x83, 0x45, 0x03, 0x00, 0x00,
}
