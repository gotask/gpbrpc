package gpbrpc

import (
	"fmt"

	. "github.com/gotask/gost/stnet"

	proto "github.com/golang/protobuf/proto"
)

var (
	_MsgTypeId   = make(map[uint64]string, 0)
	_MsgTypeName = make(map[string]uint64, 0)
)

func RegisterProtoMsg(id uint64, msg proto.Message) {
	name := proto.MessageName(msg)
	_MsgTypeId[id] = name
	_MsgTypeName[name] = id
}

func SendGpbMsg(sess *Session, cmdseq uint64, msg proto.Message) error {
	r := ProtocolMessage{}
	id, ok := _MsgTypeName[proto.MessageName(msg)]
	if !ok {
		return fmt.Errorf("cannot find msg type by name: %s", proto.MessageName(msg))
	}
	r.CmdId = id
	r.CmdSeq = cmdseq
	if msg != nil {
		m, e := proto.Marshal(msg.(proto.Message))
		if e != nil {
			return e
		}
		r.CmdData = m
	}

	buf, e := proto.Marshal(&r)
	if e != nil {
		return e
	}
	return sess.AsyncSend(PackSendProtocol(buf))
}

func SendRawGpb(sess *Session, cmdid, cmdseq uint64, msg []byte) error {
	r := ProtocolMessage{}
	r.CmdId = cmdid
	r.CmdSeq = cmdseq
	r.CmdData = msg

	buf, e := proto.Marshal(&r)
	if e != nil {
		return e
	}
	return sess.AsyncSend(PackSendProtocol(buf))
}

func SendRPCResponse(current *Current, msg interface{}) error {
	m, e := proto.Marshal(msg.(proto.Message))
	if e != nil {
		return e
	}
	req := current.Req
	rsp := &RPCResponsePacket{RPCRetCode: 0, RequestId: req.RequestId, RspPayload: m, Context: req.Context}
	buf, _ := proto.Marshal(rsp)
	return current.Sess.AsyncSend(PackSendProtocol(buf))
}

type GpbServiceImp interface {
	Init() bool
	Loop()
	Destroy()
	Handle(current *CurrentContent, cmdId, cmdSeq uint64, req proto.Message) (seq uint64, rsp proto.Message)
	HandleError(current *CurrentContent, e error)
	HashProcessor(sess *Session, msg *ProtocolMessage) (processorID int)
	SessionOpen(sess *Session)
	SessionClose(sess *Session)
}

type RpcClientImp interface {
	Loop()
	OnConnected(c *Connector)
	OnDisconnected(c *Connector)
}
