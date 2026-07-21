package security

type TokenService interface {
	Generate(userID int64, roleID int64) (string, error)
	Parse(tokenString string) (*TokenClaims, error)
}
