cd /d ".."
docker compose down -v
docker compose up -d
timeout /t 5 /nobreak
set CONFIG_PATH=.\configs\migrator.yaml
.\migrator.exe --type up
