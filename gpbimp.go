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
func (service *ServiceGpb) HandleMessage(current *CurrentContent, msgID uint64, m interface{}) {
	msg := m.(*ProtocolMessage)
	if s, ok := _MsgTypeId[msg.CmdId]; ok {
		t := proto.MessageType(s).Elem()
		req := reflect.New(t).Interface()
		e := proto.Unmarshal(msg.CmdData, req.(proto.Message))
		if e != nil {
			service.imp.HandleError(current, e)
			return
		}
		seq, rsp := service.imp.Handle(current, msg.CmdId, msg.CmdSeq, req.(proto.Message))
		r := ProtocolMessage{}
		if rsp != nil {
			r.CmdSeq = seq
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
		e = current.Sess.AsyncSend(PackSendProtocol(buf))
		if e != nil {
			service.imp.HandleError(current, e)
			return
		}
	} else {
		service.imp.HandleError(current, fmt.Errorf("cannot find msg type by id: %d", msg.CmdId))
	}
}
func (service *ServiceGpb) Unmarshal(sess *Session, data []byte) (lenParsed int, msgID int64, msg interface{}, err error) {
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
	m := &ProtocolMessage{}
	e := proto.Unmarshal(data[4:msgLen], m)
	if e != nil {
		return int(msgLen), 0, nil, e
	}
	return int(msgLen), 0, m, nil
}
func (service *ServiceGpb) HashProcessor(sess *Session, msgID uint64, msg interface{}) (processorID int) {
	m := msg.(*ProtocolMessage)
	return service.imp.HashProcessor(sess, m)
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
