package headers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"com.methompson/kcms-go/kcms"
	"com.methompson/kcms-go/kcms/jwtuserdata"
	"github.com/dgrijalva/jwt-go"
)

type authContextKey string

// JWTExtractor extracts the bearer token from the headers of a request and inserts it into the request context
func JWTExtractor(kcms kcms.KCMS) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header["Authorization"]
			if len(auth) > 0 {
				tokenString := ""

				// We will get all authorization headers, find one that has a value that starts with bearer and take the
				// value from that one. If the user sends more than on Authorization header, it's up to them to fix
				// their erroneous request.
				for _, header := range auth {
					if strings.HasPrefix(header, "bearer") {
						tokenString = strings.Split(header, " ")[1]
						break
					}
				}

				// If we found a token string, we'll go ahead and deal with it. Otherwise, we'll move on
				if tokenString != "" {
					// We're going to parse the JWT token and insert the data into a JWTUserData struct.
					decoded, err := jwt.ParseWithClaims(tokenString, &jwtuserdata.JWTUserData{}, func(token *jwt.Token) (interface{}, error) {
						_, ok := token.Method.(*jwt.SigningMethodHMAC)
						if !ok {
							return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
						}

						mySecret := []byte(kcms.JWTSecret)

						return mySecret, nil
					})

					// If no parsing errors, we're going to prep the context to save the JWTUserData struct
					if err == nil {
						claims, ok := decoded.Claims.(*jwtuserdata.JWTUserData)

						if ok && decoded.Valid {
							k := authContextKey("authToken")
							ctx := context.WithValue(r.Context(), k, claims)
							r = r.WithContext(ctx)
						}

					} else {
						// fmt.Println("Token Not Decoded")
					}
				} else {
					// fmt.Println("Empty Token")
				}
			}

			// All done, move to the next routing middleware
			next.ServeHTTP(w, r)
		})
	}
}

// GetHeaderAuth will retrieve the saved token from the headers, if it exists
func GetHeaderAuth(ctx context.Context) *jwtuserdata.JWTUserData {
	// Get the authToken from the headers
	k := authContextKey("authToken")
	token := ctx.Value(k)

	// We can't make a type assertion on nil, so we have to check if it's nil
	var decodedToken *jwtuserdata.JWTUserData
	// var decodedToken JWTUserData

	// If the token isn't nil, we assert that it's a jwt.MapClaims type
	if token != nil {
		// I'm not certain we need to maintain a reference, so this will pass a copy
		// decodedToken = *ctx.Value(k).(*JWTUserData)

		decodedToken = ctx.Value(k).(*jwtuserdata.JWTUserData)
	}

	return decodedToken
}
