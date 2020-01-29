package gpbrpc

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"
	. "github.com/gotask/gost/stnet"
)

type GpbServer struct {
	server  *Server
	clients map[string]*RPCClient
}

func NewGpbServer(name string, loopmsec uint32) *GpbServer {
	return &GpbServer{NewServer(name, loopmsec), make(map[string]*RPCClient)}
}
func (svr *GpbServer) Start() error {
	return svr.server.Start()
}
func (svr *GpbServer) Stop() {
	svr.server.Stop()
}
func (svr *GpbServer) AddRpcService(name, address string, heartbeat uint32, rpcsevice ServerInterface, threadId int) error {
	rpcImp := &RPCServer{rpcsevice, nil}
	_, err := svr.server.AddService(name, address, heartbeat, rpcImp, threadId)
	return err
}

func (svr *GpbServer) AddRpcClient(name string, threadId int) error {
	rpcimp := &RPCClient{}
	service, err := svr.server.AddService(name, "", 0, rpcimp, threadId)
	if err != nil {
		return err
	}
	rpcimp.master = service
	svr.clients[name] = rpcimp
	return nil
}

func (svr *GpbServer) NewRpcConnector(clientName, address string, handle ClientInterface) error {
	client, ok := svr.clients[clientName]
	if !ok {
		return fmt.Errorf("no this rpcclient: %s", clientName)
	}
	c := client.master.NewConnect(clientName, address, nil)
	c.SetUserData(handle)
	handle.SetRPCClient(client)
	handle.SetConn(c)
	return nil
}

func (svr *GpbServer) AddGpbService(name, address string, heartbeat uint32, gpbsevice GpbServiceImp, threadId int) (*Service, error) {
	gpbImp := &ServiceGpb{gpbsevice}
	return svr.server.AddService(name, address, heartbeat, gpbImp, threadId)
}

func (svr *GpbServer) AddGpbClient(name string, gpbclient GpbClientImp, threadId int) (*Service, error) {
	gpbImp := &ConnectGpb{gpbclient}
	return svr.server.AddService(name, "", 0, gpbImp, threadId)
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
