package main

import (
	"fmt"
	"strings"

	pb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init() {
	generator.RegisterPlugin(new(grpc))
}

// grpc is an implementation of the Go protocol buffer compiler's
// plugin architecture.  It generates bindings for gRPC support.
type grpc struct {
	gen *generator.Generator
}

// Name returns the name of this plugin, "grpc".
func (g *grpc) Name() string {
	return "gpbrpc"
}

// The names for packages imported in the generated code.
// They may vary from the final path component of the import path
// if the name is used by other packages.
var (
	gostPkg string
	grpcPkg string
)

// Init initializes the plugin.
func (g *grpc) Init(gen *generator.Generator) {
	g.gen = gen
}

// Given a type name defined in a .proto, return its object.
// Also record that we're using it, to guarantee the associated import.
func (g *grpc) objectNamed(name string) generator.Object {
	g.gen.RecordTypeUse(name)
	return g.gen.ObjectNamed(name)
}

// Given a type name defined in a .proto, return its name as we will print it.
func (g *grpc) typeName(str string) string {
	return g.gen.TypeName(g.objectNamed(str))
}

// P forwards to g.gen.P.
func (g *grpc) P(args ...interface{}) { g.gen.P(args...) }

// Generate generates code for the services in the given file.
func (g *grpc) Generate(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}

	g.P("// Reference imports to suppress errors if they are not otherwise used.")
	g.P("var _ gpbrpc.RPCClient")
	g.P("var _ stnet.Connect")
	g.P()

	for i, service := range file.FileDescriptorProto.Service {
		g.generateService(file, service, i)
	}
}

// GenerateImports generates the import declaration for this file.
func (g *grpc) GenerateImports(file *generator.FileDescriptor) {
	if len(file.FileDescriptorProto.Service) == 0 {
		return
	}
	g.P("import (")
	g.P(`"time"`)
	g.P(`"github.com/gotask/gpbrpc"`)
	g.P(`"github.com/gotask/gost/stnet"`)
	g.P(")")
	g.P()
}

