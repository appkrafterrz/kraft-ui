package handler

import (
	"net/http"

	"github.com/a-h/templ"

	docsPages "github.com/appkrafterrz/kraft-ui/web/pages/docs"
	"github.com/appkrafterrz/kraft-ui/web/pages"
)

// Docs holds all handlers for the kraft-ui documentation site.
type Docs struct{}

func NewDocs() *Docs {
	return &Docs{}
}

// render is a helper that renders a templ component to the response.
func render(w http.ResponseWriter, r *http.Request, c templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := c.Render(r.Context(), w); err != nil {
		http.Error(w, "render error: "+err.Error(), http.StatusInternalServerError)
	}
}

// ── Site pages ────────────────────────────────────────────────────────────────

func (h *Docs) Home(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.Home())
}

func (h *Docs) Introduction(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.IntroductionPage())
}

func (h *Docs) CLIPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.CLIPage())
}

// ── DaisyUI pages ─────────────────────────────────────────────────────────────

func (h *Docs) DaisyOverview(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DaisyOverviewPage())
}

func (h *Docs) DaisyButton(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DaisyButtonPage())
}

func (h *Docs) DaisyBadge(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DaisyBadgePage())
}

func (h *Docs) DaisyAlert(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DaisyAlertPage())
}

func (h *Docs) DaisyCard(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DaisyCardPage())
}

func (h *Docs) DaisyInput(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DaisyInputPage())
}

func (h *Docs) Installation(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.InstallationPage())
}

func (h *Docs) Theming(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.ThemingPage())
}

// ── Component pages ───────────────────────────────────────────────────────────

func (h *Docs) ButtonPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.ButtonPage())
}

func (h *Docs) InputPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.InputPage())
}

func (h *Docs) TextareaPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.TextareaPage())
}

func (h *Docs) CardPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.CardPage())
}

func (h *Docs) BadgePage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.BadgePage())
}

func (h *Docs) AlertPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.AlertPage())
}

func (h *Docs) AvatarPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.AvatarPage())
}

func (h *Docs) SpinnerPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.SpinnerPage())
}

func (h *Docs) ProgressPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.ProgressPage())
}

func (h *Docs) ModalPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.ModalPage())
}

func (h *Docs) ToastPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.ToastPage())
}

func (h *Docs) DropdownPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.DropdownPage())
}

func (h *Docs) TabsPage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.TabsPage())
}

func (h *Docs) TablePage(w http.ResponseWriter, r *http.Request) {
	render(w, r, docsPages.TablePage())
}
