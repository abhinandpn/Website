package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abhinandpn/Website/internal/config"
	"github.com/abhinandpn/Website/internal/handler"
	"github.com/abhinandpn/Website/internal/middleware"
)

func main() {
	// Load config from environment
	cfg := config.Load()

	// Initialise handlers (pre-parses all HTML templates)
	h, err := handler.New()
	if err != nil {
		log.Fatalf("[FATAL] Failed to load templates: %v", err)
	}

	// Create a new ServeMux (router)
	mux := http.NewServeMux()

	// Static files — CSS, images, fonts etc.
	staticFS := http.FileServer(http.Dir("web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFS))

	// Page routes
	mux.HandleFunc("/", h.ServeIndex)
	mux.HandleFunc("/about", h.ServeAbout)
	mux.HandleFunc("/form", h.ServeForm)

	// Form submission API endpoint
	mux.HandleFunc("/contact", h.HandleContact)

	// Wrap the mux with middleware (Recovery → CORS → Logger → mux)
	wrappedMux := middleware.Chain(mux,
		middleware.Recovery,
		middleware.CORS,
		middleware.Logger,
	)

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Stock-Ed server running on http://localhost%s", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		log.Fatalf("[FATAL] Server error: %v", err)
	}
}
