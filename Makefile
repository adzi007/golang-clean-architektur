create:
	protoc --proto_path=transport/grpc/proto transport/grpc/proto/*.proto --go_out=transport/grpc/gen/
	protoc --proto_path=transport/grpc/proto transport/grpc/proto/*.proto --go-grpc_out=transport/grpc/gen/
	
# protoc -I . --grpc-gateway_out ./gen/ --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true proto/test.proto

clean:
	rm transport/grpc/gen/proto/*.go