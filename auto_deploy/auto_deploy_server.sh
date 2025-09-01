#!/bin/sh
cd ~/workspace/personal_projects/auto_deploy
git pull origin master
docker run --rm -v $(pwd):/app -w /app golang:1.24 go build -o autodeploy main.go
PID=$(cat pidfile.txt 2>/dev/null)
if kill -0 $PID 2>/dev/null; then echo "Process $PID is running. Kill process"; kill $PID; fi
./autodeploy &
echo $! > pidfile.txt