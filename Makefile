.PHONY: all stream unary
all: stream unary

stream:
	@protoc --go_out=./stream --go-grpc_out=./stream ./stream/*.proto

unary:
	@protoc --go_out=./unary --go-grpc_out=./unary ./unary/*.proto
