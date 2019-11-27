protos:
	protoc -I. --go_out=plugins=grpc:./generated api/api.proto
