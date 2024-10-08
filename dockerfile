FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
FROM ubuntu:latest
WORKDIR /app
COPY --from=build /app/main .
EXPOSE 8080
RUN chmod +x /app/main
CMD ["./main"]
