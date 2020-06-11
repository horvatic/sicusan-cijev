#!/bin/sh

REPO_URI="$1"
SERVICE="$2"

if [ -d "$HOME/go/src/$SERVICE" ]; then
	CURRENT_PID=$(cat "$HOME/go/src/$SERVICE/.pid.info")
	kill -9 "$CURRENT_PID"
	rm -rf "$HOME/go/src/$SERVICE"
fi

git clone "$REPO_URI" "$HOME/go/src/$SERVICE"
cd "$HOME/go/src/$SERVICE"
go build main.go
mv main "$SERVICE"
chmod +x "$SERVICE"
./"$SERVICE" &
echo "$!" >> "$HOME/go/src/$SERVICE/.pid.info"
