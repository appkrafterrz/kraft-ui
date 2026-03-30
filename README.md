# kraft-ui

Typed UI components for the GOTH stack — Go, Templ, HTMX, Tailwind CSS, and Alpine.js.

Stop rewriting the same form fields, buttons, and cards. kraft-ui gives you a consistent,
typed component library that integrates naturally with server-side Go rendering and HTMX
interactions — no React, no build step for components, no JavaScript framework.

---

## What's included

**Primitives** — components built on Tailwind CSS with typed Go props:

| Component  | Import path                           |
|------------|---------------------------------------|
| Button     | `github.com/appkrafterrz/kraft-ui/pkg/ui/button`   |
| Input      | `github.com/appkrafterrz/kraft-ui/pkg/ui/input`    |
| Textarea   | `github.com/appkrafterrz/kraft-ui/pkg/ui/textarea` |
| Card       | `github.com/appkrafterrz/kraft-ui/pkg/ui/card`     |
| Badge      | `github.com/appkrafterrz/kraft-ui/pkg/ui/badge`    |
| Alert      | `github.com/appkrafterrz/kraft-ui/pkg/ui/alert`    |
| Avatar     | `github.com/appkrafterrz/kraft-ui/pkg/ui/avatar`   |
| Spinner    | `github.com/appkrafterrz/kraft-ui/pkg/ui/spinner`  |
| Progress   | `github.com/appkrafterrz/kraft-ui/pkg/ui/progress` |
| Modal      | `github.com/appkrafterrz/kraft-ui/pkg/ui/modal`    |
| Toast      | `github.com/appkrafterrz/kraft-ui/pkg/ui/toast`    |
| Dropdown   | `github.com/appkrafterrz/kraft-ui/pkg/ui/dropdown` |
| Tabs       | `github.com/appkrafterrz/kraft-ui/pkg/ui/tabs`     |
| Table      | `github.com/appkrafterrz/kraft-ui/pkg/ui/table`    |

**DaisyUI wrappers** — DaisyUI v5 components wrapped as typed Templ components:

| Component  | Import path                                   |
|------------|-----------------------------------------------|
| Button     | `github.com/appkrafterrz/kraft-ui/pkg/daisyui` |
| Badge      | `github.com/appkrafterrz/kraft-ui/pkg/daisyui` |
| Alert      | `github.com/appkrafterrz/kraft-ui/pkg/daisyui` |
| Card       | `github.com/appkrafterrz/kraft-ui/pkg/daisyui` |
| Input      | `github.com/appkrafterrz/kraft-ui/pkg/daisyui` |

---

## Quick start

```bash
go get github.com/appkrafterrz/kraft-ui
```

Use a primitive:

```go
import "github.com/appkrafterrz/kraft-ui/pkg/ui/button"

@button.Button(button.Props{Variant: button.VariantPrimary}) {
    Click me
}
```

Use a DaisyUI wrapper:

```go
import "github.com/appkrafterrz/kraft-ui/pkg/daisyui"

@daisyui.Input(daisyui.InputProps{
    Type:  "email",
    Label: "Email",
    Error: errors["email"],
})
```

Pass HTMX attributes to any component:

```go
@button.Button(button.Props{
    Variant: button.VariantPrimary,
    Attrs: templ.Attributes{
        "hx-post":   "/api/save",
        "hx-target": "#result",
        "hx-swap":   "innerHTML",
    },
}) { Save }
```

---

## Why kraft-ui?

**For primitives:** a form input field in raw Tailwind is 16 lines. With kraft-ui it is 4.
More importantly, every developer on your team writes it the same way.
When the design changes, you update one file.

**For DaisyUI wrappers:** DaisyUI is great, but using it directly in Templ means building
class strings by hand, no IDE autocomplete, and typos that only show up at runtime.
kraft-ui wraps every DaisyUI component in a typed Go struct. Mistype a variant constant
and `go build` catches it before your users do.

```
// Before
<span class={
  func() string {
    switch order.Status {
    case "shipped":   return "badge badge-primary"
    case "delivered": return "badge badge-success"
    default:          return "badge badge-neutral"
    }
  }()
}>{ order.Status }</span>

// After
@daisyui.Badge(daisyui.BadgeProps{
    Variant: orderVariant(order.Status),
}) { { order.Status } }
```

---

## Docs and playground

Run the docs site locally:

```bash
make setup     # install npm deps and tidy go modules
make dev       # hot reload: templ + css + go (requires air)
```

Or for a quick one-shot build:

```bash
make run       # build and run once at http://localhost:3000
```

Visit `http://localhost:3000` for the full component docs and interactive playground.

---

## Copy-paste or module

**As a module** — `go get` and import as shown above. Pull in updates with `go get -u`.

**Copy-paste** — copy any component from `pkg/ui/` or `pkg/daisyui/` directly into your project
and own it completely. No runtime dependency, no upgrade surprises. This is the recommended
approach for production apps.

---

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

---

## License

MIT — see [LICENSE](LICENSE).
