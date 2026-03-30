package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/appkrafterrz/kraft-ui/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := buildRouter()

	log.Printf("→ kraft-ui docs server running at http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func buildRouter() http.Handler {
	r := chi.NewRouter()

	// ── Global middleware ────────────────────────────────────────────────────
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	// ── Static assets ────────────────────────────────────────────────────────
	r.Handle("/static/*", http.StripPrefix("/static/",
		http.FileServer(http.Dir("web/static")),
	))

	// ── Docs site routes ─────────────────────────────────────────────────────
	h := handler.NewDocs()

	r.Get("/", h.Home)

	r.Route("/docs", func(r chi.Router) {
		r.Get("/introduction", h.Introduction)
		r.Get("/installation", h.Installation)
		r.Get("/cli", h.CLIPage)

		r.Route("/daisyui", func(r chi.Router) {
			r.Get("/", h.DaisyOverview)
			r.Get("/button", h.DaisyButton)
			r.Get("/badge", h.DaisyBadge)
			r.Get("/alert", h.DaisyAlert)
			r.Get("/card", h.DaisyCard)
			r.Get("/input", h.DaisyInput)
		})
		r.Get("/theming", h.Theming)

		r.Route("/components", func(r chi.Router) {
			// Primitives
			r.Get("/button", h.ButtonPage)
			r.Get("/input", h.InputPage)
			r.Get("/textarea", h.TextareaPage)
			r.Get("/card", h.CardPage)
			r.Get("/badge", h.BadgePage)
			r.Get("/alert", h.AlertPage)
			r.Get("/avatar", h.AvatarPage)
			r.Get("/spinner", h.SpinnerPage)
			r.Get("/progress", h.ProgressPage)

			// Composites
			r.Get("/modal", h.ModalPage)
			r.Get("/toast", h.ToastPage)
			r.Get("/dropdown", h.DropdownPage)
			r.Get("/tabs", h.TabsPage)
			r.Get("/table", h.TablePage)
		})
	})

	// ── Playground ───────────────────────────────────────────────────────────────
	r.Get("/playground", h.Playground)
	r.Get("/playground/preview/{component}", h.PlaygroundPreview)
	r.Get("/playground/component/{component}", h.PlaygroundComponent)

	return r
}
