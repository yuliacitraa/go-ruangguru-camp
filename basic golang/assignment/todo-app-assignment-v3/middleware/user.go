package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Kita dapat memperoleh session token dari requests cookies, ini disertakan dalam setiap request
        c, err := r.Cookie("session_token")
        if err != nil {
            if err == http.ErrNoCookie {
                // cookie tidak di set, return unauthorized status
                w.WriteHeader(http.StatusUnauthorized)
                json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})

                return
            }
            // Untuk jenis error lainnya, return bad request status
            w.WriteHeader(http.StatusBadRequest)
                json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})

            return
        }
        sessionToken := c.Value

        // Kita mendapatkan userSession dari session map
        userSession, exists := db.Sessions[sessionToken]

        if !exists {
            // Session token tidak ada pada session map, return unauthorized error
            w.WriteHeader(http.StatusUnauthorized)
                json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})

            return
        }

        if isExpired(userSession) {
            delete(db.Sessions, sessionToken)
            w.WriteHeader(http.StatusUnauthorized)
                json.NewEncoder(w).Encode(model.ErrorResponse{Error: "http: named cookie not present"})

            return
        }

        ctx := context.WithValue(r.Context(), "username", userSession.Username)
        next.ServeHTTP(w, r.WithContext(ctx))
	}) // TODO: replace this
}
