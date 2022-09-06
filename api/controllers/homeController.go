package controllers

import (
	"net/http"

	"github.com/LordCeilan/crud-restful-api-go/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this shitty API")
}
