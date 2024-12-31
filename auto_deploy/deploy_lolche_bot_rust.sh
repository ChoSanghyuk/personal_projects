#!/bin/sh
cd ~/workspace/lolche_bot_rust
git pull origin master
cargo build
PID=$(cat pidfile.txt 2>/dev/null)
if kill -0 $PID 2>/dev/null; then echo "Process $PID is running. Kill process"; kill $PID; fi
./target/debug/lolche_bot_rust &
echo $! > pidfile.txt

