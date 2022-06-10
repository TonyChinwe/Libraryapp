package jwt

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/TonyChinwe/libraryapp/pkg/users/model"
	"github.com/TonyChinwe/libraryapp/utils/response"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	//"github.com/stdeemene/go-travel2/utils/response"
	"github.com/TonyChinwe/libraryapp/config/env"
)

type contextKey string

func (c contextKey) String() string {
	return string(c)
}

var (
	contextTokenClaims = contextKey("tokenClaims")
	contextEmailKey    = contextKey("username")
)

type TokenClaims struct {
	UserId   string `json:"id"`
	UserName string `json:"username"`
	UserRole string `json:"role"`
	jwt.StandardClaims
}

type TokenDetails struct {
	UserId              string `json:"id"`
	UserRole            string `json:"role"`
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	UserIsActive        bool   `json:"isActive"`
	AccessUUID          string `json:"access_token_uuid"`
	RefreshUUID         string `json:"refresh_token_uuid"`
	AccessTokenExpires  int64  `json:"access_token_exp"`
	RefreshTokenExpires int64  `json:"refresh_token_exp"`
}

func GenerateJwtToken(user *model.User) (*TokenDetails, error) {
	accessKey := env.GetEnvWithKey("JWT_ACCESS_KEY")
	refreshKey := env.GetEnvWithKey("JWT_REFRESH_KEY")
	issuer := env.GetEnvWithKey("JWT_ISSUER")
	td := &TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Hour * 24).Unix()
	td.AccessUUID = uuid.New().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New().String()
	var accessSigningKey = []byte(accessKey)
	var refreshSigningKey = []byte(refreshKey)

	claims := &TokenClaims{
		UserId:   user.ID.Hex(),
		UserName: user.Username,
		UserRole: user.Role,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			ExpiresAt: td.AccessTokenExpires,
			Subject:   user.Username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString(accessSigningKey)
	if err != nil {
		return nil, err
	}

	refreshToken, err := token.SignedString(refreshSigningKey)
	if err != nil {
		return nil, err
	}

	td.AccessToken = accessToken
	td.RefreshToken = refreshToken
	td.UserRole = user.Role
	td.UserId = user.ID.Hex()
	return td, nil
}

func ProtectApi(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessKey := env.GetEnvWithKey("JWT_ACCESS_KEY")

		if strings.Contains(r.URL.Path, "/auth/") || strings.Contains(r.URL.Path, "/swagger/") {
			h.ServeHTTP(w, r)
		} else {
			authorizationHeader := r.Header.Get("Authorization")
			if authorizationHeader == "" {
				response.BaseResponse(w, http.StatusUnauthorized, "An authorization header is required")
				return
			}

			tokenString := strings.Split(authorizationHeader, " ")
			if len(tokenString) != 2 {
				response.BaseResponse(w, http.StatusUnauthorized, "Please pass the authorization header as <Bearer APIKEY>")
				return
			}
			tknStr := tokenString[1]

			tc := TokenClaims{}
			token, err := jwt.ParseWithClaims(tknStr, &tc, func(token *jwt.Token) (interface{}, error) {
				return []byte(accessKey), nil
			})
			if err != nil || !token.Valid {
				response.BaseResponse(w, http.StatusUnauthorized, "Invalid Access Token")
			} else {
				NewContext(r.Context(), &tc)
				ctx := context.WithValue(r.Context(), contextEmailKey, tc.UserName)
				h.ServeHTTP(w, r.WithContext(ctx))

			}
		}

	})

}

// 	// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, tokenClaims *TokenClaims) context.Context {
	return context.WithValue(ctx, contextTokenClaims, tokenClaims)
}

// 	// FromContext returns the User value stored øin ctx, if any.
func FromContext(ctx context.Context) (*TokenClaims, bool) {
	tc, ok := ctx.Value(contextTokenClaims).(*TokenClaims)
	return tc, ok
}

// AuthToken gets the auth token from the context.
func ContextEmail(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(contextEmailKey).(string)
	return email, ok
}

func Refresh(w http.ResponseWriter, r *http.Request) (*TokenDetails, error) {

	td := &TokenDetails{}
	td.AccessTokenExpires = time.Now().Add(time.Hour * 24).Unix()
	td.AccessUUID = uuid.New().String()

	td.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUUID = uuid.New().String()

	return td, nil
}
