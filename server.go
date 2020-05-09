package gpbrpc

import (
	. "github.com/gotask/gost/stnet"

	proto "github.com/golang/protobuf/proto"
)

type GpbServer struct {
	server      *Server
	rpcClients  map[string]*RPCClient
	gpbServices map[string]*Service
}

func NewGpbServer(name string, loopmsec uint32) *GpbServer {
	return &GpbServer{NewServer(name, loopmsec), make(map[string]*RPCClient), make(map[string]*Service)}
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

func (svr *GpbServer) AddRpcClient(clientid string, imp RpcClientImp, threadId int) error {
	rpcimp := &RPCClient{clientImp: imp}
	service, err := svr.server.AddService(clientid, "", 0, rpcimp, threadId)
	if err != nil {
		return err
	}
	rpcimp.master = service
	svr.rpcClients[clientid] = rpcimp
	return nil
}

func (svr *GpbServer) NewRpcConnector(clientid, address string, handle ClientInterface) *Connect {
	client, ok := svr.rpcClients[clientid]
	if !ok {
		panic("no rpcclient: " + clientid)
	}
	c := client.master.NewConnect(address, nil)
	handle.SetRPCClient(client)
	handle.SetConn(c)
	return c
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

func (svr *GpbServer) AddGpbService(name, address string, heartbeat uint32, gpbsevice GpbServiceImp, threadId int) (*Service, error) {
	gpbImp := &ServiceGpb{gpbsevice}
	return svr.server.AddService(name, address, heartbeat, gpbImp, threadId)
}

func (svr *GpbServer) AddLoopService(name string, imp LoopService, threadId int) (*Service, error) {
	return svr.server.AddLoopService(name, imp, threadId)
}

func (svr *GpbServer) AddGpbClient(clientid string, gpbclient GpbClientImp, threadId int) error {
	gpbImp := &ConnectGpb{gpbclient}
	s, e := svr.server.AddService(clientid, "", 0, gpbImp, threadId)
	if e != nil {
		return e
	}
	svr.gpbServices[clientid] = s
	return nil
}

func (svr *GpbServer) NewGpbConnector(clientid, address string, userdata interface{}) *Connect {
	client, ok := svr.gpbServices[clientid]
	if !ok {
		panic("no clientservice: " + clientid)
	}
	return client.NewConnect(address, userdata)
}
