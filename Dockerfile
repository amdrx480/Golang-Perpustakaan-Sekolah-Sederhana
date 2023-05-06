FROM golang:1.19-alpine

WORKDIR /app


COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
RUN go mod tidy
RUN go build -o /api

EXPOSE 8080

CMD ["/api"]