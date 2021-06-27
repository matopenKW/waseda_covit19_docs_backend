FROM golang:1.14.15 as builder

WORKDIR /go/src

COPY ./  ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go 

FROM scratch as runner

COPY --from=builder /go/src/main /app/main

ENTRYPOINT ["/app/main"]