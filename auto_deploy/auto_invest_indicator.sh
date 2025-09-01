#!/bin/sh
cd ~/workspace/invest_indicator
git pull origin main
# /home/cho/go/bin/go build -o invest_indicator ./cmd/
docker run --rm -v $(pwd):/app -w /app golang:1.24 go build -o investindicator cmd/main.go
PID=$(cat pidfile.txt 2>/dev/null)
if kill -0 $PID 2>/dev/null; then echo "Process $PID is running. Kill process"; kill $PID; fi
nohup ./invest_indicator > /var/log/invest/$(date +%F).log 2>&1 &
echo $! > pidfile.txt
