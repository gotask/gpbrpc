package gpbrpc

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	proto "github.com/gogo/protobuf/proto"
	. "github.com/gotask/gost/stnet"
)

var (
	ReqRecvdButTimeout    int32 = -1
	ErrRecvdReqTimeOut          = errors.New("received req but timeout.")
	ASyncCallTimeout      int32 = -2
	ErrASyncCallTimeout         = errors.New("async call return timeout.")
	RequestQueueIsFull    int32 = -3
	ErrReqSendQueueIsFull       = errors.New("qequest send queue is full.")
	GPBUnmarshalFailed    int32 = -4
	ErrGPBUnmarshalFailed       = errors.New("gpb proto unmarshal failed.")
	CallRemoteNullFunc    int32 = -5
	ErrCallRemoteNullFunc       = errors.New("call remote null func.")
	SyncCallTimeout       int32 = -6
	ErrSyncCallTimeout          = errors.New("sync call return timeout.")

	ErrRsqSendQueueIsFull = errors.New("response send queue is full.")
	ErrRecvdRspTimeOut    = errors.New("received rsp but timeout.")
)

type RPCRequest struct {
	Req       RPCRequestPacket
	Callback  interface{}
	Exception interface{}
	Signal    chan *RPCResponsePacket
	Handle    ClientInterface
}

type ClientInterface interface {
	HandleRSP(req *RPCRequest, rsp *RPCResponsePacket)
	SetTimeout(t uint32)
	GetTimeout() uint32
	SetRPCClient(*RPCClient)
	SetConn(*Connect)
	GetConn() *Connect
	HashProcessor(*RPCResponsePacket) int
}

type RPCClient struct {
	master      *Service
	requests    []*RPCRequest
	reqMutex    sync.Mutex
	failedReqs  []*RPCRequest
	failedMutex sync.Mutex

	sequence     uint32
	syncRequests sync.Map
}

func (rpc *RPCClient) Send(req *RPCRequest) error {
	buf, _ := proto.Marshal(&req.Req)
	return req.Handle.GetConn().AsyncSend(PackSendProtocol(buf))
}

func (rpc *RPCClient) PushRequest(req *RPCRequest) {
	req.Req.RequestId = atomic.AddUint32(&rpc.sequence, 1)
	req.Req.Timeout = req.Handle.GetTimeout()
	e := rpc.Send(req)
	if e != nil {
		if req.Signal == nil {
			rpc.failedMutex.Lock()
			rpc.failedReqs = append(rpc.failedReqs, req)
			rpc.failedMutex.Unlock()
		} else {
			rsp := &RPCResponsePacket{RPCRetCode: RequestQueueIsFull, RequestId: req.Req.RequestId, Context: req.Req.Context}
			req.Signal <- rsp
			rpc.HandleError(nil, ErrReqSendQueueIsFull)
		}
		return
	}
	req.Req.Timeout += uint32(time.Now().UnixNano() / 1e6)

	if req.Signal == nil {
		if req.Req.IsOneWay {
			return
		}
		rpc.reqMutex.Lock()
		rpc.requests = append(rpc.requests, req)
		rpc.reqMutex.Unlock()
	} else {
		rpc.syncRequests.Store(req.Req.RequestId, req)
	}
}

func (rpc *RPCClient) DeleteRequest(reqId uint32) {
	rpc.syncRequests.Delete(reqId)
}

func (rpc *RPCClient) Init() bool {
	rpc.requests = make([]*RPCRequest, 0, 1024)
	rpc.failedReqs = make([]*RPCRequest, 0, 16)
	rpc.sequence = 1
	return true
}

func (rpc *RPCClient) Loop() {
	rpc.failedMutex.Lock()
	failedTmpQ := rpc.failedReqs
	rpc.failedReqs = make([]*RPCRequest, 0, 16)
	rpc.failedMutex.Unlock()
	for _, v := range failedTmpQ {
		rsp := &RPCResponsePacket{RPCRetCode: RequestQueueIsFull, RequestId: v.Req.RequestId, Context: v.Req.Context}
		v.Handle.HandleRSP(v, rsp)
		rpc.HandleError(nil, ErrReqSendQueueIsFull)
	}

	now := time.Now().UnixNano() / 1e6
	rpc.reqMutex.Lock()
	i := 0
	for ; i < len(rpc.requests); i++ {
		if rpc.requests[i].Req.Timeout > uint32(now) {
			break
		}
	}
	timeoutReq := rpc.requests[:i]
	rpc.requests = rpc.requests[i:]
	rpc.reqMutex.Unlock()

	for _, v := range timeoutReq {
		rsp := &RPCResponsePacket{RPCRetCode: ASyncCallTimeout, RequestId: v.Req.RequestId, Context: v.Req.Context}
		v.Handle.HandleRSP(v, rsp)
		rpc.HandleError(nil, ErrASyncCallTimeout)
	}
}
func (rpc *RPCClient) Destroy() {

}

func (rpc *RPCClient) getRPCRequest(id uint32) (*RPCRequest, int, bool) {
	v := &RPCRequest{}
	ok := false
	i := 0
	for ; i < len(rpc.requests); i++ {
		v = rpc.requests[i]
		if v.Req.RequestId == id {
			ok = true
			break
		} else if v.Req.RequestId > id && v.Req.RequestId-id < 0x7FFFFFFF {
			break
		}
	}
	return v, i, ok
}

