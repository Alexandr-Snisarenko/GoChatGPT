
# syntax=docker/dockerfile:1
FROM golang:1.22

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
