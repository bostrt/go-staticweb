FROM golang:onbuild

EXPOSE 8080
VOLUME ["/go/src/app/files"]
