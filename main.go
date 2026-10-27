package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func setupRoutes(app *App) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", app.list)
	mux.HandleFunc("POST /tasks", app.create)
	mux.HandleFunc("GET /tasks/{id}", app.get)
	mux.HandleFunc("DELETE /tasks/{id}", app.delete)
	return mux
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func main() {
	app := &App{store: NewStore()}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, setupRoutes(app)))
}
