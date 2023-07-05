#!/bin/bash

mkdir -p $HOME/mysql_data

docker compose down

docker compose build

docker compose up -d
