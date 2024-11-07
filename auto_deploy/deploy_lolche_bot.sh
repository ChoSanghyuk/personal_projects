#!/bin/sh
cd ~/workspace/lolche_bot
git pull origin main
/home/cho/go/bin/go build -o lolche_bot
PID=$(cat pidfile.txt 2>/dev/null)
if kill -0 $PID 2>/dev/null; then echo "Process $PID is running. Kill process"; kill $PID; fi
./lolche_bot &
echo $! > pidfile.txt

