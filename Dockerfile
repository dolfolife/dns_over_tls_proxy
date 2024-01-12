FROM golang:latest

WORKDIR /proxy

COPY . .

RUN go build -o dns-over-tls-proxy .

EXPOSE 53/tcp
EXPOSE 53/udp

CMD ["./dns-over-tls-proxy"]
