package http

import (
	"backend/internal/config"
	"errors"
	"log"
	"net/http"
)

func StartServer(cfg config.HTTPServer, router http.Handler) {
	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	log.Println("🚀 Starting HTTP server on" + server.Addr + "...")
	err := server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Error while starting HTTP server: " + err.Error())
	}
}
