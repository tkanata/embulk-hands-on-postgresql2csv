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