#!/bin/sh

SERVICE="$2-userprocess"

if [ -d "$HOME/go/src/$SERVICE" ]; then
	rm -rf "$HOME/go/src/$SERVICE"
	killall -9 "$SERVICE"
fi

git clone "$1" "$HOME/go/src/$SERVICE"
cd "$HOME/go/src/$SERVICE"
go build main.go
mv main "$SERVICE"
chmod +x "$SERVICE"
./"$SERVICE" &

