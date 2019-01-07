
FROM golang:1.11.1-alpine3.8
WORKDIR /go/src/github.com/dinhvandung7541/vote-now
ADD . /go/src/github.com/dinhvandung7541/vote-now

COPY --from=0 /go/src/github.com/dinhvandung7541/vote-now/ /vote-now

CMD go run cmd/vote-now/main.go
