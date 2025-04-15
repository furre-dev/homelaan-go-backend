package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/furre-dev/homelaan-go-backend/internal"
)

// TODO: HANDLE VALIDATION THROUGH FIREBASE SDK

// Dummy function. Replace this with real JWT verification or token lookup.
func validateBearerToken(token string) (string, error) {
	// Example: decode and verify JWT, then return user id or info.
	if token == "validtoken" {
		return "user-id-123", nil
	}
	return "", fmt.Errorf("invalid token")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")
			userID, err := validateBearerToken(token)
			if err == nil && userID != "" {
				ctx := context.WithValue(r.Context(), internal.UserIDKey, userID)
				r = r.WithContext(ctx)
			}
			// It's OK if no valid userID. Resolvers will check if required!
		}
		next.ServeHTTP(w, r)
	})
}
