package auth

type TokenManager interface {
	GenerateJWT(id int64, role int64) (string, error)
	Parse(accessToken string) (interface{}, error)
	RefreshJWT() (string, error)
}
