package services 

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var tokenSecret = "my_secret_key"

type AccessSecretClient struct {
    ClientID string
    AccessTokenSecret string
}

var clientSecrets []AccessSecretClient

type AuthorizeClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AccessClaims struct {
	ClientId string `json:"client_id"`
	jwt.RegisteredClaims
}

func (c *AuthorizeClaims) SetExpiry(expiry time.Time) {
	c.ExpiresAt = jwt.NewNumericDate(expiry)
}

func (c *AccessClaims) SetExpiry(expiry time.Time) {
	c.ExpiresAt = jwt.NewNumericDate(expiry)
}

func GenerateToken(claims jwt.Claims, expiration time.Duration, signingKey string) (string, error) {
	expirationTime := time.Now().Add(expiration)
	if c, ok := claims.(interface{ SetExpiry(time.Time) }); ok {
		c.SetExpiry(expirationTime)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT[T jwt.Claims](tokenString string, secret string, claims T) (T, error) {
    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        return claims, err
    }

    if !token.Valid {
        return claims, fmt.Errorf("invalid token")
    }

    return claims, nil
}

func GenerateAuthorizationToken(userID string) (string, error) {
    claims := &AuthorizeClaims{Username: userID}	
    return GenerateToken(claims, 15*time.Minute, tokenSecret)
}

func GenerateAccessToken(clientID string, accessTokenSecret string) (string, error) {
    clientSecrets = append(clientSecrets, AccessSecretClient{clientID, accessTokenSecret})

    claims := &AccessClaims{ClientId: clientID}
    return GenerateToken(claims, 1*time.Hour, accessTokenSecret) 
}

func ValidateAuthorizationToken(tokenString string) (*AuthorizeClaims, error) {
    claims := &AuthorizeClaims{}
    return ValidateJWT(tokenString, tokenSecret, claims)
}

func ValidateAccessToken(tokenString string, clientID string) (*AccessClaims, error) {
    claims := &AccessClaims{}

    secret, found := findAccessTokenByClientId(clientID)
    if found == false {
        return nil, fmt.Errorf("This client is not registered on auth server")
    }
    return ValidateJWT(tokenString, secret, claims)
}

func findAccessTokenByClientId(clientID string) (string, bool) {
    for _, item := range(clientSecrets) {
        if item.ClientID == clientID {
            return item.AccessTokenSecret, true
        }
    }

    return "", false
}

