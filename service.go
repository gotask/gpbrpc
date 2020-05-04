package gpbrpc

import (
	. "github.com/gotask/gost/stnet"

	proto "github.com/golang/protobuf/proto"
)

var (
	_MsgTypeId   = make(map[uint32]string, 0)
	_MsgTypeName = make(map[string]uint32, 0)
)

func RegisterProtoMsg(id uint32, msg proto.Message) {
	name := proto.MessageName(msg)
	_MsgTypeId[id] = name
	_MsgTypeName[name] = id
}

type GpbServiceImp interface {
	Init() bool
	Loop()
	Destroy()
	HandleRequest(current *CurrentContent, reqCmdId, reqSeq uint32, req proto.Message) (errcode uint32, rsp proto.Message)
	HandleError(current *CurrentContent, e error)
	HashProcessor(sess *Session, msg *ProtocolRequest) (processorID int)
	SessionOpen(sess *Session)
	SessionClose(sess *Session)
}

type GpbClientImp interface {
	Loop()
	HandleResponse(current *CurrentContent, rspCmdId, rspSeq, pushId uint32, rsp proto.Message)
	HandleError(current *CurrentContent, errCode uint32, e error)
	OnConnected(c *Connector)
	OnDisconnected(c *Connector)
}

type RpcClientImp interface {
	Loop()
	OnConnected(c *Connector)
	OnDisconnected(c *Connector)
}
