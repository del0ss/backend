package auth

type TokenManager interface {
	GenerateJWT(id int, role int) (string, error)
	Parse(accessToken string) (interface{}, error)
	RefreshJWT() (string, error)
}
