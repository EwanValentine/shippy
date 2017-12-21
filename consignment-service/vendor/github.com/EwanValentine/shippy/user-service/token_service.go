package main

type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(token string) (interface{}, error) {
	return "", nil
}

func (srv *TokenService) Encode(data interface{}) (string, error) {
	return "", nil
}