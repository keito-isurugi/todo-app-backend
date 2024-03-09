FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download \
    && go build -o /app/main /app/cmd/api/main.go

FROM nginx:latest

# Nginxの設定ファイルをコピー
COPY nginx/nginx.conf /etc/nginx/nginx.conf

# ビルドしたGoアプリをコピー
COPY --from=builder /app/main /usr/share/nginx/html

# Nginxを起動
CMD ["nginx", "-g", "daemon off;"]
