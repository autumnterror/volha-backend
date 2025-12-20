echo off

cd ../dumps

set CONTAINER_NAME=volha-productsdb
set DB_NAME=productsdb
set DB_USER=postgres

docker exec -i %CONTAINER_NAME% psql -U %DB_USER% -d %DB_NAME% < "./demo.sql"