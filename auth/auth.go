package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/furre-dev/homelaan-go-backend/firebase"
	"github.com/furre-dev/homelaan-go-backend/internal"
)

func serializeBearerToFireabaseTokenUid(idToken string) (string, error) {
	app, err := firebase.InitFirebaseApp()
	if err != nil {
		return "", fmt.Errorf("error initializing firebase app: %v", err)
	}

	client, err := app.Auth(context.Background())
	if err != nil {
		return "", fmt.Errorf("error initializing firebase app: %v", err)
	}

	token, tokenErr := client.VerifyIDToken(context.Background(), idToken)
	if tokenErr != nil {
		return "", fmt.Errorf("error verifying idtoken: %v", err)
	}

	return token.UID, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			userID, err := serializeBearerToFireabaseTokenUid(token)
			if err != nil {
				println(err.Error())
			}
			if err == nil && userID != "" {
				ctx := context.WithValue(r.Context(), internal.UserIDKey, userID)
				r = r.WithContext(ctx)
			}
			// It's OK if no valid userID. Resolvers will check if required!
		}
		next.ServeHTTP(w, r)
	})
}
