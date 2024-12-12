FROM golang:latest

WORKDIR /app
COPY . .
WORKDIR /app/cmd/client
RUN go build -o /app/client-server .

EXPOSE 8081
CMD ["/app/client-server"]
