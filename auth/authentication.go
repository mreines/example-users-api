package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
)

var sst = []byte("BUcQ1zIUiAZcgLOmRfYhmUBjBzXuEe3uWW5jYYEfJk8HW3KrHf8dFUJ55q0n2Huz")

type Authenticator struct {
}

type ClaimKey interface {
}

var UserKey = ClaimKey("Username")

type customClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Validate a google login token.
func (a *Authenticator) TokenSignIn(w http.ResponseWriter, r *http.Request) {
	validator, err := idtoken.NewValidator(r.Context(),
		option.WithHTTPClient(http.DefaultClient))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}

	r.ParseForm()
	id := r.Form["idtoken"][0]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal("tokenid parameter is required")
		return
	}

	payload, err := validator.Validate(r.Context(), id, "")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	if got, ok := payload.Claims["email"]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalf("missing email claim. payload.Claims = %+v", payload.Claims)
	} else {
		got, ok := got.(string)
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalf("email wasn't a string. payload.Claims = %+v", payload.Claims)
		} else {
			a.buildJwt(w, got)
		}
	}
}

func (a *Authenticator) buildJwt(w http.ResponseWriter, email string) {
	var secondsInDay = 86400
	claims := customClaims{
		Username: email,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(secondsInDay),
			Issuer:    "http://users-database.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(sst)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
	fmt.Fprint(w, signedToken)
}

func TokenCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/v1") {
			bearerToken := r.Header.Get("Authorization")
			if bearerToken == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				tokenized := strings.Split(bearerToken, " ")
				sentToken := tokenized[1]
				// Initialize a new instance of `Claims`
				claims := &customClaims{}

				// Parse the JWT string and store the result in `claims`.
				// Note that we are passing the key in this method as well. This method will return an error
				// if the token is invalid (if it has expired according to the expiry time we set on sign in),
				// or if the signature does not match
				token, err := jwt.ParseWithClaims(sentToken, claims, func(token *jwt.Token) (interface{}, error) {
					return sst, nil
				})
				if err != nil {
					if err == jwt.ErrSignatureInvalid {
						w.WriteHeader(http.StatusUnauthorized)
						return
					}
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if !token.Valid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				r = r.WithContext(context.WithValue(r.Context(), UserKey, claims.Username))
			}
		}
		next.ServeHTTP(w, r)
	})
}
