package gpbrpc

import (
	proto "github.com/gogo/protobuf/proto"
	. "github.com/gotask/gost/stnet"
)

var (
	MsgHeaderLen        = 4
	MinMsgLen    uint32 = 4
	MaxMsgLen    uint32 = 1024 * 1024 * 100
)

type ServiceGpb struct {
	imp GpbServiceImp
}

func (service *ServiceGpb) Init() bool {
	return service.imp.Init()
}
func (service *ServiceGpb) Loop() {
	service.imp.Loop()
}
func (service *ServiceGpb) Destroy() {
	service.imp.Destroy()

}
func (service *ServiceGpb) HandleMessage(current *CurrentContent, msgID uint32, msg interface{}) {
	req := msg.(*ProtocolRequest)
	service.imp.HandleMessage(current, req)
}
func (service *ServiceGpb) Unmarshal(sess *Session, data []byte) (lenParsed int, msgID int32, msg interface{}, err error) {
	if len(data) < 4 {
		return 0, 0, nil, nil
	}
	msgLen := DecodeLength(data)
	if msgLen < MinMsgLen || msgLen > MaxMsgLen {
		return int(msgLen), 0, nil, ErrInvalidMsgLen
	}
	if len(data) < int(msgLen) {
		return 0, 0, nil, nil
	}
	req := &ProtocolRequest{}
	e := proto.Unmarshal(data[4:msgLen], req)
	if e != nil {
		return int(msgLen), 0, nil, e
	}
	return int(msgLen), 0, req, nil
}
func (service *ServiceGpb) HashProcessor(sess *Session, msgID int32, msg interface{}) (processorID int) {
	req := msg.(*ProtocolRequest)
	return service.imp.HashProcessor(sess, req)
}
func (service *ServiceGpb) SessionOpen(sess *Session) {
	service.imp.SessionOpen(sess)
}
func (service *ServiceGpb) SessionClose(sess *Session) {
	service.imp.SessionClose(sess)
}
func (service *ServiceGpb) HeartBeatTimeOut(sess *Session) {
	sess.Close()
}
func (service *ServiceGpb) HandleError(current *CurrentContent, err error) {
	SysLog.Error(err.Error())
	current.Sess.Close()
}

type ConnectGpb struct {
	imp GpbClientImp
}

func (cs *ConnectGpb) Init() bool {
	return true
}
func (cs *ConnectGpb) Loop() {
	cs.imp.Loop()
}
func (cs *ConnectGpb) Destroy() {

}
func (cs *ConnectGpb) HandleMessage(current *CurrentContent, msgID uint32, msg interface{}) {
	rsp := msg.(*ProtocolResponse)
	cs.imp.HandleMessage(current, rsp)
}
func (cs *ConnectGpb) Unmarshal(sess *Session, data []byte) (lenParsed int, msgID int32, msg interface{}, err error) {
	if len(data) < 4 {
		return 0, 0, nil, nil
	}
	msgLen := DecodeLength(data)
	if msgLen < MinMsgLen || msgLen > MaxMsgLen {
		return int(msgLen), 0, nil, ErrInvalidMsgLen
	}
	if len(data) < int(msgLen) {
		return 0, 0, nil, nil
	}

	rsp := &ProtocolResponse{}
	e := proto.Unmarshal(data[4:msgLen], rsp)
	if e != nil {
		return int(msgLen), 0, nil, e
	}

	return int(msgLen), 0, rsp, nil
}
func (cs *ConnectGpb) HashProcessor(sess *Session, msgID int32, msg interface{}) (processorID int) {
	rsp := msg.(*ProtocolResponse)
	return cs.imp.HashProcessor(sess, rsp)
}
func (cs *ConnectGpb) SessionOpen(sess *Session) {
	cs.imp.OnConnected()
}
func (cs *ConnectGpb) SessionClose(sess *Session) {
	cs.imp.OnDisconnected()
}
func (cs *ConnectGpb) HeartBeatTimeOut(sess *Session) {

}
func (cs *ConnectGpb) HandleError(current *CurrentContent, err error) {
	SysLog.Error(err.Error())
	current.Sess.Close()
}
