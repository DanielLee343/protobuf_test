generate_chat:
	protoc -I. \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:. \
	chat.proto