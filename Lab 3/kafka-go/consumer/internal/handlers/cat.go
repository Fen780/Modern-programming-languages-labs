package handlers

import (
	"consumer/internal/entities"
	"database/sql"
	"encoding/json"
	"net/http"
)

type CatHandler struct {
	DB *sql.DB
}

func (h *CatHandler) GetAllCats(w http.ResponseWriter, r *http.Request) {
	rows, err := h.DB.Query("SELECT * FROM cats")
	if err != nil {
		http.Error(w, "db error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var res []entities.Cat

	for rows.Next() {
		var cat entities.Cat
		rows.Scan(&cat.ID, &cat.Name, &cat.Age, &cat.Breed, &cat.Weight)
		res = append(res, cat)
	}

	json.NewEncoder(w).Encode(res)
}

func (h *CatHandler) AddCat(cat *entities.Cat) error {
	_, err := h.DB.Exec("INSERT INTO cats (name, age, breed, weight) VALUES (?, ?, ?, ?)", cat.Name, cat.Age, cat.Breed, cat.Weight)
	if err != nil {
		return err
	}
	return nil
}
