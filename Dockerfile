FROM golang:alpine AS builder

WORKDIR $GOPATH/src/shortlink
COPY . .

RUN adduser -D -g '' shortlink

RUN CGO_ENABLED=0 go build -o /go/bin/shortlink


FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/shortlink /shortlink

USER shortlink

EXPOSE 8080

ENTRYPOINT ["/shortlink"]
