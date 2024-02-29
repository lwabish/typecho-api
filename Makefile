GEN_TOOL = gentool

ensure-gen-tool:
	@if ! which $(GEN_TOOL) >/dev/null; then \
	echo "Install $(GEN_TOOL) to generate models"; \
	go install gorm.io/gen/tools/gentool@latest; fi

gen-model: ensure-gen-tool
	rm -f models/*.gen.go
	$(GEN_TOOL) -c ./models/gorm-gen.yml