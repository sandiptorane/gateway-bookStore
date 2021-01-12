GRPC:
	protoc --proto_path=. --go_out=plugins=grpc:. --go_opt=paths=source_relative ./bookproto/book.proto
	protoc --proto_path=. --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt=paths=source_relative ./bookproto/book.proto