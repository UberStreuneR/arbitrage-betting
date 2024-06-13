package auth

type Manager interface {
	Authenticate(string, string) (bool, error)
}

type authManager struct{}

func NewAuthManager() *authManager {
	return &authManager{}
}

func (authManager) Authenticate(username string, password string) (bool, error) {
	return true, nil
}
