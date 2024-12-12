FROM golang:latest

WORKDIR /app
COPY . .
WORKDIR /app/cmd/resource
RUN go build -o /app/resource-server .

EXPOSE 8082
CMD ["/app/resource-server"]
