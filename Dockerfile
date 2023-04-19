FROM golang:1.20-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out ./cmd/main/main.go
EXPOSE 8080:8080
CMD ./out