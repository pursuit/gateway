FROM golang:alpine3.13 as builder

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main .

ENTRYPOINT ["/main"]
