FROM golang:latest

ENV API_KEY=DEMO_KEY \
    PORT=8080 \
    CONCURRENT_REQUESTS=5 \
    TRANSPORT_TIMEOUT=5 \
    HANDSGAKE_TIMEOUT=5 \
    CLIENT_TIMEOUT=10

COPY . /go/src/app

WORKDIR /go/src/app/cmd/gogospace

RUN go build -o gogospace main.go

EXPOSE 8080

CMD ["./gogospace"]