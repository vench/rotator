V := @

# Build
OUT_DIR = ./bin

.PHONY: vendor
vendor:
	$(V)go mod tidy -compat=1.17
	$(V)go mod vendor

.PHONY: build
build:
	$(V)CGO_ENABLED=1 go build -o ${OUT_DIR}/app ./cmd/rotator/main.go