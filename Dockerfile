FROM golang:1.22.4-alpine

WORKDIR /nlw-journey

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

WORKDIR ./cmd/journey

RUN go build -o ./bin/journey .

EXPOSE 8080
ENTRYPOINT [ "./bin/journey" ]


