#!/bin/bash

cd ../dumps

CONTAINER_NAME=volha-productsdb
DB_NAME=productsdb
DB_USER=postgres

export CONTAINER_NAME DB_NAME DB_USER

docker exec -i $CONTAINER_NAME psql -U $DB_USER -d $DB_NAME < "./demo.sql"

unset CONTAINER_NAME DB_NAME DB_USER
