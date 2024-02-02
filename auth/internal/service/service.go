package service

type AuthService struct {
}

func (a *AuthService) Register(email, password string) (string, error) {

}

func (a *AuthService) Login(email, password string) (string, error) {

}

func (a *AuthService) ItsValid(token string) (bool, error) {

}
