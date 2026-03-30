.PHONY: dev build gen lint css install setup clean

# ── Bootstrap ──────────────────────────────────────────────
setup: install
	go mod tidy
	@mkdir -p web/static/css web/static/js tmp

install:
	npm install

# ── Development ────────────────────────────────────────────
dev:
	@echo "→ Starting kraft-ui dev server..."
	@$(MAKE) -j3 dev/go dev/css dev/templ

dev/go:
	air

dev/css:
	npm run css:watch

dev/templ:
	templ generate --watch

# ── Build ──────────────────────────────────────────────────
build: gen css/build
	go build -o ./bin/kraft-ui-server ./cmd/server

gen:
	templ generate

css/build:
	npm run css:build

# ── Code quality ───────────────────────────────────────────
lint:
	golangci-lint run ./...

fmt:
	gofmt -s -w .
	templ fmt .

test:
	go test ./... -v

# ── Utilities ──────────────────────────────────────────────
clean:
	rm -rf ./bin ./tmp
	find . -name "*_templ.go" -delete

# ── Quick run (no hot reload) ──────────────────────────────
run: gen css/build
	go run ./cmd/server
