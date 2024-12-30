package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/robertpuiu/go-coins-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) // show line and file when loging something
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API...")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
