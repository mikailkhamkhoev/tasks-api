package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type App struct {
	store *Store
}

func (a *App) list(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, a.store.List())
}

func (a *App) create(w http.ResponseWriter, r *http.Request) {
	var in Task
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		http.Error(w, "bad json", 400)
		return
	}
	if in.Title == "" {
		http.Error(w, "title required", 400)
		return
	}
	writeJSON(w, 201, a.store.Create(in.Title))
}

func (a *App) get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad id", 400)
		return
	}
	t, ok := a.store.Get(id)
	if !ok {
		http.Error(w, "not found", 404)
		return
	}
	writeJSON(w, 200, t)
}

func (a *App) delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "bad id", 400)
		return
	}
	if !a.store.Delete(id) {
		http.Error(w, "not found", 404)
		return
	}
	w.WriteHeader(204)
}
