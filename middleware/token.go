package middleware

import (
  "net/http"
  "log"

  "test/auth"
)

func ValidTokenMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Do stuff here
        log.Println(r.RequestURI)

        token := auth.ParseToken(r.Header["Token"][0])
        if token != nil {
          if token.Valid {
            next.ServeHTTP(w, r)
          } else {
            w.WriteHeader(http.StatusNotFound)
        		w.Write([]byte("Not Logged in"))
            return
          }
        } else {
          w.WriteHeader(http.StatusNotFound)
          w.Write([]byte("Not Logged in"))
          return
        }
    })
}
