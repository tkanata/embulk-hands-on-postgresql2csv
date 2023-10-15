export PSQL_ENDPOINT=localhost
export PSQL_USER=postgres
export PSQL_PASSWORD=password
export PSQL_DBNAME=embulk-handson
export PSQL_PORT=5432
export PSQL_TZ=Asia/Tokyo

.PHONY: build run

# バイナリの出力名を指定します。
OUTPUT := tea_importer

# Goのビルドコマンド
build:
	go build -o $(OUTPUT) main.go

# Goのプログラムを直接実行するコマンド
run: build
	./$(OUTPUT)

# ビルドで生成されたファイルを削除するコマンド
clean:
	rm -f $(OUTPUT)

init:
	@echo "#---------------------------#"
	@echo "# dockerを起動"
	@echo "# 必要なツールをインストール、DBを最新化します"
	@echo "# ent: ORM"
	@echo "#---------------------------#"
	docker-compose up -d
	go get -d entgo.io/ent/cmd/ent
	go generate ./ent
	brew install ariga/tap/atlas
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "postgres://postgres:password@localhost:5432/embulk-handson?search_path=public&sslmode=disable"
	$(MAKE) run

init-arm64:
	@echo "#---------------------------#"
	@echo "# dockerを起動"
	@echo "# 必要なツールをインストール、DBを最新化します"
	@echo "# ent: ORM"
	@echo "#---------------------------#"
	docker-compose up -d
	go get -d entgo.io/ent/cmd/ent
	go generate ./ent
	arch -arm64 brew install ariga/tap/atlas
	atlas migrate apply --dir "file://ent/migrate/migrations" --url "postgres://postgres:password@localhost:5432/embulk-handson?search_path=public&sslmode=disable"
	$(MAKE) run