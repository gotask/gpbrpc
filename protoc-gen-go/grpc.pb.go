// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	"time"
	"github.com/gotask/gpbrpc"
	"github.com/gotask/gost/stnet"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SimpleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ed5933cbeff6746b, []int{0}
}
func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (dst *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(dst, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

type SimpleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ed5933cbeff6746b, []int{1}
}
func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (dst *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(dst, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

type StreamMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMsg) Reset()         { *m = StreamMsg{} }
func (m *StreamMsg) String() string { return proto.CompactTextString(m) }
func (*StreamMsg) ProtoMessage()    {}
func (*StreamMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ed5933cbeff6746b, []int{2}
}
func (m *StreamMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg.Unmarshal(m, b)
}
func (m *StreamMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg.Marshal(b, m, deterministic)
}
func (dst *StreamMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg.Merge(dst, src)
}
func (m *StreamMsg) XXX_Size() int {
	return xxx_messageInfo_StreamMsg.Size(m)
}
func (m *StreamMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg proto.InternalMessageInfo

type StreamMsg2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamMsg2) Reset()         { *m = StreamMsg2{} }
func (m *StreamMsg2) String() string { return proto.CompactTextString(m) }
func (*StreamMsg2) ProtoMessage()    {}
func (*StreamMsg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_grpc_ed5933cbeff6746b, []int{3}
}
func (m *StreamMsg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg2.Unmarshal(m, b)
}
func (m *StreamMsg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg2.Marshal(b, m, deterministic)
}
func (dst *StreamMsg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg2.Merge(dst, src)
}
func (m *StreamMsg2) XXX_Size() int {
	return xxx_messageInfo_StreamMsg2.Size(m)
}
func (m *StreamMsg2) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg2.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "main.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "main.SimpleResponse")
	proto.RegisterType((*StreamMsg)(nil), "main.StreamMsg")
	proto.RegisterType((*StreamMsg2)(nil), "main.StreamMsg2")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ gpbrpc.RPCClient
var _ stnet.Connect

// TestClient is the client API for Test service.
type Test_UnaryCall_Exception = func(int32)
type Test_UnaryCall_Callback = func(*SimpleResponse)
type Test_Downstream_Exception = func(int32)
type Test_Downstream_Callback = func(*StreamMsg)
type Test_Upstream_Exception = func(int32)
type Test_Upstream_Callback = func(*SimpleResponse)
type Test_Bidi_Exception = func(int32)
type Test_Bidi_Callback = func(*StreamMsg2)

type TestClient struct {
	conn      *stnet.Connect
	rpcclient *gpbrpc.RPCClient
	content   map[string]string
	timeout   uint32
	hashfunc  func(*gpbrpc.RPCResponsePacket) int
}

// new client for Test service.
func NewTestClient() *TestClient {
	return &TestClient{nil, nil, make(map[string]string), 5000, nil}
}

// Change the content of this client which is send to service.
func (c *TestClient) SetContent(k, v string) {
	c.content[k] = v
}

func (c *TestClient) DelContent(k string) {
	delete(c.content, k)
}

func (c *TestClient) GetContent() map[string]string {
	return c.content
}

// set timeout of the response from service.
func (c *TestClient) SetTimeout(t uint32) {
	c.timeout = t
}

func (c *TestClient) GetTimeout() uint32 {
	return c.timeout
}

// set rpcclient.
func (c *TestClient) SetRPCClient(rpc *gpbrpc.RPCClient) {
	c.rpcclient = rpc
}

// set connection which is connecting the service.
func (c *TestClient) SetConn(conn *stnet.Connect) {
	c.conn = conn
}

func (c *TestClient) GetConn() *stnet.Connect {
	return c.conn
}

// decide which thread is used to handle the response.
func (c *TestClient) SetHashFunc(f func(*gpbrpc.RPCResponsePacket) int) {
	c.hashfunc = f
}

func (c *TestClient) HashProcessor(rsp *gpbrpc.RPCResponsePacket) int {
	if c.hashfunc != nil {
		return c.hashfunc(rsp)
	}
	return -1
}

// async call functions
func (_c *TestClient) UnaryCall(in *SimpleRequest, cb Test_UnaryCall_Callback, exp Test_UnaryCall_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "UnaryCall", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.rpcclient.PushRequest(req)
}

func (_c *TestClient) Downstream(in *SimpleRequest, cb Test_Downstream_Callback, exp Test_Downstream_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "Downstream", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.rpcclient.PushRequest(req)
}

func (_c *TestClient) Upstream(in *StreamMsg, cb Test_Upstream_Callback, exp Test_Upstream_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "Upstream", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.rpcclient.PushRequest(req)
}

func (_c *TestClient) Bidi(in *StreamMsg, cb Test_Bidi_Callback, exp Test_Bidi_Exception) {
	buf, _ := proto.Marshal(in)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "Bidi", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}
	if cb == nil && exp == nil {
		req.Req.IsOneWay = true
	}
	_c.rpcclient.PushRequest(req)
}

// sync call functions
func (_c *TestClient) UnaryCall_Sync(in *SimpleRequest) (out *SimpleResponse, ret int32) {
	out = &SimpleResponse{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "UnaryCall", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.rpcclient.PushRequest(req)
	to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)
	select {
	case s := <-sg:
		ret = s.RPCRetCode
		if len(s.RspPayload) > 0 {
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				ret = gpbrpc.GPBUnmarshalFailed
			}
		}
	case <-to.C:
		ret = gpbrpc.SyncCallTimeout
	}
	to.Stop()
	_c.rpcclient.DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *TestClient) Downstream_Sync(in *SimpleRequest) (out *StreamMsg, ret int32) {
	out = &StreamMsg{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "Downstream", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.rpcclient.PushRequest(req)
	to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)
	select {
	case s := <-sg:
		ret = s.RPCRetCode
		if len(s.RspPayload) > 0 {
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				ret = gpbrpc.GPBUnmarshalFailed
			}
		}
	case <-to.C:
		ret = gpbrpc.SyncCallTimeout
	}
	to.Stop()
	_c.rpcclient.DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *TestClient) Upstream_Sync(in *StreamMsg) (out *SimpleResponse, ret int32) {
	out = &SimpleResponse{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "Upstream", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.rpcclient.PushRequest(req)
	to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)
	select {
	case s := <-sg:
		ret = s.RPCRetCode
		if len(s.RspPayload) > 0 {
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				ret = gpbrpc.GPBUnmarshalFailed
			}
		}
	case <-to.C:
		ret = gpbrpc.SyncCallTimeout
	}
	to.Stop()
	_c.rpcclient.DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *TestClient) Bidi_Sync(in *StreamMsg) (out *StreamMsg2, ret int32) {
	out = &StreamMsg2{}
	buf, _ := proto.Marshal(in)
	sg := make(chan *gpbrpc.RPCResponsePacket, 1)
	req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "Test", FuncName: "Bidi", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}
	_c.rpcclient.PushRequest(req)
	to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)
	select {
	case s := <-sg:
		ret = s.RPCRetCode
		if len(s.RspPayload) > 0 {
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				ret = gpbrpc.GPBUnmarshalFailed
			}
		}
	case <-to.C:
		ret = gpbrpc.SyncCallTimeout
	}
	to.Stop()
	_c.rpcclient.DeleteRequest(req.Req.RequestId)
	return out, ret
}

