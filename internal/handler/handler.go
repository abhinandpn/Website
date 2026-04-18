package handler

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// templateDir is the path to the HTML templates folder.
const templateDir = "web/templates"

// Handler holds shared dependencies for all route handlers.
type Handler struct {
	templates map[string]*template.Template
}

// New creates a Handler and pre-parses all HTML templates.
func New() (*Handler, error) {
	pages := []string{"index.html", "about.html", "form.html"}

	templates := make(map[string]*template.Template, len(pages))
	for _, page := range pages {
		tmplPath := filepath.Join(templateDir, page)
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			return nil, err
		}
		templates[page] = tmpl
	}

	return &Handler{templates: templates}, nil
}

// render is a helper that executes a named template and writes it to the response.
func (h *Handler) render(w http.ResponseWriter, name string, data any) {
	tmpl, ok := h.templates[name]
	if !ok {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("[ERROR] rendering template %s: %v", name, err)
		http.Error(w, "Render error", http.StatusInternalServerError)
	}
}

// ServeIndex handles GET /
func (h *Handler) ServeIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	h.render(w, "index.html", nil)
}

// ServeAbout handles GET /about
func (h *Handler) ServeAbout(w http.ResponseWriter, r *http.Request) {
	h.render(w, "about.html", nil)
}

// ServeForm handles GET /form
func (h *Handler) ServeForm(w http.ResponseWriter, r *http.Request) {
	h.render(w, "form.html", nil)
}

// HandleContact handles POST /contact (form submission from index.html)
func (h *Handler) HandleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	message := r.FormValue("message")

	// Log the submission — replace with DB insert in future.
	log.Printf("[CONTACT] Name: %s | Email: %s | Phone: %s | Message: %s",
		name, email, phone, message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok","message":"Thank you! We will get back to you soon."}`))
}
