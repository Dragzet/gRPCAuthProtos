.PHONY: build

build:
	protoc -I proto proto/auth/auth.proto --go_out=./gen/go --go_opt=paths=source_relative --go-grpc_out=./gen/go

.DEfAULT_GOAL := build