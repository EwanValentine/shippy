package main

type Authable interface {
	Decode(token string) (interface{}, error)
	Encode(data interface{}) (string, error)
}

type TokenService {
	repo Repository
}

func (srv *TokenService) Decode(token string) (interface{}, error) {

}

func (srv *TokenService) Endode(data interface{}) (string, error) {

}