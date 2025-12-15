package handlers

import (
	"net/http"
	"producer-go/internal/kafka"
	"producer-go/internal/middleware"
)

type CatHandler struct {
}

func (h *CatHandler) AddCat(w http.ResponseWriter, r *http.Request) {

	cat, ok := middleware.GetValidatedCat(r)
	if !ok {
		http.Error(w, "validation middleware not applied", http.StatusInternalServerError)
		return
	}

	kafka.SendMessage(cat)
}
