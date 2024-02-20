GO_WORKSPACE := .

protoc: 
	protoc --proto_path=protos protos/*.proto --go_out=$(GO_WORKSPACE) --go-grpc_out=$(GO_WORKSPACE)
	@echo "Protoc Compile Finish!"