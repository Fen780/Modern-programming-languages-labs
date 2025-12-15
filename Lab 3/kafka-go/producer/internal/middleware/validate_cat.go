package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"producer-go/internal/entities"
)

type myCat struct{}

var catContextKey = myCat{}

func ValidateCat(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var cat entities.Cat

		err := json.NewDecoder(r.Body).Decode(&cat)
		if err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}

		err = validate(cat)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), catContextKey, cat)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validate(c entities.Cat) error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Age <= 0 {
		return errors.New("age must be > 0")
	}
	if c.Breed == "" {
		return errors.New("breed is required")
	}
	if c.Weight <= 0 {
		return errors.New("weight must be > 0")
	}
	return nil
}

func GetValidatedCat(r *http.Request) (entities.Cat, bool) {
	cat, ok := r.Context().Value(catContextKey).(entities.Cat)
	return cat, ok
}
