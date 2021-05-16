## Ledger - Protos

Protocol buffer files and generated Go code for use in the Ledger project.

## Tool Installation

```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

## Generating Files

1. Enter a package directory, in this example `session`
2. Run `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./session.proto`

