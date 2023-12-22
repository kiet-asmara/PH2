package middleware

import (
	"context"
	"database/sql"
	"encoding/json"
	"example/entity"
	"example/helpers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ResponseJson(w http.ResponseWriter, code int, body any) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}

type Auth struct {
	DB *sql.DB
}

func NewAuth(db *sql.DB) Auth {
	return Auth{DB: db}
}

func (auth Auth) Authentication(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		auth_token := r.Header.Get("Authorization")

		if auth_token == "" {
			ResponseJson(w, http.StatusUnauthorized, map[string]any{
				"message": "please provide valid access token",
			})
			return
		}

		claims, err := helpers.DecodeToken(auth_token)
		if err != nil {
			ResponseJson(w, http.StatusUnauthorized, map[string]any{
				"message": "please provide valid access token",
			})
			return
		}

		row, err := auth.DB.Query(
			`
			SELECT id, email, address
			FROM Users
			WHERE id = ?
			LIMIT 1
			`,
			claims["id"],
		)

		if err != nil || !row.Next() {
			ResponseJson(w, http.StatusUnauthorized, map[string]any{
				"message": "please provide valid access token",
			})
			return
		}

		loggedInUser := entity.User{}
		row.Scan(&loggedInUser.Id, &loggedInUser.Email, &loggedInUser.Address)

		reqWithUser := r.WithContext(context.WithValue(r.Context(), "user", loggedInUser))

		next(w, reqWithUser, p)
	}
}
