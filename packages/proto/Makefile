GOPATH := $(shell go env GOPATH)

all: go ts

go: go/messages/runtimefilter.pb.go go/messages/inspection.pb.go

ts: ts/messages/runtimefilter.pb.ts ts/messages/inspection.pb.ts

clean:
	rm -f go/messages/*.pb.go

go/messages/runtimefilter.pb.go:
	protoc \
	 	--go_opt=paths=source_relative \
		--go_out=./go \
		./messages/runtimefilter.proto

go/messages/inspection.pb.go:
	protoc \
	 	--go_opt=paths=source_relative \
		--go_out=./go \
		./messages/inspection.proto

ts/messages/runtimefilter.pb.ts:
	protoc \
		--plugin=protoc-gen-ts=../../node_modules/.bin/protoc-gen-ts_proto \
		--ts_out=./ts \
		./messages/runtimefilter.proto

ts/messages/inspection.pb.ts:
	protoc \
		--plugin=protoc-gen-ts=../../node_modules/.bin/protoc-gen-ts_proto \
		--ts_out=./ts \
		./messages/inspection.proto