.DEFAULT_GOAL := help
.PHONY: setup build run fmt lint help

setup: ## 開発に必要なツールをインストールする
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/tools/cmd/goimports@latest

build: ## 本番用APIサーバのコンテナイメージをビルドする
	@docker build \
            --tag=teratera-api:latest \
            --target=prod .

run: ## 本番用APIサーバを実行する
	@docker container run \
            --name teratera-api \
            -p 8080:8080 \
            --rm \
            teratera-api

fmt: ## コードを整形する
	@goimports -w .

lint: ## 静的解析を実行する
	@go vet $$(go list ./... | grep -v -e /teraterapb)
	@staticcheck $$(go list ./... | grep -v -e /teraterapb)

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
