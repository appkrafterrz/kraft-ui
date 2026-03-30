<div align="center">

# kraft-ui

**Typed UI components for the GOTH stack**

[![CI](https://github.com/appkrafterrz/kraft-ui/actions/workflows/ci.yml/badge.svg)](https://github.com/appkrafterrz/kraft-ui/actions/workflows/ci.yml)
[![Go](https://img.shields.io/badge/go-1.23-00ADD8?logo=go&logoColor=white)](https://go.dev)
[![DaisyUI](https://img.shields.io/badge/daisyui-v5-FF9903?logo=daisyui&logoColor=white)](https://daisyui.com)
[![Tailwind CSS](https://img.shields.io/badge/tailwind-v4-38BDF8?logo=tailwindcss&logoColor=white)](https://tailwindcss.com)
[![License: MIT](https://img.shields.io/badge/license-MIT-22c55e)](LICENSE)

Go · Templ · HTMX · Tailwind CSS · Alpine.js

**[Docs & Playground →](https://kraft-ui.appkrafterrz.com)**

</div>

---

Stop rewriting the same form fields, buttons, and cards on every project.
kraft-ui gives you a consistent, typed component library that integrates naturally
with server-side Go rendering and HTMX — no React, no JavaScript build step, no magic.

```go
// One line. Typed. Consistent. HTMX-ready.
@button.Button(button.Props{
    Variant: button.VariantPrimary,
    Attrs:   templ.Attributes{"hx-post": "/save", "hx-target": "#result"},
}) { Save changes }
```

---

## Why?

A form input field in raw Tailwind is 16 lines. With kraft-ui it is 4.

```html
<!-- Before: 16 lines, memorise every class, repeat on every field -->
<div class="flex flex-col gap-1.5">
  <label class="text-sm font-medium text-text-primary">Email</label>
  <input type="email" placeholder="you@example.com"
    class="rounded-md border border-error-500 px-3 py-2 text-sm bg-surface
           text-text-primary placeholder:text-text-muted focus:outline-none
           focus:ring-2 focus:ring-error-500 transition-colors" />
  <p class="text-xs text-error-600">Invalid email address</p>
</div>
```

```go
// After: typed, consistent, one place to update
@input.Input(input.Props{
    Type:  "email",
    Label: "Email",
    Error: "Invalid email address",
})
```

More importantly — every developer on your team writes it the same way.
When the design changes, you update one file.

---

## What's included

### Primitives — Tailwind CSS, typed Go props

| Component  | Import |
|------------|--------|
| Button     | `pkg/ui/button`   |
| Input      | `pkg/ui/input`    |
| Textarea   | `pkg/ui/textarea` |
| Card       | `pkg/ui/card`     |
| Badge      | `pkg/ui/badge`    |
| Alert      | `pkg/ui/alert`    |
| Avatar     | `pkg/ui/avatar`   |
| Spinner    | `pkg/ui/spinner`  |
| Progress   | `pkg/ui/progress` |
| Modal      | `pkg/ui/modal`    |
| Toast      | `pkg/ui/toast`    |
| Dropdown   | `pkg/ui/dropdown` |
| Tabs       | `pkg/ui/tabs`     |
| Table      | `pkg/ui/table`    |

### DaisyUI wrappers — 30+ themes, typed variants

| Component  | Before (raw DaisyUI) | After (kraft-ui) |
|------------|----------------------|------------------|
| Button     | `"btn btn-primary"`  | `daisyui.ButtonPrimary` |
| Badge      | `"badge badge-success"` | `daisyui.BadgeSuccess` |
| Alert      | `"alert alert-error"` | `daisyui.AlertError` |
| Card       | nested divs          | `daisyui.Card(...)` |
| Input      | fieldset boilerplate | `daisyui.Input(...)` |

Mistype `daisyui.BadgePrimry` — `go build` catches it. Raw class strings don't.

---

## Quick start

```bash
go get github.com/appkrafterrz/kraft-ui
```

```go
import (
    "github.com/appkrafterrz/kraft-ui/pkg/ui/button"
    "github.com/appkrafterrz/kraft-ui/pkg/ui/input"
    "github.com/appkrafterrz/kraft-ui/pkg/daisyui"
)

// Primitive
@button.Button(button.Props{Variant: button.VariantPrimary}) { Click me }

// DaisyUI wrapper
@daisyui.Input(daisyui.InputProps{
    Type:  "email",
    Label: "Email",
    Error: errors["email"],
})

// HTMX — pass any attribute via templ.Attributes
@button.Button(button.Props{
    Variant: button.VariantPrimary,
    Attrs: templ.Attributes{
        "hx-post":   "/api/submit",
        "hx-target": "#result",
        "hx-swap":   "innerHTML",
    },
}) { Submit }
```

### Copy-paste (recommended for production)

```bash
# Copy any component directly into your project — own it completely
cp -r pkg/ui/button your-project/components/button
```

No runtime dependency. No upgrade surprises. MIT licensed.

---

## Run locally

```bash
git clone https://github.com/appkrafterrz/kraft-ui
cd kraft-ui
make setup   # npm install + go mod tidy
make dev     # hot reload at http://localhost:3000
```

`make dev` runs templ, Tailwind, and the Go server in parallel with live reload.

---

## Stack

| Layer      | Technology |
|------------|------------|
| Language   | Go 1.23    |
| Templates  | [Templ](https://templ.guide) |
| Interactivity | [HTMX](https://htmx.org) |
| Styling    | [Tailwind CSS v4](https://tailwindcss.com) |
| Components | [DaisyUI v5](https://daisyui.com) |
| Client state | [Alpine.js](https://alpinejs.dev) |

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). PRs welcome — new components, better patterns, bug fixes.

## License

MIT — see [LICENSE](LICENSE).
