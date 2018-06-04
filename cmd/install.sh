#!/usr/bin/env bash

type go >/dev/null 2>&1 || { echo >&2 "Go runtime is required."; exit 1; }
type npm >/dev/null 2>&1 || { echo >&2 "Node.js is required in your system."; exit 1; }

go get -u github.com/gorilla/handlers
go get -u github.com/gorilla/mux
go get -u github.com/rrborja/shift-schedule

npm install --cwd "../www" --prefix "../www"

d="6.4.2018"

mkdir -p $d
if [ ! -f 1.txt ]; then printf "Jack:17:0:6" > $d/1.txt; fi
if [ ! -f 2.txt ]; then printf "Jill:12:6:10" > $d/2.txt; fi
if [ ! -f 3.txt ]; then printf "Jim:4:26:48" > $d/3.txt; fi
if [ ! -f 4.txt ]; then printf "John:7:10:19" > $d/4.txt; fi
if [ ! -f 5.txt ]; then printf "John:7:19:23" > $d/5.txt; fi

exec go run main.go