func (rpc *RPCClient) HandleMessage(current *CurrentContent, msgID uint32, msg interface{}) {
	rsp := msg.(*RPCResponsePacket)

	rpc.reqMutex.Lock()
	v, i, ok := rpc.getRPCRequest(rsp.RequestId)
	if !ok {
		rpc.reqMutex.Unlock()
		rpc.HandleError(current, ErrRecvdRspTimeOut)
		return
	}
	rpc.requests = append(rpc.requests[:i], rpc.requests[i+1:]...)
	rpc.reqMutex.Unlock()
	v.Handle.HandleRSP(v, rsp)
}

func (rpc *RPCClient) Unmarshal(sess *Session, data []byte) (lenParsed int, msgID int32, msg interface{}, err error) {
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
	rsp := &RPCResponsePacket{}
	e := proto.Unmarshal(data[4:msgLen], rsp)
	if e != nil {
		return int(msgLen), 0, nil, e
	}
	v, ok := rpc.syncRequests.Load(rsp.RequestId)
	if ok { //sync call
		req := v.(*RPCRequest)
		req.Signal <- rsp
		return int(msgLen), -1, nil, nil
	}
	return int(msgLen), 0, rsp, nil
}

func (rpc *RPCClient) HashProcessor(sess *Session, msgID int32, msg interface{}) (processorID int) {
	rsp := msg.(*RPCResponsePacket)
	return sess.UserData.(ClientInterface).HashProcessor(rsp)
}
func (rpc *RPCClient) SessionOpen(sess *Session) {
}
func (rpc *RPCClient) SessionClose(sess *Session) {
}
func (rpc *RPCClient) HeartBeatTimeOut(sess *Session) {
}
func (rpc *RPCClient) HandleError(current *CurrentContent, err error) {
	SysLog.Error(err.Error())
}

type Current struct {
	Sess *Session
	Req  *RPCRequestPacket
}

type ServerInterface interface {
	HandleReq(req *RPCRequestPacket) (msg []byte, ret int32, err error)
	HashProcessor(req *RPCRequestPacket) int
	NewHandle() ServerInterface
	SetCurrent(Current)
	SetResponse(bool)
	IsResponse() bool
}
type RPCServer struct {
	handle     ServerInterface
	realHandle []ServerInterface
}

func (rpc *RPCServer) Init() bool {
	for i := 0; i < ProcessorThreadsNum; i++ {
		rpc.realHandle = append(rpc.realHandle, rpc.handle.NewHandle())
	}
	return true
}
func (rpc *RPCServer) Loop() {
}
func (rpc *RPCServer) Destroy() {

}
func (rpc *RPCServer) HandleMessage(current *CurrentContent, msgID uint32, msg interface{}) {
	req := msg.(*RPCRequestPacket)
	now := time.Now().UnixNano() / 1e6
	if req.Timeout < uint32(now) {
		rpc.SendResponse(current, req, ReqRecvdButTimeout, "")
		rpc.HandleError(current, ErrRecvdReqTimeOut)
		return
	}

	cur := Current{current.Sess, req}
	handle := rpc.realHandle[current.GoroutineID]
	handle.SetCurrent(cur)
	m, r, e := handle.HandleReq(req)
	if !req.IsOneWay || handle.IsResponse() {
		rpc.SendResponse(current, req, r, string(m))
	}
	if e != nil {
		rpc.HandleError(current, e)
	}
	handle.SetResponse(true)
}

func (rpc *RPCServer) SendResponse(current *CurrentContent, req *RPCRequestPacket, ret int32, msg string) error {
	rsp := &RPCResponsePacket{RPCRetCode: ret, RequestId: req.RequestId, RspPayload: msg, Context: req.Context}
	buf, _ := proto.Marshal(rsp)
	e := current.Sess.AsyncSend(PackSendProtocol(buf))
	if e != nil {
		rpc.HandleError(current, ErrRsqSendQueueIsFull)
	}
	return e
}

func (rpc *RPCServer) Unmarshal(sess *Session, data []byte) (lenParsed int, msgID int32, msg interface{}, err error) {
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
	req := &RPCRequestPacket{}
	e := proto.Unmarshal(data[4:msgLen], req)
	if e != nil {
		return int(msgLen), 0, nil, e
	}
	req.Timeout += uint32(time.Now().UnixNano() / 1e6)
	return int(msgLen), 0, req, nil
}
func (rpc *RPCServer) HashProcessor(sess *Session, msgID int32, msg interface{}) (processorID int) {
	req := msg.(*RPCRequestPacket)
	return rpc.handle.HashProcessor(req)
}
func (rpc *RPCServer) SessionOpen(sess *Session) {
}
func (rpc *RPCServer) SessionClose(sess *Session) {
}
func (rpc *RPCServer) HeartBeatTimeOut(sess *Session) {
}
func (rpc *RPCServer) HandleError(current *CurrentContent, err error) {
	SysLog.Error(err.Error())
}

type DefaultRPCService struct {
	current     Current
	notresponse bool
}

//Get msg processor by hash
func (rpc *DefaultRPCService) HashProcessor(req *RPCRequestPacket) int {
	return -1
}

//record current content for sending rpcresponse async
func (rpc *DefaultRPCService) SetCurrent(cur Current) {
	rpc.current = cur
}

//get current content for sending rpcresponse async
func (rpc *DefaultRPCService) GetCurrent() Current {
	return rpc.current
}

//rpc request is need response immediately
func (rpc *DefaultRPCService) SetResponse(b bool) {
	rpc.notresponse = !b
}

//default return should be true
func (rpc *DefaultRPCService) IsResponse() bool {
	return !rpc.notresponse
}