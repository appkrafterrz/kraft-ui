// kraft.js — Alpine.js stores + HTMX configuration for kraft-ui

// ── HTMX global config ────────────────────────────────────────────────────────
document.addEventListener('DOMContentLoaded', () => {
    if (typeof htmx !== 'undefined') {
        htmx.config.defaultSwapStyle = 'outerHTML';
        htmx.config.includeIndicatorStyles = false; // we use our own
    }
});

// ── Alpine.js stores ──────────────────────────────────────────────────────────
document.addEventListener('alpine:init', () => {

    // ── Modal store ───────────────────────────────────────────────────────────
    // Manages open/close state for all modals.
    //
    // Usage:
    //   Open:  $store.modals.open('my-modal')
    //   Close: $store.modals.close('my-modal')
    //   Check: $store.modals.isOpen('my-modal')
    //
    //   Or from Go via HTMX response header:
    //   w.Header().Set("HX-Trigger", `{"kraft:openModal": "my-modal"}`)
    Alpine.store('modals', {
        _open: new Set(),
        open(id) {
            this._open.add(id);
            document.body.style.overflow = 'hidden';
        },
        close(id) {
            this._open.delete(id);
            if (this._open.size === 0) {
                document.body.style.overflow = '';
            }
        },
        isOpen(id) {
            return this._open.has(id);
        },
        closeAll() {
            this._open.clear();
            document.body.style.overflow = '';
        },
    });

    // Listen for HTMX-triggered modal open/close events
    document.addEventListener('kraft:openModal', (e) => Alpine.store('modals').open(e.detail));
    document.addEventListener('kraft:closeModal', (e) => Alpine.store('modals').close(e.detail));

    // ── Sidebar store (docs site) ─────────────────────────────────────────────
    Alpine.store('sidebar', {
        open: false,
        toggle() { this.open = !this.open; },
        close() { this.open = false; },
    });

    // ── Toast helpers ─────────────────────────────────────────────────────────
    // Programmatic toast trigger (without HTMX)
    // Usage: $store.toasts.push({ variant: 'success', title: 'Done!', message: 'Saved.' })
    Alpine.store('toasts', {
        items: [],
        push(toast) {
            const id = `toast-${Date.now()}`;
            this.items.push({ id, ...toast });
            setTimeout(() => this.remove(id), 4200);
        },
        remove(id) {
            this.items = this.items.filter(t => t.id !== id);
        },
    });

    // ── Reusable Alpine components ────────────────────────────────────────────

    // dropdown: can be used as x-data="dropdown"
    Alpine.data('dropdown', () => ({
        open: false,
        toggle() { this.open = !this.open; },
        close() { this.open = false; },
    }));

    // clipboard: copy text to clipboard
    Alpine.data('clipboard', (text = '') => ({
        copied: false,
        text,
        async copy(t) {
            const value = t || this.text;
            try {
                await navigator.clipboard.writeText(value);
                this.copied = true;
                setTimeout(() => { this.copied = false; }, 2000);
            } catch {}
        },
    }));
});

// ── HTMX event bindings ───────────────────────────────────────────────────────

// Show/hide loading indicators for HTMX requests
document.addEventListener('htmx:beforeRequest', (e) => {
    const indicator = e.detail.elt.getAttribute('hx-indicator');
    if (indicator) {
        document.querySelectorAll(indicator).forEach(el => {
            el.classList.add('htmx-request');
        });
    }
});

document.addEventListener('htmx:afterRequest', (e) => {
    const indicator = e.detail.elt.getAttribute('hx-indicator');
    if (indicator) {
        document.querySelectorAll(indicator).forEach(el => {
            el.classList.remove('htmx-request');
        });
    }
});

// Close open modals on Escape key (backup for non-trap cases)
document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') {
        Alpine.store('modals').closeAll();
    }
});
