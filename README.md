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

Run `make pkg=<package name> generate`

