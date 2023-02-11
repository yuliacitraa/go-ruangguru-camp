package main

import (
	"net/http"
)

func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
            return
        }

        next.ServeHTTP(w, r)

	}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("role")
		if role != "ADMIN" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Role not authorized"))
		} else {
			next.ServeHTTP(w, r)
		}


	}) // TODO: replace this
}

func main() {
	http.Handle("/admin", RequestMethodGetMiddleware(AdminMiddleware(AdminHandler())))

	http.ListenAndServe("localhost:8080", nil)
}
