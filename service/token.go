package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func init()  {
	newTokenService(tokenConfig)
}

type TokenConfig struct {
	Security  []byte
	Issuer    string
	ExpiresAt int
}

type TokenService struct {
	Config *TokenConfig
}

var tokenConfig = &TokenConfig{
	Security: []byte("1111221"),
	Issuer: "neko.studio",
	ExpiresAt: 3600,
}
var tokenService *TokenService
func newTokenService(config *TokenConfig) *TokenService {
	tokenService= &TokenService {
		Config: config,
	}
	return tokenService
}
func GetTokenService() *TokenService {
 return tokenService
}
func(ts *TokenService) CreateToken(username string) string {
	exp :=time.Duration(ts.Config.ExpiresAt)*time.Second
	claims := &jwt.StandardClaims{
		Subject:   username,
		ExpiresAt: time.Now().Add(exp).Unix(),
		Issuer: ts.Config.Issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(ts.Config.Security)
	if err != nil {
		return ""
	}
	return ss
}
func (ts *TokenService) ValidateToken(token string) (bool,*jwt.Token) {
	parse, err := jwt.Parse(token, ts.tokenHandler)
	if err != nil {
		return false,nil
	}
	return parse.Valid,parse
}

func (ts *TokenService)GetSubject(tokenString string) string {
	valid, token := ts.ValidateToken(tokenString)
	if valid {
		claims,ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return ""
		}
		if v,k:=claims["sub"];k{
			return v.(string)
		}
	}
	return ""
}
func (ts *TokenService)tokenHandler(token *jwt.Token) (interface{}, error) {
	return  ts.Config.Security, nil
}