#!/bin/sh

REPO_STORE=""

git clone "$1" "$REPO_STORE/go/src/$2"
cd "$REPO_STORE/go/src/$2"
go build main.go
mv main "$2"
chmod +x "$2"
./"$2" &

