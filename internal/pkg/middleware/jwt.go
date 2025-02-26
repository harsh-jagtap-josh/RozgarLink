package middleware

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func GenerateToken(userID int, role string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("JWT_PRIVATE_KEY")
	jwtToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

// Middleware to validate JWT and extract user info
func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		err := godotenv.Load()
		if err != nil {
			logger.Errorw(ctx, "error occured while loading env")
			http.Error(w, "failed to fetch env - "+err.Error(), http.StatusInternalServerError)
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			logger.Errorw(ctx, "missing authorization token", zap.String("URL", r.URL.Path), zap.String("Method", r.Method))
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.Split(authHeader, "Bearer ")
		if len(tokenStr) != 2 {
			logger.Errorw(ctx, "invalid authorization token format", zap.String("URL", r.URL.Path), zap.String("Method", r.Method))
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		secret := os.Getenv("JWT_PRIVATE_KEY")

		token, err := jwt.Parse(tokenStr[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			logger.Errorw(ctx, "invalid authorization token", zap.String("URL", r.URL.Path), zap.String("Method", r.Method))
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Extract jst data
		data, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Errorw(ctx, "invalid jwt data", zap.String("URL", r.URL.Path), zap.String("Method", r.Method))
			http.Error(w, "invalid token data", http.StatusUnauthorized)
			return
		}

		ctx = context.WithValue(r.Context(), "user_id", int(data["user_id"].(float64))) // pass the jst data into request context
		ctx = context.WithValue(ctx, "role", data["role"].(string))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RequireWorkerRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("role")
		if userRole != "worker" {
			logger.Errorw(r.Context(), "unauthorized access", zap.String("required_role", "worker"))
			http.Error(w, "unauthorized access to api", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireEmployerRole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Context().Value("role")
		if userRole != "worker" {
			logger.Errorw(r.Context(), "unauthorized access", zap.String("required_role", "worker"))
			http.Error(w, "unauthorized access to api", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireSameUserOrAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authUserID, ok := r.Context().Value("user_id").(int)
		if !ok {
			http.Error(w, "unauthorized access", http.StatusUnauthorized)
			return
		}

		role, _ := r.Context().Value("role").(string)

		vars := mux.Vars(r)
		paramID, err := strconv.Atoi(vars[role+"_id"])
		if err != nil {
			logger.Errorw(r.Context(), "invalid user id provided", zap.Error(err), zap.String("id", vars[role+"_id"]))
			http.Error(w, "invalid user id - "+err.Error(), http.StatusBadRequest)
			return
		}

		// allow access only if it the user itself trying to access its data or its the Admin
		if role != "admin" && authUserID != paramID {
			logger.Errorw(r.Context(), "access forbidden")
			http.Error(w, "forbidden: you are not authorized", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
