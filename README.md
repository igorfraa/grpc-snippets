# grpc-snippets

# ts+vue 
```sh
protoc --plugin=protoc-gen-grpc-web=frontend/node_modules/.bin/protoc-gen-grpc-web \
  		--grpc-web_out=import_style=commonjs+dts,mode=grpcweb:frontend/src/lib/ \
  		--js_out=import_style=commonjs:frontend/src/lib/ \
  		-I internal/model/grpc \
  		protobuf.proto
```

# go (easy as in docs) 
```go
//go:generate protoc --go_out=internal/model/grpc --go_opt=paths=source_relative --go-grpc_out=internal/model/grpc --go-grpc_opt=paths=source_relative --proto_path=internal/model/grpc service.proto
```
