FROM golang:1.24-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOFLAGS="-mod=mod"
RUN go build -o car ./cmd/main.go

VOLUME ["/app/uploads"]

CMD ["./car"]