func (_c *TestClient) HandleRSP(r *gpbrpc.RPCRequest, s *gpbrpc.RPCResponsePacket) {
	if r.Req.FuncName == "UnaryCall" {
		if s.RPCRetCode != 0 {
			if r.Exception.(Test_UnaryCall_Exception) != nil {
				r.Exception.(Test_UnaryCall_Exception)(s.RPCRetCode)
			}
		} else {
			out := &SimpleResponse{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(Test_UnaryCall_Exception) != nil {
					r.Exception.(Test_UnaryCall_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(Test_UnaryCall_Callback) != nil {
					r.Callback.(Test_UnaryCall_Callback)(out)
				}
			}
		}
	} else if r.Req.FuncName == "Downstream" {
		if s.RPCRetCode != 0 {
			if r.Exception.(Test_Downstream_Exception) != nil {
				r.Exception.(Test_Downstream_Exception)(s.RPCRetCode)
			}
		} else {
			out := &StreamMsg{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(Test_Downstream_Exception) != nil {
					r.Exception.(Test_Downstream_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(Test_Downstream_Callback) != nil {
					r.Callback.(Test_Downstream_Callback)(out)
				}
			}
		}
	} else if r.Req.FuncName == "Upstream" {
		if s.RPCRetCode != 0 {
			if r.Exception.(Test_Upstream_Exception) != nil {
				r.Exception.(Test_Upstream_Exception)(s.RPCRetCode)
			}
		} else {
			out := &SimpleResponse{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(Test_Upstream_Exception) != nil {
					r.Exception.(Test_Upstream_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(Test_Upstream_Callback) != nil {
					r.Callback.(Test_Upstream_Callback)(out)
				}
			}
		}
	} else if r.Req.FuncName == "Bidi" {
		if s.RPCRetCode != 0 {
			if r.Exception.(Test_Bidi_Exception) != nil {
				r.Exception.(Test_Bidi_Exception)(s.RPCRetCode)
			}
		} else {
			out := &StreamMsg2{}
			if err := proto.Unmarshal(s.RspPayload, out); err != nil {
				if r.Exception.(Test_Bidi_Exception) != nil {
					r.Exception.(Test_Bidi_Exception)(gpbrpc.GPBUnmarshalFailed)
				}
			} else {
				if r.Callback.(Test_Bidi_Callback) != nil {
					r.Callback.(Test_Bidi_Callback)(out)
				}
			}
		}
	}
}

// TestServer is the server API for Test service.
type TestServer interface {
	UnaryCall(*SimpleRequest) *SimpleResponse
	// This RPC streams from the server only.
	Downstream(*SimpleRequest) *StreamMsg
	// This RPC streams from the client.
	Upstream(*StreamMsg) *SimpleResponse
	// This one streams in both directions.
	Bidi(*StreamMsg) *StreamMsg2
	//Get msg processor by hash
	HashProcessor(req *gpbrpc.RPCRequestPacket) int
	//each thread has one handle
	NewHandle(*gpbrpc.RPCHelper) TestServer
}

type Test struct {
	*gpbrpc.RPCHelper
	TestServer
}

func NewTestServer(s TestServer) *Test {
	return &Test{&gpbrpc.RPCHelper{}, s}
}

func (s *Test) NewHandle() gpbrpc.ServerInterface {
	h := &gpbrpc.RPCHelper{}
	return &Test{h, s.TestServer.NewHandle(h)}
}

func (s *Test) HandleReq(req *gpbrpc.RPCRequestPacket) (msg []byte, ret int32, err error) {
	if req.FuncName == "UnaryCall" {
		in := &SimpleRequest{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.UnaryCall(in))
		}
	} else if req.FuncName == "Downstream" {
		in := &SimpleRequest{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.Downstream(in))
		}
	} else if req.FuncName == "Upstream" {
		in := &StreamMsg{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.Upstream(in))
		}
	} else if req.FuncName == "Bidi" {
		in := &StreamMsg{}
		if err = proto.Unmarshal(req.ReqPayload, in); err != nil {
			ret = gpbrpc.GPBUnmarshalFailed
			err = gpbrpc.ErrGPBUnmarshalFailed
		} else {
			msg, err = proto.Marshal(s.Bidi(in))
		}
	} else {
		ret = gpbrpc.CallRemoteNullFunc
		err = gpbrpc.ErrCallRemoteNullFunc
	}
	return msg, ret, err
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_grpc_ed5933cbeff6746b) }

var fileDescriptor_grpc_ed5933cbeff6746b = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0xe2, 0xe7, 0xe2,
	0x0d, 0xce, 0xcc, 0x2d, 0xc8, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x12, 0xe0,
	0xe2, 0x83, 0x09, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x71, 0x73, 0x71, 0x06, 0x97, 0x14,
	0xa5, 0x26, 0xe6, 0xfa, 0x16, 0xa7, 0x2b, 0xf1, 0x70, 0x71, 0xc1, 0x39, 0x46, 0x46, 0x37, 0x18,
	0xb9, 0x58, 0x42, 0x52, 0x8b, 0x4b, 0x84, 0xcc, 0xb8, 0x38, 0x43, 0xf3, 0x12, 0x8b, 0x2a, 0x9d,
	0x13, 0x73, 0x72, 0x84, 0x84, 0xf5, 0x40, 0x46, 0xeb, 0xa1, 0x98, 0x2b, 0x25, 0x82, 0x2a, 0x08,
	0x31, 0x5b, 0xc8, 0x84, 0x8b, 0xcb, 0x25, 0xbf, 0x3c, 0xaf, 0x18, 0x6c, 0x24, 0x76, 0x8d, 0xfc,
	0x50, 0x41, 0x98, 0xad, 0x06, 0x8c, 0x42, 0xc6, 0x5c, 0x1c, 0xa1, 0x05, 0x50, 0x3d, 0xe8, 0xd2,
	0xd8, 0x2d, 0xd2, 0x60, 0x14, 0xd2, 0xe5, 0x62, 0x71, 0xca, 0x4c, 0xc9, 0xc4, 0xd4, 0x20, 0x80,
	0x26, 0x60, 0xa4, 0xc1, 0x68, 0xc0, 0x98, 0xc4, 0x06, 0x0e, 0x25, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe9, 0x64, 0xa2, 0xdc, 0x33, 0x01, 0x00, 0x00,
}
