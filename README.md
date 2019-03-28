# ftserver
Simple CLI Quiz using Cobra and gRPC

Command to compile proto files
protoc --proto_path=proto --go_out=plugins=grpc:proto questions.proto
