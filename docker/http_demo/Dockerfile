FROM golang:alpine AS builder

ENV GO11MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go build -o app .

FROM scratch

COPY --from=builder /build/app /

EXPOSE 9090

CMD ["/app"]