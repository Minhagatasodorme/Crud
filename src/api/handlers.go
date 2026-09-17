package api

import (
	"crud/src/database"
	"net/http"
)

func get_usuario(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	email := r.URL.Query().Get("email")

	resp_json := database.GetUser(nome, email)

	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)

	w.Write(resp_json)
}

func get_all(w http.ResponseWriter, r *http.Request) {
	resp_json := database.GetUserAll()

	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusOK)

	w.Write(resp_json)
}

func post_usuario(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	email := r.URL.Query().Get("email")

	resp_json := database.PostUser(nome, email)

	w.Header().Set("Content-Type", "applications/json")

	w.WriteHeader(http.StatusCreated)

	w.Write(resp_json)
}

func delete_user(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	nome := r.URL.Query().Get("nome")
	email := r.URL.Query().Get("email")

	data := database.DeleteUser(id, nome, email)

	w.Header().Set("Content-Type", "Applications/json")

	w.WriteHeader(http.StatusOK)

	w.Write(data)
}
