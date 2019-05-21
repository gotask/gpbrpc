package gpbrpc

import (
	. "github.com/gotask/gost/stnet"
)

type GpbServiceImp interface {
	Init() bool
	Loop()
	Destroy()
	HandleMessage(current *CurrentContent, msg *ProtocolRequest)
	HashProcessor(sess *Session, msg *ProtocolRequest) (processorID int)
	SessionOpen(sess *Session)
	SessionClose(sess *Session)
}

type GpbClientImp interface {
	Loop()
	HandleMessage(current *CurrentContent, msg *ProtocolResponse)
	HashProcessor(sess *Session, msg *ProtocolResponse) (processorID int)
	OnConnected()
	OnDisconnected()
}
