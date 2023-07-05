#!/bin/bash

mkdir -p $HOME/workspace/mysql_data

docker compose down

docker compose build

docker compose up -d
