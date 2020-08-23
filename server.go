package gpbrpc

import (
	. "github.com/gotask/gost/stnet"
)

type GpbServer struct {
	server     *Server
	rpcClients map[string]*RPCClient
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

func (svr *GpbServer) AddGpbService(name, address string, heartbeat uint32, gpbsevice GpbServiceImp, threadId int) (*Service, error) {
	gpbImp := &ServiceGpb{gpbsevice}
	return svr.server.AddService(name, address, heartbeat, gpbImp, threadId)
}

func (svr *GpbServer) AddLoopService(name string, imp LoopService, threadId int) (*Service, error) {
	return svr.server.AddLoopService(name, imp, threadId)
}
