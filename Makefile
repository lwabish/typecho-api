GEN_TOOL = gentool
GOBIN ?= $(GOPATH)/bin

ensure-gen-tool:
	@if ! which $(GEN_TOOL) >/dev/null; then \
	echo "Install $(GEN_TOOL) to generate models"; \
	go install gorm.io/gen/tools/gentool@latest; fi

gen-model: ensure-gen-tool
	rm -f models/*.gen.go
	$(GEN_TOOL) -c ./models/gorm-gen.yml

setup-env:
	docker compose up --build -d

teardown-env:
	docker compose down -v

run-test:
	go test -v ./...

test: setup-env run-test teardown-env

install-client:
	go build -o output/typora-client client/main.go
	ln -sf $(shell pwd)/output/typora-client $(GOBIN)/typora-client
