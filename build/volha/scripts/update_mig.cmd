
docker compose down 

docker pull zitrax78/volha-gateway
docker pull zitrax78/product-service
docker pull zitrax78/dumper

docker compose up -d
timeout /t 5 /nobreak
set CONFIG_PATH=.\configs\migrator.yaml
.\migrator.exe --type up


