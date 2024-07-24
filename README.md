# grpc-snippets
```sh
go generate ./...
```

# ts+vue 
```sh
npm i --save protoc-gen-grpc-web
protoc --plugin=protoc-gen-grpc-web=frontend/ping-example/node_modules/.bin/protoc-gen-grpc-web \
  		--grpc-web_out=import_style=commonjs+dts,mode=grpcweb:frontend/ping-example/src/lib/ \
  		--js_out=import_style=commonjs:frontend/ping-example/src/lib/ \
  		-I proto/ \
  		ping.proto
```

# go (easy as in docs) 
```go
//go:generate protoc --go_out=internal/model/grpc --go_opt=paths=source_relative --go-grpc_out=internal/model/grpc --go-grpc_opt=paths=source_relative --proto_path=internal/model/grpc service.proto
```
