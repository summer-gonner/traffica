goctl rpc protoc sys/sys.proto --go_out=./sys/ --go-grpc_out=./sys/ --zrpc_out=./sys/ -m
goctl api go -api admin/doc/api/admin.api -dir admin admin/doc/api/admin.api