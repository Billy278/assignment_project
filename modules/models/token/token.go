package models

type TokenType string

const (
	ID_TOKEN      TokenType = "id_token"
	ACCESS_TOKEN  TokenType = "access_token"
	REFRESH_TOKEN TokenType = "refresh_token"
)

type TokensPromo struct {
	IDToken      string `json:"id_token"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DefaultClaimPromo struct {
	Expired   int       `json:"exp"`
	NotBefore int       `json:"nbf"`
	IssuedAt  int       `json:"iat"`
	Issuer    string    `json:"iss"`
	Audience  string    `json:"aud"`
	JTI       string    `json:"jti"`
	Type      TokenType `json:"typ"`
}

type IdClaimPromo struct {
	UserId string `json:"preffend_user_id"`
	Name   string `json:"preffend_name"`
	Email  string `json:"preffend_email"`
}
type AccessClaimPromo struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}
