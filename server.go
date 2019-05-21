package gpbrpc

import (
	proto "github.com/gogo/protobuf/proto"
	. "github.com/gotask/gost/stnet"
)

type GpbServer struct {
	*Server
	defaultRPCService *RPCClient
}

func NewGpbServer(name string, loopmsec uint32) (*GpbServer, error) {
	svr := NewServer(name, loopmsec)
	rpcimp := &RPCClient{}
	service, err := svr.AddService("defaultRPCService", "", 0, rpcimp, 0)
	rpcimp.master = service
	return &GpbServer{svr, rpcimp}, err
}
func (svr *GpbServer) Start() error {
	return svr.Server.Start()
}
func (svr *GpbServer) Stop() {
	svr.Server.Stop()
}
func (svr *GpbServer) AddRpcService(name, address string, heartbeat uint32, rpcsevice ServerInterface, threadId int) error {
	rpcImp := &RPCServer{rpcsevice, nil}
	_, err := svr.AddService(name, address, heartbeat, rpcImp, threadId)
	return err
}

func (svr *GpbServer) NewRpcClient(name, address string, handle ClientInterface) {
	c := svr.NewConnect(svr.defaultRPCService.master, name, address, nil)
	c.SetUserData(handle)
	handle.SetRPCClient(svr.defaultRPCService)
	handle.SetConn(c)
}

func (svr *GpbServer) AddGpbService(name, address string, heartbeat uint32, gpbsevice GpbServiceImp, threadId int) (*Service, error) {
	gpbImp := &ServiceGpb{gpbsevice}
	return svr.AddService(name, address, heartbeat, gpbImp, threadId)
}

func (svr *GpbServer) AddGpbClient(name, address string, gpbclient GpbClientImp, threadId int) (*Connect, error) {
	gpbImp := &ConnectGpb{gpbclient}
	return svr.AddConnect(name, address, gpbImp, threadId)
}

func (svr *GpbServer) NewGpbClient(service *Service, name, address string, userdata interface{}) *Connect {
	return svr.NewConnect(service, name, address, userdata)
}

func SendRPCResponse(current *Current, msg interface{}) error {
	m, e := proto.Marshal(msg.(proto.Message))
	if e != nil {
		return e
	}
	req := current.Req
	rsp := &RPCResponsePacket{RPCRetCode: 0, RequestId: req.RequestId, RspPayload: string(m), Context: req.Context}
	buf, _ := proto.Marshal(rsp)
	return current.Sess.AsyncSend(PackSendProtocol(buf))
}
