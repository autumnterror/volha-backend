#!/bin/bash

cd ..
CONFIG_PATH=./configs/migrator.yaml
export CONFIG_PATH
./migrator --type up
unset CONFIG_PATH
