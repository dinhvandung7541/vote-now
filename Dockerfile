FROM golang:1.11.1-alpine3.8
WORKDIR /go/src/github.com/dinhvandung7541/vote-now
ADD . /go/src/github.com/dinhvandung7541/vote-now
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o vote-now ./ext/cmd/vote-now

FROM alpine:3.8
WORKDIR /
RUN apk add --no-cache ca-certificates

COPY --from=0 /go/src/github.com/dinhvandung7541/vote-now/vote-now /vote-now

CMD ["/vote-now"]