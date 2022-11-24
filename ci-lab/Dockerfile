ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine

WORKDIR /app

COPY go.mod ./
COPY ./cmd/cat-service/ ./
RUN go mod download

RUN go build main.go

CMD [ "./main" ]


