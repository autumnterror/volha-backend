
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
cddb:
	docker compose -f ./build/volha/docker-compose.db.yml down
cddb-v:
	docker compose -f ./build/volha/docker-compose.db.yml down -v
cu:
	docker compose -f ./build/volha/docker-compose.yml up -d
cd:
	docker compose -f ./build/volha/docker-compose.yml down

bldmig:
	go build -o ./build/volha/migrator.exe ./cmd/product-service/migrator
bldmig-mac:
	go build -o ./build/volha/migrator ./cmd/product-service/migrator
docx:
	swag init --dir ./cmd/gateway,./internal/gateway/net/handlers,./pkg/views,./pkg/proto/gen --output ./docs
bld-a:
	docker build -t zitrax78/volha-gateway --file ./build/gateway/dockerfile .
	docker build -t zitrax78/product-service --file ./build/product-service/dockerfile .
	docker build -t zitrax78/dumper --file ./build/dumper/dockerfile .
bld-a-mac:
	docker build --platform linux/amd64 -t zitrax78/volha-gateway --file ./build/gateway/dockerfile .
	docker build --platform linux/amd64 -t zitrax78/product-service --file ./build/product-service/dockerfile .
	docker build --platform linux/amd64 -t zitrax78/dumper --file ./build/dumper/dockerfile .
push-a:
	docker push zitrax78/volha-gateway
	docker push zitrax78/product-service
	docker push zitrax78/dumper

#-------------------------test-------------------------
t-m:
	go test -run $(METHOD) ./... -v
t-m-h:
	go test -run TestGetDictionaries  ./... -v
t-psql:
	go test ./internal/product-service/psql -v
t-red:
	go test ./internal/gateway/redis -v
