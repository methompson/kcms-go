package headers

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"com.methompson/kcms-go/kcms"
	"com.methompson/kcms-go/kcms/jwtuserdata"
	"github.com/dgrijalva/jwt-go"
)

// We'll just run JWTExtractor and assign it to a variable of the proper type to
// make sure that it returns what is necessary
func TestJWTExtractorReturnsProperResult(t *testing.T) {
	k := &kcms.KCMS{}

	var result func(http.Handler) http.Handler

	result = JWTExtractor(k)

	reflect.TypeOf(result)
	// fmt.Println(valType)
}

func TestJWTExtractorResultReturnsProperResult(t *testing.T) {
	k := &kcms.KCMS{}

	extr := JWTExtractor(k)

	var result http.Handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	result = extr(handler)
	reflect.TypeOf(result)
}

func jwtShouldBeNil(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := authContextKey("authToken")
		ctx := r.Context()
		requestToken := ctx.Value(key)

		if requestToken != nil {
			t.Error("JWT is not nil when it should have been nil")
		}
	}
}

func jwtShouldNotBeNil(t *testing.T) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		key := authContextKey("authToken")
		ctx := r.Context()
		requestToken := ctx.Value(key)

		if requestToken == nil {
			t.Error("JWT is nil when it should have contained data")
		}
	}
}

func makeDummyJwt(exp int64, secret string) (string, error) {
	claims := jwtuserdata.JWTUserData{
		ID:        "69",
		FirstName: "firstName",
		LastName:  "lastName",
		Username:  "username",
		Email:     "email",
		UserType:  "userType",
		UserMeta:  "{}",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func TestExtractionExtractsAuthorizationData(t *testing.T) {
	// Making a dummy JWT
	exp := time.Now().Unix() + 4*60*60
	k := &kcms.KCMS{}

	signedToken, err := makeDummyJwt(exp, k.JWTSecret)

	if err != nil {
		t.Fatal(err)
	}

	// Making a dummy request and setting the token in the header
	request := httptest.NewRequest("POST", "localhost:8080", nil)
	request.Header.Set("Authorization", "bearer "+signedToken)

	// Making a new handler that will check that things have been set in
	// the context
	nextHandler := http.HandlerFunc(jwtShouldNotBeNil(t))

	handler := JWTExtractor(k)(nextHandler)
	handler.ServeHTTP(httptest.NewRecorder(), request)
}

func TestExtractionDoesntExtractExpiredData(t *testing.T) {
	// Making a dummy JWT
	exp := time.Now().Unix() - 4*60*60

	k := &kcms.KCMS{}

	signedToken, err := makeDummyJwt(exp, k.JWTSecret)

	if err != nil {
		t.Fatal(err)
	}

	// Making a dummy request and setting the token in the header
	request := httptest.NewRequest("POST", "localhost:8080", nil)
	request.Header.Set("Authorization", "bearer "+signedToken)

	// Making a new handler that will check that things have been set in
	// the context
	nextHandler := http.HandlerFunc(jwtShouldBeNil(t))

	handler := JWTExtractor(k)(nextHandler)
	handler.ServeHTTP(httptest.NewRecorder(), request)
}

func TestExtractionDoesntExtractWithInvalidSecret(t *testing.T) {
	// Making a dummy JWT
	exp := time.Now().Unix() - 4*60*60

	k := &kcms.KCMS{}

	signedToken, err := makeDummyJwt(exp, "A Test Secret that doesn't work")

	if err != nil {
		t.Fatal(err)
	}

	// Making a dummy request and setting the token in the header
	request := httptest.NewRequest("POST", "localhost:8080", nil)
	request.Header.Set("Authorization", "bearer "+signedToken)

	// Making a new handler that will check that things have been set in
	// the context
	nextHandler := http.HandlerFunc(jwtShouldBeNil(t))

	handler := JWTExtractor(k)(nextHandler)
	handler.ServeHTTP(httptest.NewRecorder(), request)
}

func TestExtractionDoesntExtractImproperHeaders(t *testing.T) {
	// Making a dummy JWT
	exp := time.Now().Unix() + 4*60*60

	k := &kcms.KCMS{}

	signedToken, err := makeDummyJwt(exp, k.JWTSecret)

	if err != nil {
		t.Fatal(err)
	}

	// Making a new handler that will check that things have been set in
	// the context
	nextHandler := http.HandlerFunc(jwtShouldBeNil(t))
	handler := JWTExtractor(k)(nextHandler)

	var request *http.Request

	headers := make([]string, 0)
	headers = append(headers, "bear "+signedToken)
	headers = append(headers, "bearer"+signedToken)
	headers = append(headers, signedToken)
	headers = append(headers, signedToken+" bearer")

	for _, h := range headers {
		// Making a dummy request and setting the token in the header
		request = httptest.NewRequest("POST", "localhost:8080", nil)
		request.Header.Set("Authorization", h)

		handler.ServeHTTP(httptest.NewRecorder(), request)
	}
}

func TestExtractionDoesntHaveData(t *testing.T) {
	k := &kcms.KCMS{}

	// Making a dummy request and setting the token in the header
	request := httptest.NewRequest("POST", "localhost:8080", nil)

	// Making a new handler that will check that things have not been set in
	// the context
	nextHandler := http.HandlerFunc(jwtShouldBeNil(t))

	handler := JWTExtractor(k)(nextHandler)
	handler.ServeHTTP(httptest.NewRecorder(), request)
}
