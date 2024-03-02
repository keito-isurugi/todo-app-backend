FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download \
    && go build -o /app/main /app/cmd/api/main.go

FROM golang:1.22-alpine

COPY --from=builder /app/main ./

CMD ["./main"]