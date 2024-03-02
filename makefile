build:
	docker-compose --env-file .env -f docker-compose.yml build --no-cache
start:
	docker-compose --env-file .env -f docker-compose.yml up -d
stop:
	docker-compose stop
restart:
	@make stop
	@make start
logs:
	docker-compose logs api -f --tail 100
init:
	@cp .env.example .env
	@make download-tools      # ツールをインストール
	@make build
	@make start
	@make swag # swaggerの初期化
	@make refresh-schema
destroy:
	@make stop
	rm -rf persist/
	docker-compose down --rmi local --volumes --remove-orphans
ps:
	docker-compose ps
shell:
	docker-compose exec api sh
tidy:
	docker-compose exec api go mod tidy
mod-download:
	go mod download
# sqlファイルをコンテナに流す
exec-schema:
	cat ./DDL/*.up.sql > ./DDL/schema.sql
	docker cp DDL/schema.sql db:/ && docker exec -it db psql -U postgres -d todo_app -f /schema.sql
	docker cp DDL/schema.sql db-test:/ && docker exec -it db-test psql -U test -d todo_app_test -f /schema.sql
	rm ./DDL/schema.sql
exec-dummy:
	docker cp DDL/insert_dummy_data.sql db:/ && docker exec -it db psql -U postgres -d todo_app -f /insert_dummy_data.sql
# テーブルをリフレッシュする
refresh-schema:
	@make exec-schema
	@make exec-dummy
	@make exec-dummy-appointments
# テストを実行して、カバレッジを出力する
coverage:
	docker compose exec api go test -coverpkg=./... -coverprofile=coverage.txt ./...
	docker compose exec api go tool cover -html=coverage.txt -o coverage.html
# テストを実行する
test:
	@docker compose exec api go test ./...
# フォーマットを整える
fmt:
	find . \( -type d -path './internal/domain/repository/mock' -o -path './internal/domain/repository/bigadvance/mock' \) -prune -o -name '*.go' -print | xargs goimports -l -w -local "github.com/keito-isurugi/todo-app-backend"
	find . \( -type d -path './internal/domain/repository/mock' -o -path './internal/domain/repository/bigadvance/mock' \) -prune -o -name '*.go' -print | xargs gofmt -l -w
# mockを生成
mockgen:
	go generate ./...
# ツールをインストール
download-tools:
	@go install github.com/golang/mock/mockgen@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/k1LoW/tbls@latest
lint:
	golangci-lint run ./...
# swaggerを生成
swag:
	swag init -g ./cmd/api/main.go -o ./swagger/src/
# tblsを使ってスキーマを生成
generate-schema:
	docker run --rm --network=todo-app-network -v $(PWD):/work -w /work ghcr.io/k1low/tbls doc --force
#ダミーの予約データを流す
exec-dummy-appointments:
	docker compose exec api go run cmd/seeder/main.go
