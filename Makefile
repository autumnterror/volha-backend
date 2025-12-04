
gen:
	protoc -I ./pkg/proto/filez/ \
	  products.proto \
	  domain.proto \
	  --go_out=./pkg/proto/gen/ \
	  --go_opt=paths=source_relative \
	  --go-grpc_out=./pkg/proto/gen/ \
	  --go-grpc_opt=paths=source_relative
cudb:
	docker compose -f ./build/volha/docker-compose.db.yml up -d
cddb-v:
	docker compose -f ./build/volha/docker-compose.db.yml down -v
bldmig:
	go build -o ./build/volha/migrator.exe ./cmd/product-service/migrator
bldmig-mac:
	go build -o ./build/volha/migrator ./cmd/product-service/migrator

#-------------------------test-------------------------
t-m:
	go test -run $(METHOD) ./... -v
t-m-h:
	go test -run TestGetDictionaries  ./... -v
t-psql:
	go test ./internal/product-service/psql -v
