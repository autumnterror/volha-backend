#!/bin/bash

cd ..
CONFIG_PATH=./configs/migrator.yaml
export CONFIG_PATH
./migrator --type down
unset CONFIG_PATH
