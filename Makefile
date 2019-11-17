protos:
	protoc -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:./generated \
		--swagger_out=logtostderr=true:. \
		--go_out=plugins=grpc:./generated \
		api/api.proto
