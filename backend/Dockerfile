FROM golang:1-alpine

# Required to build gccgo packages (sqlite)
RUN apk add build-base
# Required to replace env variables in config.json
RUN apk add gettext

WORKDIR /go/src/app
COPY . .

RUN go mod download -x
RUN go install -v ./...

ENTRYPOINT [ "sh", "docker-backend-entry.sh" ]