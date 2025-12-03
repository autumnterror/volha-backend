gen:
	protoc -I ./pkg/proto/filez/ \
	  products.proto \
	  domain.proto \
	  --go_out=./pkg/proto/gen/ \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/proto/gen/ \
	  --go-grpc_opt=paths=source_relative
