#!/bin/sh

git clone "$1" ~/go/src/
cd "~/go/src/$2"
go build main.go
mv main.go "$2"
./"$2" &

