# gpbrpc
rpc using protocol of google protocol buffer

protoc-gen-go  generate rpc files
	protoc.exe --go_out=plugins=gpbrpc:. grpc.proto