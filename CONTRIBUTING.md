# Contributing to kraft-ui

Thank you for considering a contribution. kraft-ui is MIT licensed and welcomes
community improvements — new components, bug fixes, better patterns, and docs.

---

## Prerequisites

| Tool          | Version   | Purpose                     |
|---------------|-----------|-----------------------------|
| Go            | 1.22+     | Build and test              |
| templ         | latest    | Compile `.templ` files      |
| Node.js       | 18+       | Tailwind CSS build          |
| air           | latest    | Hot reload in development   |

Install Go tools:

```bash
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/air-verse/air@latest
```

---

## Development setup

```bash
git clone https://github.com/appkrafterrz/kraft-ui
cd kraft-ui
make setup        # npm install + go mod tidy
make dev          # hot reload server at http://localhost:3000
```

`make dev` runs three watchers in parallel:

- `air` — recompiles and restarts the Go server on `.go` or `.templ` changes
- `templ generate --watch` — recompiles `.templ` files to `_templ.go`
- `npm run css:watch` — rebuilds Tailwind CSS on template changes

---

## Build without hot reload

```bash
make build        # produces ./bin/kraft-ui-server
make run          # build and run in one step
```

**Important:** after editing any `.templ` file you must run all three steps before
changes appear in the browser:

```bash
templ generate ./...      # compile .templ → _templ.go
npm run css:build         # rebuild Tailwind (picks up new utility classes)
# restart the server
```

---

## Adding a component

1. Create `pkg/ui/<name>/<name>.templ` with a `Props` struct and a `templ` function.
2. Run `templ generate` to produce `<name>_templ.go`.
3. Add a doc page at `web/pages/docs/<name>.templ` with a live preview and props table.
4. Register the route in `cmd/server/main.go` and the handler in `internal/handler/docs.go`.
5. Add the component to the playground in `web/pages/playground.templ` and `web/pages/playground_repl.templ`.
6. Run `npm run css:build` to include any new Tailwind classes.

---

## Adding a DaisyUI wrapper

1. Add the component function to `pkg/daisyui/<name>.templ`.
2. Follow the same doc and playground steps as above, using the `daisy-` prefix for playground IDs.

---

## Code style

- `gofmt -s` and `templ fmt` before committing.
- Props structs use exported fields with Go-idiomatic names (`Variant`, `Size`, `Disabled`).
- Variants and sizes are typed string constants — never bare strings.
- Pass arbitrary HTML attributes via `templ.Attributes` so HTMX and Alpine.js work out of the box.

---

## Pull requests

- Keep PRs focused — one component or one fix per PR.
- Include a live preview in the docs page for new components.
- Generated `_templ.go` files should be committed alongside the `.templ` source.
