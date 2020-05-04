package gpbrpc

import (
	"fmt"
	"reflect"

	. "github.com/gotask/gost/stnet"

	proto "github.com/golang/protobuf/proto"
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
func (service *ServiceGpb) HandleMessage(current *CurrentContent, msgID uint32, m interface{}) {
	msg := m.(*ProtocolRequest)
	if s, ok := _MsgTypeId[msg.CmdId]; ok {
		t := proto.MessageType(s).Elem()
		req := reflect.New(t).Interface()
		e := proto.Unmarshal(msg.CmdData, req.(proto.Message))
		if e != nil {
			service.imp.HandleError(current, e)
			return
		}
		err, rsp := service.imp.HandleRequest(current, msg.CmdId, msg.CmdSeq, req.(proto.Message))
		r := ProtocolResponse{}
		if err != 0 {
			r.ErrorCode = err
		} else if rsp != nil {
			r.CmdSeq = msg.CmdSeq
			id, ok := _MsgTypeName[proto.MessageName(rsp)]
			if !ok {
				service.imp.HandleError(current, fmt.Errorf("cannot find rsp msg type by name: %s", proto.MessageName(rsp)))
				return
			}
			r.CmdId = id
			r.CmdData, e = proto.Marshal(rsp)
			if e != nil {
				service.imp.HandleError(current, e)
				return
			}
		} else {
			return
		}
		buf, e := proto.Marshal(&r)
		if e != nil {
			service.imp.HandleError(current, e)
			return
		}
		e = current.Sess.Send(PackSendProtocol(buf))
		if e != nil {
			service.imp.HandleError(current, e)
			return
		}
	} else {
		service.imp.HandleError(current, fmt.Errorf("cannot find req msg type by id: %d", msg.CmdId))
	}
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
func (cs *ConnectGpb) HandleMessage(current *CurrentContent, msgID uint32, m interface{}) {
	msg := m.(*ProtocolResponse)

	if msg.ErrorCode != 0 {
		cs.imp.HandleError(current, msg.ErrorCode, nil)
		return
	}
	if s, ok := _MsgTypeId[msg.CmdId]; ok {
		t := proto.MessageType(s).Elem()
		rsp := reflect.New(t).Interface()
		e := proto.Unmarshal(msg.CmdData, rsp.(proto.Message))
		if e != nil {
			cs.imp.HandleError(current, 0, e)
			return
		}
		cs.imp.HandleResponse(current, msg.CmdId, msg.CmdSeq, msg.PushSeq, rsp.(proto.Message))
	} else {
		cs.imp.HandleError(current, 0, fmt.Errorf("cannot find rsp msg type by id: %d", msg.CmdId))
	}
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
	return -1
}
func (cs *ConnectGpb) SessionOpen(sess *Session) {
	cs.imp.OnConnected(sess.Connector())
}
func (cs *ConnectGpb) SessionClose(sess *Session) {
	cs.imp.OnDisconnected(sess.Connector())
}
func (cs *ConnectGpb) HeartBeatTimeOut(sess *Session) {

}
func (cs *ConnectGpb) HandleError(current *CurrentContent, err error) {
	SysLog.Error(err.Error())
	current.Sess.Close()
}
