package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/EwanValentine/shippy/user-service/proto/user"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"github.com/micro/go-web"
)

const (
	defaultName = "Ewan Valentine"
	defaultEmail = "ewan.valentine89@gmail.com"
	defaultPassword = "Testing123"
	defaultCompany = "BBC"
)


func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	name := defaultName
	email := defaultEmail
	password := defaultPassword
	company := defaultCompany

	// Contact the server and print out its response.
	if len(os.Args) > 1 {
		name = os.Args[1]
		email = os.Args[2]
		password = os.Args[3]
		company = os.Args[4]
	}

	r, err := client.Create(context.TODO(), &pb.User{
		Name: name,
		Email: email,
		Password: password,
		Company: company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}
}
