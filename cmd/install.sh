#!/usr/bin/env bash

clear

printf "Setting up Shift Scheduler's core package and test data\n\n"

type go >/dev/null 2>&1 || { echo >&2 "Go runtime is required."; exit 1; }
type npm >/dev/null 2>&1 || { echo >&2 "Node.js is required in your system."; exit 1; }

g=`go version`
n=`node -v`

printf "Runtime environments:\n$g\nnode.js $n\n\n"

go get -u github.com/gorilla/handlers
go get -u github.com/gorilla/mux
go get -u github.com/rrborja/shift-schedule

npm install --cwd "../www" --prefix "../www"

echo

d=`date +%-m.%-d.%Y`

if [ ! -d "$d" ]; then
    echo "Test date will be populated for today's shift. To delete Today's shift, remove:"
    echo $PWD"/$d/"
    echo

    mkdir -p $d
    if [ ! -f 1.txt ]; then printf "Jack:17:0:6" > $d/1.txt; fi
    if [ ! -f 2.txt ]; then printf "Jill:12:6:10" > $d/2.txt; fi
    if [ ! -f 3.txt ]; then printf "Jim:4:26:48" > $d/3.txt; fi
    if [ ! -f 4.txt ]; then printf "John:7:10:19" > $d/4.txt; fi
    if [ ! -f 5.txt ]; then printf "John:7:19:23" > $d/5.txt; fi
fi

echo "Installation done. Wait up to one minute before you may access:"
echo "http://localhost:3000"
echo

exec go run main.go