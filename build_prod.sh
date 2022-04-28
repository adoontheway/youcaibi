#! /bin/bash

# Build web and other service

cd ~/work/src/myworkpath/api
env GOOS=linux GOARCH=amd64 go build -o ../bin/api


cd ~/work/src/myworkpath/scheduler
env GOOS=linux GOARCH=amd64 go build -o ../bin/scheduler

cd ~/work/src/myworkpath/stream
env GOOS=linux GOARCH=amd64 go build -o ../bin/stream

cd ~/work/src/myworkpath/web
env GOOS=linux GOARCH=amd64 go build -o ../bin/web

# todo git push remote 