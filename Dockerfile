FROM golang:1.22-alpine

WORKDIR /app

COPY .env go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./... -v
RUN go build -o /pos-go-temperature-by-location cmd/server/main.go

EXPOSE 8080

CMD ["/pos-go-temperature-by-location"]
