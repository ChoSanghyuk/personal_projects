#!/bin/sh
cd ~/workspace/invest_indicator
git pull origin main
/home/cho/go/bin/go build -o invest_indicator
PID=$(cat pidfile.txt 2>/dev/null)
if kill -0 $PID 2>/dev/null; then echo "Process $PID is running. Kill process"; kill $PID; fi
./invest_indicator &
echo $! > pidfile.txt

