package headers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

type authContextKey string

func JWTExtractor() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header["Authorization"]
			if len(auth) > 0 {
				token := ""
				for _, header := range auth {
					if strings.HasPrefix(header, "bearer") {
						tokens := strings.Split(header, " ")
						token = tokens[1]
					}
				}

				if token != "" {
					k := authContextKey("auth")
					ctx := context.WithValue(r.Context(), k, token)
					r = r.WithContext(ctx)
				}
				// fmt.Println("Middleware", token)
			}

			next.ServeHTTP(w, r)
		})
	}
}

func GetHeaderAuth(ctx context.Context) string {
	k := authContextKey("auth")

	raw := ctx.Value(k)
	fmt.Println(raw)

	return "test"
}
