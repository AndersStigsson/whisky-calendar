package middlewares

import (
	"net/http"

	"github.com/AndersStigsson/whisky-calendar/token"
)

func VerifyJWTToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		}

		_, err := token.ExtractJWTFromRequest(r)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
