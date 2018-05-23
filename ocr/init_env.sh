#!/usr/bin/env bash

env=$1
if [ ! $env ];then
env=dev
fi

go run db/db.go --action=drop --env=$env
go run db/db.go --action=create --env=$env
go run db/db.go --action=migrate --env=$env
