goctl rpc protoc sys/sys.proto --go_out=./sys/ --go-grpc_out=./sys/ --zrpc_out=./sys/ -m
goctl api go -api admin/doc/api/admin.api -dir admin admin/doc/api/admin.api
goctl rpc protoc record/record.proto --go_out=./record/ --go-grpc_out=./record/ --zrpc_out=./record/ -m

goctl api go -api agent/doc/api/agent.api -dir agent agent/doc/api/agent.api
