1. gRPC
    ```
    protoc --go_out=address/pb --go_opt=paths=source_relative --go-grpc_out=address/pb --go-grpc_opt=paths=source_relative --proto_path=address/protofiles address/protofiles/*.proto
    ```

2. GraphQL
   ```
   go get github.com/99designs/gqlgen
   ```

   ```
   go run github.com/99designs/gqlgen init
   ```

   ```
   go run github.com/99designs/gqlgen generate
   ```
   