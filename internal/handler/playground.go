package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/appkrafterrz/kraft-ui/web/pages"
)

func parseParams(r *http.Request) map[string]string {
	params := make(map[string]string)
	for k, v := range r.URL.Query() {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	return params
}

// Playground serves the interactive component explorer.
func (h *Docs) Playground(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("c")
	render(w, r, pages.Playground(active))
}

// PlaygroundPreview returns the full REPL panel (preview + controls) for HTMX injection.
func (h *Docs) PlaygroundPreview(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "component")
	render(w, r, pages.PlaygroundREPL(id, parseParams(r)))
}

// PlaygroundComponent returns just the live component for HTMX prop-change updates.
func (h *Docs) PlaygroundComponent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "component")
	render(w, r, pages.PlaygroundLiveComponent(id, parseParams(r)))
}
