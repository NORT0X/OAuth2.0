FROM golang:latest

WORKDIR /app
COPY . .
WORKDIR /app/cmd/auth
RUN go build -o /app/auth-server .

EXPOSE 8080
CMD ["/app/auth-server"]

