package api

import (
	"log"
	"net/http"
)

func RunAPI() {
	http.HandleFunc("GET /get/usuarios/", get_usuario)

	http.HandleFunc("GET /get/usuarios/all", get_all)

	http.HandleFunc("POST /post/usuarios/", post_usuario)

	http.HandleFunc("DELETE /delete/usuarios/", delete_user)

	log.Println("Servidor rodando em http://localhost:3000")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
