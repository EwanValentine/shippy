package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/EwanValentine/shippy/shippy-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type handler struct {
	repository   Repository
	tokenService authable
}

func (s *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	result, err := s.repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}

	user := UnmarshalUser(result)
	res.User = user

	return nil
}

func (s *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}

	users := UnmarshalUserCollection(results)
	res.Users = users

	return nil
}

func (s *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return err
	}

	token, err := s.tokenService.Encode(req)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	log.Println("user:", req)
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	if err := s.repository.Create(ctx, MarshalUser(req)); err != nil {
		return err
	}

	// Strip the password back out, so's we're not returning it
	req.Password = ""
	res.User = req

	return nil
}

func (s *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	claims, err := s.tokenService.Decode(req.Token)
	if err != nil {
		return err
	}

	if claims.User.Id == "" {
		return errors.New("invalid user")
	}

	res.Valid = true
	return nil
}