// reservedClientName records whether a client name is reserved on the client side.
var reservedClientName = map[string]bool{
	// TODO: do we need any in gRPC?
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

// deprecationComment is the standard comment added to deprecated
// messages, fields, enums, and enum values.
var deprecationComment = "// Deprecated: Do not use."

// generateService generates all the code for the named service.
func (g *grpc) generateService(file *generator.FileDescriptor, service *pb.ServiceDescriptorProto, index int) {
	path := fmt.Sprintf("6,%d", index) // 6 means service.

	origServName := service.GetName()
	fullServName := origServName
	if pkg := file.GetPackage(); pkg != "" {
		fullServName = pkg + "." + fullServName
	}
	servName := generator.CamelCase(origServName)
	deprecated := service.GetOptions().GetDeprecated()

	g.P()
	g.P(fmt.Sprintf(`// %sClient is the client API for %s service.`, servName, servName))

	// Client interface.
	if deprecated {
		g.P("//")
		g.P(deprecationComment)
	}
	for _, method := range service.Method {
		outType := g.typeName(method.GetOutputType())

		g.P("type ", servName, "_", method.GetName(), "_Exception = func(int32)")
		g.P("type ", servName, "_", method.GetName(), "_Callback = func(*", outType, ")")
	}
	g.P()

	// Client structure.
	g.P("type ", servName, "Client struct {")
	g.P("conn *stnet.Connect")
	g.P("rpcclient *gpbrpc.RPCClient")
	g.P("content map[string]string")
	g.P("timeout uint32")
	g.P("hashfunc func(*gpbrpc.RPCResponsePacket) int")
	g.P("}")
	g.P()

	// NewClient factory.
	if deprecated {
		g.P(deprecationComment)
	}
	g.P(fmt.Sprintf(`// new client for %s service.`, servName))
	g.P("func New", servName, "Client () *", servName, "Client {")
	g.P("return &", servName, "Client{nil, nil, make(map[string]string), 5000, nil}")
	g.P("}")
	g.P()

	g.P("// Change the content of this client which is send to service.")
	g.P("func (c *", servName, "Client) SetContent(k, v string) {")
	g.P("c.content[k] = v")
	g.P("}")
	g.P()
	g.P("func (c *", servName, "Client) DelContent(k string) {")
	g.P("delete(c.content, k)")
	g.P("}")
	g.P()
	g.P("func (c *", servName, "Client) GetContent() map[string]string {")
	g.P("return c.content")
	g.P("}")
	g.P()
	g.P("// set timeout of the response from service.")
	g.P("func (c *", servName, "Client) SetTimeout(t uint32) {")
	g.P("c.timeout = t")
	g.P("}")
	g.P()
	g.P("func (c *", servName, "Client) GetTimeout() uint32 {")
	g.P("return c.timeout")
	g.P("}")
	g.P()
	g.P("// set rpcclient.")
	g.P("func (c *", servName, "Client) SetRPCClient(rpc *gpbrpc.RPCClient) {")
	g.P("c.rpcclient = rpc")
	g.P("}")
	g.P()
	g.P("// set connection which is connecting the service.")
	g.P("func (c *", servName, "Client) SetConn(conn *stnet.Connect) {")
	g.P("c.conn = conn")
	g.P("}")
	g.P()
	g.P("func (c *", servName, "Client) GetConn() *stnet.Connect {")
	g.P("return c.conn")
	g.P("}")
	g.P()
	g.P("// decide which thread is used to handle the response.")
	g.P("func (c *", servName, "Client) SetHashFunc(f func(*gpbrpc.RPCResponsePacket) int) {")
	g.P("c.hashfunc = f")
	g.P("}")
	g.P()
	g.P("func (c *", servName, "Client) HashProcessor(rsp *gpbrpc.RPCResponsePacket) int {")
	g.P("if c.hashfunc!=nil{")
	g.P("return c.hashfunc(rsp)")
	g.P("}")
	g.P("return -1")
	g.P("}")
	g.P()

	g.P("// async call functions")
	for _, method := range service.Method {
		inType := g.typeName(method.GetInputType())

		g.P("func (_c *", servName, "Client) ", method.GetName(), "(in *", inType, ",cb ", servName, "_", method.GetName(), "_Callback", ",exp ", servName, "_", method.GetName(), "_Exception){")
		g.P("buf, _ := proto.Marshal(in)")
		g.P(`req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "`, servName, `", FuncName: "`, method.GetName(), `", ReqPayload: buf, Context: _c.GetContent()}, Callback: cb, Exception: exp, Handle: _c}`)
		g.P("if cb == nil && exp == nil {")
		g.P("req.Req.IsOneWay = true")
		g.P("}")
		g.P("_c.rpcclient.PushRequest(req)")
		g.P("}")
		g.P()
	}

	g.P("// sync call functions")
	for _, method := range service.Method {
		inType := g.typeName(method.GetInputType())
		outType := g.typeName(method.GetOutputType())

		g.P("func (_c *", servName, "Client) ", method.GetName(), "_Sync(in *", inType, ")(out *", outType, ",ret int32){")
		g.P(`out=&`, outType, "{}")
		g.P("buf, _ := proto.Marshal(in)")
		g.P("sg := make(chan *gpbrpc.RPCResponsePacket, 1)")
		g.P(`req := &gpbrpc.RPCRequest{Req: gpbrpc.RPCRequestPacket{ServiceName: "`, servName, `", FuncName: "`, method.GetName(), `", ReqPayload: buf, Context: _c.GetContent()}, Signal: sg, Handle: _c}`)
		g.P("_c.rpcclient.PushRequest(req)")
		g.P("to := time.NewTimer(time.Duration(_c.GetTimeout()) * time.Millisecond)")
		g.P("select {")
		g.P("case s := <-sg:")
		g.P("ret = s.RPCRetCode")
		g.P(`if len(s.RspPayload) > 0 {`)
		g.P(`if err := proto.Unmarshal(s.RspPayload, out); err != nil {`)
		g.P("ret = gpbrpc.GPBUnmarshalFailed")
		g.P("}")
		g.P("}")
		g.P("case <-to.C:")
		g.P("ret = gpbrpc.SyncCallTimeout")
		g.P("}")
		g.P("to.Stop()")
		g.P("_c.rpcclient.DeleteRequest(req.Req.RequestId)")
		g.P("return out, ret")
		g.P("}")
		g.P()
	}

	g.P("func (_c *", servName, "Client) HandleRSP(r *gpbrpc.RPCRequest, s *gpbrpc.RPCResponsePacket) {")
	if len(service.Method) > 0 {
		for i, method := range service.Method {
			outType := g.typeName(method.GetOutputType())

			if i == 0 {
				g.P(`if r.Req.FuncName == "`, method.GetName(), `" {`)
			} else {
				g.P(`} else if r.Req.FuncName == "`, method.GetName(), `" {`)
			}
			g.P("if s.RPCRetCode != 0 {")
			g.P(`if r.Exception.(`, servName, "_", method.GetName(), "_Exception) != nil {")
			g.P(`r.Exception.(`, servName, "_", method.GetName(), "_Exception)(s.RPCRetCode)")
			g.P(`}`)
			g.P(`} else {`)
			g.P(`out:=&`, outType, "{}")
			g.P(`if err := proto.Unmarshal(s.RspPayload, out); err != nil {`)
			g.P(`if r.Exception.(`, servName, "_", method.GetName(), "_Exception) != nil {")
			g.P(`r.Exception.(`, servName, "_", method.GetName(), "_Exception)(gpbrpc.GPBUnmarshalFailed)")
			g.P(`}`)
			g.P(`} else {`)
			g.P(`if r.Callback.(`, servName, "_", method.GetName(), "_Callback) != nil {")
			g.P(`r.Callback.(`, servName, "_", method.GetName(), "_Callback)(out)")
			g.P(`}`)
			g.P("}")
			g.P("}")
		}
	}
	g.P("}")
	g.P("}")
	g.P()

	// Server interface.
	serverType := servName + "Server"
	g.P("// ", serverType, " is the server API for ", servName, " service.")
	if deprecated {
		g.P("//")
		g.P(deprecationComment)
	}
	g.P("type ", serverType, " interface {")
	for i, method := range service.Method {
		g.gen.PrintComments(fmt.Sprintf("%s,2,%d", path, i)) // 2 means method in a service.
		g.P(g.generateServerSignature(servName, method))
	}
	g.P("//Get msg processor by hash")
	g.P("HashProcessor(req *gpbrpc.RPCRequestPacket) int")
	g.P("//each thread has one handle")
	g.P("NewHandle(*gpbrpc.RPCHelper) " + serverType)
	g.P("}")
	g.P()

	g.P("type ", servName, " struct {")
	g.P("*gpbrpc.RPCHelper")
	g.P(serverType)
	g.P("}")
	g.P()
	// Server registration.
	if deprecated {
		g.P(deprecationComment)
	}
	g.P("func New", servName, "Server(s ", serverType, ") *", servName, "{")
	g.P("return &", servName, "{&gpbrpc.RPCHelper{}, s}")
	g.P("}")
	g.P()

	g.P("func (s *", servName, ") NewHandle() gpbrpc.ServerInterface {")
	g.P("h := &gpbrpc.RPCHelper{}")
	g.P("return &", servName, "{h, s.", servName, "Server.NewHandle(h)}")
	g.P("}")
	g.P()

	g.P("func (s *", servName, ") HandleReq(req *gpbrpc.RPCRequestPacket) (msg []byte, ret int32, err error) {")
	if len(service.Method) > 0 {
		for i, method := range service.Method {
			inType := g.typeName(method.GetInputType())

			if i == 0 {
				g.P(`if req.FuncName == "`, method.GetName(), `" {`)
			} else {
				g.P(`} else if req.FuncName == "`, method.GetName(), `" {`)
			}
			g.P("in:=&", inType, "{}")
			g.P(`if err = proto.Unmarshal(req.ReqPayload, in); err != nil {`)
			g.P(`ret = gpbrpc.GPBUnmarshalFailed`)
			g.P(`err = gpbrpc.ErrGPBUnmarshalFailed`)
			g.P(`} else {`)
			g.P("msg,err = proto.Marshal(s.", method.GetName(), "(in))")
			g.P("}")
		}
		g.P("}else{")
		g.P(`ret = gpbrpc.CallRemoteNullFunc`)
		g.P(`err = gpbrpc.ErrCallRemoteNullFunc`)
		g.P("}")
	}
	g.P("return msg, ret, err")
	g.P("}")
	g.P()
}

// generateServerSignature returns the server-side signature for a method.
func (g *grpc) generateServerSignature(servName string, method *pb.MethodDescriptorProto) string {
	origMethName := method.GetName()
	methName := generator.CamelCase(origMethName)
	if reservedClientName[methName] {
		methName += "_"
	}

	var reqArgs []string
	var ret string
	if method.GetOutputType() != "" {
		ret = "(*" + g.typeName(method.GetOutputType()) + ")"
	}
	if method.GetInputType() != "" {
		reqArgs = append(reqArgs, "*"+g.typeName(method.GetInputType()))
	}

	return methName + "(" + strings.Join(reqArgs, ", ") + ") " + ret
}
