#!/bin/sh

# MUST
# export GO111MODULE=on
CC=i686-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=386 go build \
    -tags="bdebug" \
    -ldflags="-H=windowsgui" \
    -o main.exe


#CC=i686-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=386 go build \
#    -tags="bdebug" \
#    -o main.exe
