.DEFAULT_GOAL := help
.PHONY: setup gen build run fmt lint help

setup: ## 開発に必要なツールをインストールする
	brew install sqldef/sqldef/mysqldef
	go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

gen: ## コードを生成する
	@buf generate
	@$(MAKE) fmt

fmt: ## コードを整形する
	@goimports -w .

lint: ## 静的解析を実行する
	@go vet $$(go list ./... | grep -v -e /teraterapb)
	@staticcheck $$(go list ./... | grep -v -e /teraterapb)

help: ## ヘルプを表示する
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) \
      | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-14s\033[0m %s\n", $$1, $$2}'
